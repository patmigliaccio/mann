# mann
## Your personal man pages

A command-line interface tool for storing and running customized commands.

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
#   1.      service --status-all
#   2.      service --status-all | grep -E 'httpd|vsftpd'
```


#### Add

Store a command including any flags and predicates for reuse later.

```bash
$ mann add service --status-all

# Added: service --status-all
```

Using quotes allows for the addition of more complex commands such as pipes.

```bash
$ mann add "service --status-all | grep -E 'httpd|vsftpd'"

# Added: service --status-all | grep -E 'httpd|vsftpd'
```

#### Run

Execute a command by passing the list item number. 

```bash
$ mann run service 2

# service --status-all | grep -E 'httpd|vsftpd'
# httpd (pid  2301) is running...
# vsftpd (pid 14070 2061) is running...
```