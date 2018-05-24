[![Mozilla Public License](https://img.shields.io/badge/license-MPL-blue.svg)](https://www.mozilla.org/MPL)
[![Go Report Card](https://goreportcard.com/badge/github.com/leprechau/ipman)](https://goreportcard.com/report/github.com/leprechau/ipman)

# ipman

## Summary

IPman is a simple tool to automatically update DNS records (A and AAAA) based on the external local IPv4 and/or IPv6 address of the local machine.  Currently it uses myexternalip.com for the external address lookup and supports writing records to GoDaddy's DNS API.  Both backend providers are modeled as interfaces to allow adding additional backends as needed in the future.

## Installing

Users with a proper Go environment (1.8+ required) ...

```
go get -u github.com/leprechau/ipman
```

Developers that wish to take advantage of vendoring and other options ...

```
git clone https://github.com/leprechau/ipman.git
cd ipman
make
```

## Usage

### Summary

```
ahurt$ ipman --help
Usage: ipman [--version] [--help] <command> [<args>]

Available commands are:
    check     Return current external ip address of local machine.
    update    Update DNS registry with external ip address of local machine.
```

### Check Options

| Option | Description |
|--------|-------------|
| `4`   | Get external IPv4 address if available.
| `6`   | Get external IPv6 address if available.

### Update Options

| Option    | Description |
|-----------|-------------|
| `4`      | Get external IPv4 address if available.
| `6`      | Get external IPv6 address if available.
| `key`    | The DNS API access key.  This may also be specified by setting the `IPMAN_DNS_KEY` environment variable.
| `secret` | The DNS API access secret.  This may also be specified by setting the `IPMAN_DNS_SECRET` environment variable.
| `domain` | The DNS domain name. This value defaults to the domain portion of the local hostname.
| `name`   | The DNS record name. This defaults to the domain root alias ("@").
| `ttl`    | The DNS record ttl in seconds.  This defaults to 600 seconds (5 minutes).

### Example

```
ahurt$ ./ipman update -4 -6 -key=YourAccessKey -secret=YourSuperSecretKey
2018/01/18 10:12:35 [IPv4] local/remote 67.187.109.252/67.187.109.252
2018/01/18 10:12:37 [IPv6] local/remote 2601:484:c000:5203:ec4:7aff:feb0:4068/2601:484:c000:5203:ec4:7aff:feb0:4068
```

The `key` and `secret` flags are optional when using environment variables.  This will keep your keys from potentially showing up in the system process list.  If Updates are required, the local address differs from the remote DNS record, the performed update will be printed.

### Crontab (automatic updates)

To dynamically update your records in near real time you can crontab the utillity using something similar to the below in your local users crontab.

```
## ipman secrets for godaddy api
IPMAN_DNS_KEY=YourAccessKey
IPMAN_DNS_SECRET=YourSuperSecretKey

## check and possibly update dns entries every 30 minutes
*/30 * * * * /path/to/ipman update -4 -6 >/dev/null 2>&1
```

Please see the crontab documentation of your local system for more information.

