[![Mozilla Public License](https://img.shields.io/badge/license-MPL-blue.svg)](https://www.mozilla.org/MPL)
[![Go Report Card](https://goreportcard.com/badge/github.com/leprechau/ipman)](https://goreportcard.com/report/github.com/leprechau/ipman)

# ipman

## Summary

IPman is a simple tool to automatically update DNS records (A and AAAA) based on the external local IPv4 and/or IPv6
address of the local machine.  It uses ipify.com for the external address lookup and supports writing
records to Cloudflare and GoDaddy's DNS API.  Both backend providers are modeled as interfaces to allow adding
additional backends as needed in the future.

Note: It appears GoDaddy has blocked API access for most customers. The GoDaddy API in this tool should
still function, but I am no longer able to test it and have moved my domain hosting to Cloudflare.
* https://www.reddit.com/r/selfhosted/comments/1cnipp3/warning_godaddy_silently_cut_access_to_their_dns/
* https://www.reddit.com/r/godaddy/comments/1bl0f5r/am_i_the_only_one_who_cant_use_the_api/

## Installing

Users with a proper Go environment (1.21+ required) ...

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

| Option | Description                                             |
|--------|---------------------------------------------------------|
| `4`    | Get external IPv4 address if available.                 |
| `6`    | Get external IPv6 address if available.                 |
| `ipbe` | IP lookup backend (`ipify` or `local`) default: `ipify` |


### Update Options

| Option   | Description                                                                              |
|----------|------------------------------------------------------------------------------------------|
| `4`      | Update external IPv4 address if available.                                               |
| `6`      | Update external IPv6 address if available.                                               |
| `key`    | The DNS API access key.  This defaults tp `$IPMAN_DNS_KEY` from the environment.         |
| `secret` | The DNS API access secret.  This defaults to `$IPMAN_DNS_SECRET` from the environment.   |
| `zone`   | The DNS zone ID or domain name. This defaults to `$IPMAN_DNS_ZONE` from the environment. |
| `name`   | The DNS record name. This defaults to the dns zone apex ("@").                           |
| `ttl`    | The DNS record ttl in seconds.  This defaults to 600 seconds (5 minutes).                |
| `ipbe`   | IP lookup backend (`ipify` or `local`). This defaults to `ipify`.                        |
| `dnsbe`  | DNS update backend (`cloudflare` or `godaddy`). This defaults to `cloudflare`.           |

### Example

```
ahurt$ ./ipman update -4 -6 -secret=<CF token or GD secret> -zone=<CF zone ID or GD domain name>
time=2024-07-12T13:19:49.872-05:00 level=INFO msg="local" iType=IPv4 addr=151.182.28.185
time=2024-07-12T13:19:50.538-05:00 level=INFO msg="remote" iType=IPv4 addr=151.182.28.180
time=2024-07-12T13:19:50.628-05:00 level=INFO msg="updated remote" record=anbcs.com rType=AAAA data=151.182.28.185
time=2024-07-12T13:19:50.731-05:00 level=INFO msg="local" iType=IPv6 addr=2601:1702:22d2:3c60:94c6:6adc:c162:dcd6
time=2024-07-12T13:19:51.146-05:00 level=INFO msg="remote" iType=IPv6 addr=2601:1702:22d2:3c60:3ecc:effe:fe22:4810
time=2024-07-12T13:19:51.146-05:00 level=INFO msg="updated remote" record=anbcs.com rType=AAAA data=2601:1702:22d2:3c60:94c6:6adc:c162:dcd6
```

The `key` and `secret` flags are optional when using environment variables.  This will keep your keys from potentially
showing up in the system process list.  Updates are only triggered when the local address differs from the remote record.

### Crontab (automatic updates)

To dynamically update your records in near real time you can crontab the utility using something similar to the below
in your local users crontab.

```
## ipman vars for cloudflare
export IPMAN_DNS_SECRET=<CF API Token>
export IPMAN_DNS_ZONE=<CF Zone ID>

## check and possibly update dns entries every 30 minutes
*/30 * * * * /path/to/ipman update -4 -6 >/dev/null 2>&1
```

Please see the crontab documentation of your local system for more information.

