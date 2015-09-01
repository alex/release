# "Latest" Artifacts

* Kernel CloudFormation template on s3://convox/release/latest/formation.json
* Kernel Lambda binary on s3://convox/release/latest/formation.zip
* `app`, `build`, `kernel`, `registry`, and `service` images tagged with "latest" on Docker Hub
* CLI binary on https://www.convox.com/downloads/{arch}/convox.zip
* A pointer to a version tag, i.e. '20150901101300', on s3://convox/release/latest/version

# Version Tag Artifacts

Every artifact is built and tagged with a timestamp in a "YYYYMMDDHHMMSS" format.

* A kernel CloudFormation template on s3://convox/release/20150902130909/formation.json
* A kernel Lambda binary on s3://convox/release/20150902130909/formation.zip
* `app`, `build`, `kernel`, `registry`, and `service` images tagged with "20150902130909" on Docker Hub

It is currently not possible / easy to get a versioned CLI artifact.

# Publishing a "latest" Release

When you are confident that everything in master will be an awesome experience
for new users, publish new "latest" artifacts:

```bash
$ convox switch demo.convox.io
$ convox run --app release release cli
$ convox run --app release release kernel --latest
```

**Note the `--latest` flag.**

Wait ~5m to see a "kernel released: 20150901101300" message in Slack.

Wait ~20m for https://circleci.com/gh/convox/ci/tree/master to run to "success".

If CI is not green, you are not done.

* Trigger a rebuild to test if the error was transient
* Browse the CircleCI artifacts to determine a root cause

If you discover a show-stopping bug in a "latest" release, prioritize work to 
fix the root cause or revert offending changes until CI is green.

# Releasing a Small Experimental Kernel Change

You can still confidently merge small Pull Requests or patches into master,
knowing that nobody will get them until a `--latest` release is published. This
is useful for small changes that require more testing on AWS.

```bash
$ convox switch demo.convox.io
$ convox run --app release release kernel
```

**Note the lack of the `--latest` flag.**

Wait ~5m to see a "kernel released: 20150902130909" message in Slack.

Wait ~20m to see CircleCI results.

Now a developer can perform manual testing with:

```bash
$ convox switch convox-862259866.us-east-1.elb.amazonaws.com
$ convox system update 20150902130909
```

Or:

```bash
$ VERSION=20150902130909 convox install
```

When you are satisfied that the change is working, publish a "latest" release.

# Releasing a Large Experimental System-Wide Change

Changes that span repos, even if just CLI and Kernel, are much harder to verify.
You can release artifacts 

* Pick a feature branch name, i.e. "multi-region"
* Push "multi-region" branches to GitHub for every component you want to test together, i.e. `cli`, `kernel` and `app`

**Note if the release and CI scripts do not see a "multi-region" branch they will use "master"**

```bash
$ convox switch demo.convox.io
$ convox run --app release release kernel --branch multi-region
```

**Note the `--branch` flag.**

Wait ~5m to see a "kernel released: 20150905161552-multi-region" message in Slack.

Wait ~20m to see CircleCI results.

Now a developer can perform manual testing with:

```bash
$ convox switch convox-862259866.us-east-1.elb.amazonaws.com
$ convox system update 20150905161552-multi-region
```

Or:

```bash
$ VERSION=20150905161552-multi-region convox install
```


