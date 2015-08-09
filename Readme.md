# convox/release

Convox release management.

## Development

Pre-reqs

* An .env file with AWS, Circle, Docker, Equinox and Slack keys.

```bash
$ make build

$ docker run --env-file=.env convox/release cli
$ docker run --env-file=.env convox/release kernel
```

## Production

Deploy the app to Convox and set the env.

```
$ convox run --app release release cli
$ convox run --app release release kernel
```

When CI has passed, you can cut a "latest" release, which every fresh install will use.

```
$ convox run --app release release kernel --latest
```

## Contributing

* Open a [GitHub Issue](https://github.com/convox/release/issues/new) for bugs and feature requests
* Initiate a [GitHub Pull Request](https://help.github.com/articles/using-pull-requests/) for patches

## See Also

* [convox/app](https://github.com/convox/app)
* [convox/build](https://github.com/convox/build)
* [convox/cli](https://github.com/convox/cli)
* [convox/kernel](https://github.com/convox/kernel)

## License

Apache 2.0 &copy; 2015 Convox, Inc.
