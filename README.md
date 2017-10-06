# wpcm
WebPage Change Monitor

Simple tool to monitor many webpages for a change.
Written in Go.
You can configure it from the `config.json`

```
{
    "timeInterval" : 33, //Time interval in minutes
    "cookie" : "", //Set custom cookie
    "urls" : [ //Array of URLs to monitor
        "https://example.com"
    ]
}
```
