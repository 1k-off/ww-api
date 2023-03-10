# Web Watch API

## Config file

### Example

```yaml
database:
  provider: mongodb
  connection_string: mongodb://root:password123@localhost:27017/ww?authSource=admin
server:
  port: 8080
  access_token:
    private_key: LS0tLS1CRUdJTiBSU0EgUFJJVkFURSBLRVktLS0tLQpNSUlDV3dJQkFBS0JnUUNTVGcyWTJNUHc1Uys3M3RJWW5JRXBSZytxMm4zajVVNVlML2EwRndpTFFacGVUdXpwCi9ka0NTU280REgrSXU2eHR5Z1RaczNDWTJJQVhsZEVaQ2VnUGwxWUNvT1F4aVJqUXRDbUR6elBGZW1ocy9tSTYKUnZVaTlvT1BSSGdvVHY0eW0vNkpOOFhUV1Y5bmpBYm9BMTlSZGY0eXFiSTJKTUlkR21Dc3U1Z3ZZUUlEQVFBQgpBb0dBRnkzVWk2VzBEV05TRjdLcW1JbWJFUDN2bDRMOE5QSHNZcDVldUpONW9BNXB0WTFOalpqUkc4S1p0TDJPCkt3eUl1ZkFxcW83NlJNendKa3h2RDBkNWcxakkrV3VqeDBnSXlFMUlwUUg4RnhBa2YxTDlHK2pqUm9lMWNsazkKcC9IQTcwMEYvZnIxRzh2WCtkejlCdzBSc2RVRkNId1luWWExL1VIZy9EKzY5SzBDUVFEaDQ1cFJZdGtkdlM4RApBd3lPVVNsc1pVcEcxQXRadFNWK2RYMGU3ZHVKa3FmYkgvUitBQ1h3SExPN2VxWjRxcTZhTmIvcjY1U3ZhRllxClM0UTk4U1VqQWtFQXBjNnArbW5nQjc5VXR4NUY0SlBoLzBkcm5LVTIxMkEwYlVmdnRkbVFPQkp2WEtsWlo4UGkKTXo3ZEErSTZ2QVh6KzVlWHhWNE04dnU5dUVmbGU5S3Jxd0pBZkpKMlVoZS9RS1ZLUGRENnBhbWd2SVNIbDlQcwpob1pkclFYQ0FNS1A1YWlaSlVEVUpvQ1NhMzZJcUFXVnRNbjhERk5FQ2lrYkVEanIrOXMxakt0bUhRSkFUamJkCjRmMTlxOG5xcVhNRFhYd0ZHTW5WRHBDMC9RWXAxUDhoS2JSV25zeTdjWWVGWURoOEJOWjdwYkJiS29UWVllOVIKcmMyKzZBUXVxN1ptbjNGeWZ3SkFMVDViS2lXbEJ1c1NoM2NucEFMc3BCbk5PdFZIOVJkSFgrZ2Q3S1I2Q3k0Twp3RlJWREh0OExMNlRzLzhxRlEzeEcrcjVheHRjcHMzcUF6Ryt6bnlZTkE9PQotLS0tLUVORCBSU0EgUFJJVkFURSBLRVktLS0tLQ==
    public_key: LS0tLS1CRUdJTiBQVUJMSUMgS0VZLS0tLS0KTUlHZk1BMEdDU3FHU0liM0RRRUJBUVVBQTRHTkFEQ0JpUUtCZ1FDU1RnMlkyTVB3NVMrNzN0SVluSUVwUmcrcQoybjNqNVU1WUwvYTBGd2lMUVpwZVR1enAvZGtDU1NvNERIK0l1Nnh0eWdUWnMzQ1kySUFYbGRFWkNlZ1BsMVlDCm9PUXhpUmpRdENtRHp6UEZlbWhzL21JNlJ2VWk5b09QUkhnb1R2NHltLzZKTjhYVFdWOW5qQWJvQTE5UmRmNHkKcWJJMkpNSWRHbUNzdTVndllRSURBUUFCCi0tLS0tRU5EIFBVQkxJQyBLRVktLS0tLQ==
    expires_in: 900
  refresh_token:
    private_key: LS0tLS1CRUdJTiBSU0EgUFJJVkFURSBLRVktLS0tLQpNSUlDWEFJQkFBS0JnUUNGUFhZYWNSVFU3d0lJTWhhTmgvbzcvaWhhUSsxK09ETUxGVlVUdDA3WG83L285K3ArCm5jZTFCVkJZWEpTRG5qaGtTY3h2ZUlYK0ZaZnk1Ti8rQldsMDg1bFZLcGZRQ3NzN3RQZ0NNNTVSNnE4SU9wT0QKeEZqcnQ3c3c0N0VuNEN5RWxXSmZnSnV2dGJRS0t3OUFhY0RwTHBIOVR1bjBHb25kcll1TnRzNlRyUUlEQVFBQgpBb0dBSUt6WjhkSVprYjZyZU9jUHNOWFFBRFpzSHZiWm4wS1pBZmJSVG14OTRWUS9GQnI0WHVUQ3ZSbjRnakx5CkdQVU1UMHZwM3N4bno4L3IyNTFWY1M1U2xnV2dkYmYwUVJpYjZkYUdPcG9md0VkellVZUd1M2tIZ2xFSU00VUgKcDRaWlJoVnFWdDcrWjVOOUFpOFpEOUFNNUY1YnFsUXhIY2J3S1hqL0JCc2VpdUVDUVFEMFMvVytybWU1Zi9kbgptZVJ6V1hPTTBNMFpFZWlidEltSllQWHB5S1NZT2Y3clFVdkJlQ0Z4aEwwZUU4eEhlZDB6dTJPZS9lMHBYSEt1CmFuYnFDWS9iQWtFQWk1K0NkUTRXOFNKemllVHFaSHdCakpOSzRxOU44U29BTG5DVnYzOE8rTzRlZUV4MkdVWGIKTTZ4NGtMQ0xUczFJdGVaajZsc2NtbzluQnlZN2R2Z2xGd0pBZmhOTmlkQzhHeEdkZnN2L0NFQ2J6NHBhcnB0OQpiZzNvQlF3VEhVbTlHQXFtTW9jS0w1aHR1Z1lGQzZhd0FCczJPMkp6OFRPZTEzK3NkN0xlRjc1RGpRSkFXdUtGCjVQUTY4dFBlS1pDZEVyTzF2bS9TZXlScHMrWUhJRE9oQm5vS29QYy9Wa2RQU0x3MEo3ckk4RVk3S0J4d2pCZGcKU3Bqc1VaK3ZEUFJTR09zR0dRSkJBTHYwUjdzUlBNTXorVE85Tkk3Z0FDSHJhSjRuS1VXaFJ2Sm5TWWQ3WUZKWApEMWg1eGJHM1JvNkM1bmNpTlg5YTZXdTBUV1BzZ1J6cUtiL3QyN0JJb0RvPQotLS0tLUVORCBSU0EgUFJJVkFURSBLRVktLS0tLQ==
    public_key: LS0tLS1CRUdJTiBQVUJMSUMgS0VZLS0tLS0KTUlHZk1BMEdDU3FHU0liM0RRRUJBUVVBQTRHTkFEQ0JpUUtCZ1FDRlBYWWFjUlRVN3dJSU1oYU5oL283L2loYQpRKzErT0RNTEZWVVR0MDdYbzcvbzkrcCtuY2UxQlZCWVhKU0RuamhrU2N4dmVJWCtGWmZ5NU4vK0JXbDA4NWxWCktwZlFDc3M3dFBnQ001NVI2cThJT3BPRHhGanJ0N3N3NDdFbjRDeUVsV0pmZ0p1dnRiUUtLdzlBYWNEcExwSDkKVHVuMEdvbmRyWXVOdHM2VHJRSURBUUFCCi0tLS0tRU5EIFBVQkxJQyBLRVktLS0tLQ==
    expires_in: 3600
queue:
  provider: memphis
  memphis:
    client_login: fd51c04d.4239.41cd.be48.481beeb4ea80
    client_token: memphis
    url: localhost
    ssl_targets_station_name: ssl-targets
    uptime_targets_station_name: uptime-targets
    domain_expiration_targets_station_name: domain-expiration-targets
    ssl_metrics_station_name: ssl-metrics
    uptime_metrics_station_name: uptime-metrics
    domain_expiration_metrics_station_name: domain-expiration-metrics

```
### Description
| Name                                                 | Description                                                                 | Required | Default                   |
|------------------------------------------------------|-----------------------------------------------------------------------------|----------|---------------------------|
| database.provider                                    | The database provider to use.                                               | No       | mongodb                   |
| database.connection_string                           | The connection string to use for the database.                              | Yes      | -                         |
| server.port                                          | The port to run the server on.                                              | No       | 8080                      |
| server.access_token.private_key                      | The private key to use for signing access tokens. Must be base64 encoded.   | Yes      | -                         |
| server.access_token.public_key                       | The public key to use for verifying access tokens. Must be base64 encoded.  | Yes      | -                         |
| server.access_token.expires_in                       | The number of seconds before an access token expires.                       | No       | 900                       |
| server.refresh_token.private_key                     | The private key to use for signing refresh tokens. Must be base64 encoded.  | Yes      | -                         |
| server.refresh_token.public_key                      | The public key to use for verifying refresh tokens. Must be base64 encoded. | Yes      | -                         |
| server.refresh_token.expires_in                      | The number of seconds before a refresh token expires.                       | No       | 3600                      |
| queue.provider                                       | The queue provider to use.                                                  | No       | memphis                   |
| queue.memphis.client_login                           | The client login to use for the memphis queue.                              | Yes      | -                         |
| queue.memphis.client_token                           | The client token to use for the memphis queue.                              | Yes      | -                         |
| queue.memphis.url                                    | The url to use for the memphis queue.                                       | Yes      | -                         |
| queue.memphis.ssl_targets_station_name               | The name of the station to use for ssl targets.                             | No       | ssl-targets               |
| queue.memphis.uptime_targets_station_name            | The name of the station to use for uptime targets.                          | No       | uptime-targets            |
| queue.memphis.domain_expiration_targets_station_name | The name of the station to use for domain expiration targets.               | No       | domain-expiration-targets |
| queue.memphis.ssl_metrics_station_name               | The name of the station to use for ssl metrics.                             | No       | ssl-metrics               |
| queue.memphis.uptime_metrics_station_name            | The name of the station to use for uptime metrics.                          | No       | uptime-metrics            |
| queue.memphis.domain_expiration_metrics_station_name | The name of the station to use for domain expiration metrics.               | No       | domain-expiration-metrics |

## Memphis Queue

### Schemas

SSL target schema:
```json
{
  "url": ""
}
```
SSL metric schema:
```json
{
    "metadata": {
        "url": "",
        "location": ""
    },
    "timestamp": "",
    "expirationDate": "",
    "certData": {
        "host": "",
        "commonName": "",
        "alternativeNames": [
            ""
        ],
        "issuer": "",
        "validFrom": "",
        "validTo": ""
    },
    "error": "",
    "expiringSoon": false
}
```