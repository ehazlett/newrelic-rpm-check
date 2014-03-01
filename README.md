# newrelic-rpm-check
This is a simple Go application that will check the NewRelic Requests Per Minute metric for applications and return the results.

# Build
`make`

# Arguments

* `-a`: NewRelic Application ID
* `-k`: NewRelic API Key
* `-t`: RPM Threshold (under this is considered critical)
* `-q`: Only show the `STATUS:VALUE`

# Usage
`newrelic-rpm-check -a <app-id> -k <api-key> -t <threshold> [-q]`

# Examples

```
$> newrelic-rpm-check -a 12345 -k abcdefg123456790ghijklmnop -t 1000
OK: Sample App Throughput: 1250 rpm
```

```
$> newrelic-rpm-check -a 12345 -k abcdefg123456790ghijklmnop -t 1000 -q
OK:1250
```
