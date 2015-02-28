# Hoi

An easy file/message transfer tool using http file server.

## Installation

### Developer

```sh
$ go get github.com/monochromegane/hoi/...
```

### User

Download from the following url.

- [https://github.com/monochromegane/hoi/releases](https://github.com/monochromegane/hoi/releases)


## Usage

### Send file

You can build a download url.

```sh
$ hoi hoge.txt
http://192.168.0.100:8081/tx1lrmkkqenh5fy0izbassidb9t9l1na/hoge.txt
```

`hoi` links the file to `~/.hoi/temp_public/random`, and it's document root directory for `hoi server`.

### Send message

```sh
$ hoi message1 message2
http://192.168.0.100:8081/tx1lrmkkqenh5fy0izbassidb9t9l1na/message.txt
```
`hoi` create a message file to `~/.hoi/temp_public/random`, and it's document root directory for `hoi server`.

### Notify

You can notify the url to `Slack` account.

```sh
$ hoi file|message @user
http://192.168.0.100:8081/tx1lrmkkqenh5fy0izbassidb9t9l1na/message.txt
Message sent successfully to @user
```

Hoi supports [Slack API](https://api.slack.com/) and [takosan (a simple web interface to Slack)](https://github.com/kentaro/takosan).

### Configuration

`hoi` allows to change port number for `hoi server`

`~/.hoi/conf.json`

```json
{
  "port": 8082
}
```

if you want to use notify for Slack:

```json
{
  "notification": {
    "from":  "YOUR SLACK ACCOUNT",
    "to":    "slack",
    "token": "YOUR SLACK API TOKEN"
  }
}
```

Or takosan:

```json
{
  "notification": {
    "from": "YOUR SLACK ACCOUNT",
    "to":   "takosan",
    "host": "TAKOSAN HOST NAME",
    "port": TAKOSAN PORT NUMBER,
  }
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

