# hoi

An easy file transfer tool using http file server.

## Installation

### Developer

```sh
$ go get github.com/monochromegane/hoi
```

### User

Download from the following url.

- [https://github.com/monochromegane/hoi/releases](https://github.com/monochromegane/hoi/releases)


## Usage

You can build download url.

```sh
$ hoi hoge.txt
http://192.168.0.100:8081/tx1lrmkkqenh5fy0izbassidb9t9l1na/hoge.txt
```

`hoi` links the file to `~/.hoi/temp_public/random`, and it's document root directory for `hoi server`.

### Configuration

`hoi` allows to change port number for `hoi server`

`~/.hoi/conf.json`

```json
{
  "port": 8082
}
```

### Clear public directory

```sh
$ hoi --clear
```

## Tasks

- Add `hoi server` commands. (start|stop|status)
- Add scheduler that clear public directory every some minutes.

## Code status

[![wercker status](https://app.wercker.com/status/69e506b638dc36b51678adce4cd215b6/m/master "wercker status")](https://app.wercker.com/project/bykey/69e506b638dc36b51678adce4cd215b6)

[![Build Status](https://travis-ci.org/monochromegane/hoi.svg?branch=master)](https://travis-ci.org/monochromegane/hoi)

## Contributing

1. Fork it
2. Create your feature branch (`git checkout -b my-new-feature`)
3. Commit your changes (`git commit -am 'Add some feature'`)
4. Push to the branch (`git push origin my-new-feature`)
5. Create new Pull Request

