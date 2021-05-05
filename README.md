# Mat Ryer's: How I Write HTTP Web Services after Eight Years

Compilable code for Mat Ryer's GopherCon 2019 talk, (How I Write HTTP Web Services after Eight
Years)[https://youtu.be/rWBSMsLG8po].

I took these notes while watching the talk. I obviously did not create this code, and Mat Ryer has
not, as far as I know, licensed it. For this reason, I will be more than happy to take this repo
private.

The fact that this code compiles is not that important since it does nothing interesting. But the
`ANCILLARY STUB` sections do show what types might satisfy some of Mat's slide code, and in addition
it makes the somewhat disconnected and incompatible slide code consistent.

Code discussion starts at (9:12)(https://youtu.be/rWBSMsLG8po?t=552)

## ToC/Timecodes

| Timecode                                     | Topic                                                              |
|----------------------------------------------|--------------------------------------------------------------------|
| (9:12)[https://youtu.be/rWBSMsLG8po?t=552]   | main and run <main.go>                                             |
| (10:04)[https://youtu.be/rWBSMsLG8po?t=604]  | server type <server.go>, <email/email.go>                          |
| (12:52)[https://youtu.be/rWBSMsLG8po?t=772]  | server as http.Handler <server.go>                                 |
| (11:12)[https://youtu.be/rWBSMsLG8po?t=672]  | newServer() <server.go>                                            |
| (14:02)[https://youtu.be/rWBSMsLG8po?t=842]  | routes file <routes.go>                                            |
| (15:13)[https://youtu.be/rWBSMsLG8po?t=913]  | handlers as server methods <handlermethods.go>                     |
| (16:59)[https://youtu.be/rWBSMsLG8po?t=1019] | handler-specific setup <handlermethods.go>                         |
| (18:00)[https://youtu.be/rWBSMsLG8po?t=1080] | handler method arguments <handlermethods.go>                       |
| (19:31)[https://youtu.be/rWBSMsLG8po?t=1171] | multiple server types <server.go>                                  |
| (21:54)[https://youtu.be/rWBSMsLG8po?t=1314] | middleware <middleware.go>                                         |
| (25:08)[https://youtu.be/rWBSMsLG8po?t=1508] | some common (premature?) abstractions <encodedecode.go>            |
| (27:52)[https://youtu.be/rWBSMsLG8po?t=1672] | inner types <handlermethods.go>                                    |
| (29:13)[https://youtu.be/rWBSMsLG8po?t=1753] | lazy setup with sync.Once <handlermethods.go>                      |
| (31:54)[https://youtu.be/rWBSMsLG8po?t=1914] | testing with (httptest)[https://golang.org/pkg/net/http/httptest/] |
| (33:31)[https://youtu.be/rWBSMsLG8po?t=2011] | testing servers <server_test.go>                                   |
| (37:09)[https://youtu.be/rWBSMsLG8po?t=2229] | testing inner types <server_test.go>                               |
| (38:44)[https://youtu.be/rWBSMsLG8po?t=2324] | integration vs. unit tests <server_test.go>                        |
| (40:08)[https://youtu.be/rWBSMsLG8po?t=2408] | e2e tests <server_test.go>                                         |
| (41:14)[https://youtu.be/rWBSMsLG8po?t=2474] | testing middleware <server_test.go>                                |
