# Mat Ryer's: How I Write HTTP Web Services after Eight Years

Compilable code for Mat Ryer's GopherCon 2019 talk, [How I Write HTTP Web Services after Eight
Years](https://youtu.be/rWBSMsLG8po).

I took these notes while watching the talk. I obviously did not create this code, and Mat Ryer has
not, as far as I know, licensed it. For this reason, I will be more than happy to take this repo
private.

The fact that this code compiles is not that important since it does nothing interesting. But the
`ANCILLARY STUB` sections do show what types might satisfy some of Mat's slide code, and in addition
it makes the somewhat disconnected and incompatible slide code consistent.


## ToC/Timecodes

| Timecode                                     | Topic                  | File                                                                                     |
|----------------------------------------------|------------------------|------------------------------------------------------------------------------------------|
| [8:18](https://youtu.be/rWBSMsLG8po?t=8m18s) | code discussion starts |                                                                                          |
| [9:25](https://youtu.be/rWBSMsLG8po?t=9m25s) | main and run           | [main.go](https://github.com/AndrewLivingston/mr-http-service/blob/main/main.go#L11-L32) |
| [10:08](https://youtu.be/rWBSMsLG8po?t=10m8s)  | server type | [server.go](https://github.com/AndrewLivingston/mr-http-service/blob/main/server.go#L5-L13)
| [11:24](https://youtu.be/rWBSMsLG8po?t=11m24s)  | server constructor | [server.go](https://github.com/AndrewLivingston/mr-http-service/blob/main/server.go#L15-L25)                                            |
| [12:57](https://youtu.be/rWBSMsLG8po?t=12m57s)  | server as http.Handler | [server.go](https://github.com/AndrewLivingston/mr-http-service/blob/main/server.go#L27-L35)                                 |
| [14:11](https://youtu.be/rWBSMsLG8po?t=14m11s)  | routes file | [routes.go](https://github.com/AndrewLivingston/mr-http-service/blob/main/routes.go#L5-L14)                                            |
| [15:18](https://youtu.be/rWBSMsLG8po?t=15m18s)  | handlers as server methods | [handlermethods.go](https://github.com/AndrewLivingston/mr-http-service/blob/main/handlermethods.go#L11-L33)                     |
| [17:03](https://youtu.be/rWBSMsLG8po?t=17m3s) | handler-specific setup | [handlermethods.go](https://github.com/AndrewLivingston/mr-http-service/blob/main/handlermethods.go#L18-L25)                         |
| [18:08](https://youtu.be/rWBSMsLG8po?t=18m8s) | handler method arguments | [handlermethods.go](https://github.com/AndrewLivingston/mr-http-service/blob/main/handlermethods.go#L35-L54)                       |
| [19:39](https://youtu.be/rWBSMsLG8po?t=19m39s) | multiple server types | [server.go](https://github.com/AndrewLivingston/mr-http-service/blob/main/server.go#L37-L52)                                  |
| [20m24s](https://youtu.be/rWBSMsLG8po?t=20m24s) | HandlerFunc vs. Handler | (none) |
| [21:54](https://youtu.be/rWBSMsLG8po?t=21m54s) | middleware | [middleware.go](https://github.com/AndrewLivingston/mr-http-service/blob/main/middleware.go#L5-L28)                                         |
| [24:39](https://youtu.be/rWBSMsLG8po?t=24m39s) | some common (premature?) abstractions | [encodedecode.go](https://github.com/AndrewLivingston/mr-http-service/blob/main/encodedecode.go#L8-L36)            |
| [27:57](https://youtu.be/rWBSMsLG8po?t=27m57s) | inner types | [handlermethods.go](https://github.com/AndrewLivingston/mr-http-service/blob/main/handlermethods.go#L56-L72)                                    |
| [29:21](https://youtu.be/rWBSMsLG8po?t=29m21s) | lazy setup with sync.Once | [handlermethods.go](https://github.com/AndrewLivingston/mr-http-service/blob/main/handlermethods.go#L74-L96)                      |
| [31:56](https://youtu.be/rWBSMsLG8po?t=31m56s) | testing with [httptest](https://golang.org/pkg/net/http/httptest/) and [is](https://github.com/matryer/is) | |
| [33:43](https://youtu.be/rWBSMsLG8po?t=33m43s) | testing servers | [server_test.go](https://github.com/AndrewLivingston/mr-http-service/blob/main/server_test.go#L13-L27)                                   |
| [37:18](https://youtu.be/rWBSMsLG8po?t=37m18s) | testing inner types | [server_test.go](https://github.com/AndrewLivingston/mr-http-service/blob/main/server_test.go#L29-L46)                               |
| [38:46](https://youtu.be/rWBSMsLG8po?t=38m46s) | integration vs. unit tests | [server_test.go](https://github.com/AndrewLivingston/mr-http-service/blob/main/server_test.go#L52-L64)                        |
| [40:13](https://youtu.be/rWBSMsLG8po?t=40m13s) | e2e tests | [server_test.go](https://github.com/AndrewLivingston/mr-http-service/blob/main/server_test.go#L66-L85)                                         |
| [41:15](https://youtu.be/rWBSMsLG8po?t=41m15s) | testing middleware | [server_test.go](https://github.com/AndrewLivingston/mr-http-service/blob/main/server_test.go#L87-L115)                                |
