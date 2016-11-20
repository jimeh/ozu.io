# ozu.io - A drunken, shitty URL shortener

![build-status](https://api.travis-ci.org/jimeh/ozu.io.svg)

What do you do when you wanna mess with a new programming language? You make a
URL shortener, duh!

## Usage

### via CLI

```
$ ozuio --help
usage: ozuio [<flags>]

Flags:
      --help                  Show context-sensitive help (also try --help-long and --help-man).
  -p, --port="8080"           Port to listen to.
  -b, --bind="0.0.0.0"        Bind address.
  -d, --dir="ozuio_database"  Directory to store database file.
  -v, --version               Print version info.
```

### via Docker

```
$ docker run -d --name ozu.io -p 8080:8080 -v /ozuio_data:/data jimeh/ozu.io
```
