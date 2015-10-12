# convox/release

Convox release management

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
