# convox/release

Convox release management.

To release via Slack:

    /release create [branch] | publish <version>

    /release create my-branch
    /release create
    /release publish 20151123210132


To release via Convox:

    $ convox run --app release release create my-branch
    $ convox run --app release release create
    $ convox run --app release release publish 20151123210132


The release workflow is:

* Create a Pull Request on Rack with a patch
* Create a release with the branch name
* `convox rack update 20151123210132-my-branch` to test the release
* Merge the pull request into master
* Release Rack from master
* Publish the resulting release number

See [Releases.md](Releases.md) for full release management guidelines.

## Development

> You will need a `.env` with AWS, Circle, Docker, Equinox and Slack keys.


```
$ make build

$ docker run --env-file=.env convox/release cli
$ docker run --env-file=.env convox/release kernel
```

## Production

#### Cut a release of the CLI

```
$ convox run --app release release cli
```

#### Cut a release of the API

```
$ convox run --app release release kernel
```

#### Promote a release of the API to stable

```
$ convox run --app release release version -publish update $VERSION
```

## License

Apache 2.0 &copy; 2015 Convox, Inc.
