# mann
## Your personal man pages

A command-line interface tool for storing customized commands and descriptions.

### Install

```bash
$ go get github.com/patmigliaccio/mann
```

### Usage

#### Retrieve

```bash
$ mann service

#   Name: service
#
#   Commands:
#        service --status-all
#        service --status-all | grep -E 'httpd'
```


#### Add


```bash
$ mann add service --status-all

# Added: service --status-all
```

Using quotes allows for the addition of more complex commands such as pipes.

```bash
$ mann add "service --status-all | grep -E 'httpd'"

# Added: service --status-all | grep -E 'httpd'
```