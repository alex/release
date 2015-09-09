# Kernel Artifacts

Every artifact is built and tagged with a timestamp in a "YYYYMMDDHHMMSS" format.

Some artifacts also have a branch suffix in a "YYYYMMDDHHMMSS-branch-name" format.

* Kernel CloudFormation template on s3://convox/release/20150902130909/formation.json
* Kernel Lambda binary on s3://convox/release/20150902130909/formation.zip
* `app`, `build`, `kernel`, `registry`, and `service` images tagged with "20150902130909" on Docker Hub
* A kernel CloudFormation template on s3://convox/release/20150902130909/formation.json
* A kernel Lambda binary on s3://convox/release/20150902130909/formation.zip
* `app`, `build`, `kernel`, `registry`, and `service` images tagged with "20150902130909" on Docker Hub

# CLI Artifacts

It is currently not possible / easy to get a versioned CLI artifact. Via Equinox we have:

* CLI binary on https://www.convox.com/downloads/{arch}/convox.zip

# Building Artifacts

At any time you can build artifacts:

```bash
$ convox switch demo.convox.io
$ convox run --app release release cli
$ convox run --app release release kernel
```

**Note the lack of flags.**

Wait ~5m to see a "20150901101300 (published: false, required: false)" message in Slack.

Wait ~20m for https://circleci.com/gh/convox/ci/tree/master to run to "success".

If CI is not green, you are not done.

* Trigger a rebuild to test if the error was transient
* Browse the CircleCI artifacts to determine a root cause

Now a developer can perform manual testing with:

```bash
$ convox switch convox-862259866.us-east-1.elb.amazonaws.com
$ convox system update 20150901101300
```

Or:

```bash
$ VERSION=20150901101300 convox install
```

# Publishing Artifacts

When you are confident that the artifacts will be an awesome experience for users, "publish" them:

```bash
$ convox switch demo.convox.io
$ convox run --app release release version -publish 20150901101300
```

**Note the explicit `-publish` flag.**

Now users that issue a `convox install` or `convox rack update` will get this latest published version.

# Requiring Artifacts

Some CloudFormation templates are required for migration purposes. To denote a "required" template:

```bash
$ convox switch demo.convox.io
$ convox run --app release release version -publish -require 20150901101300
```

**Note the explicit `-publish` and `-require` flag.**

# Building a Large Experimental System-Wide Change

Changes that span repos, even if just CLI and Kernel, are much harder to verify. A great strategy
is to push all changes to test together with the same branch name, then build and test these artifacts
together.

* Pick a feature branch name, i.e. "multi-region"
* Push "multi-region" branches to GitHub for every component you want to test together, i.e. `cli`, `kernel` and `app`

**Note if the release and CI scripts do not see a "multi-region" branch they will use "master"**

```bash
$ convox switch demo.convox.io
$ convox run --app release release kernel --branch multi-region
```

**Note the `--branch` flag.**

Wait ~5m to see a "kernel released: 20150905161552-multi-region (published: false, required: false)" message in Slack.

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
