# finch

A Golang Telegram Bot framework

Unlike the [Telegram Bot API](https://github.com/Syfaro/telegram-bot-api), this is a framework for writing commands, not just low level dealings with the API directly. 

It currently is in its early infancy and needs lots of work. Expect frequent breaking changes. 

You can see how to write some commands from the example commands provided in the `commands` folder. 

## Example

It's fairly easy to get this bot running, it requires few lines of code.

```go
package main

import (
	"github.com/syfaro/finch"
	_ "github.com/syfaro/finch/commands/help"
	_ "github.com/syfaro/finch/commands/info"
	_ "github.com/syfaro/finch/commands/stats"
)

func main() {
	f := finch.NewFinch("MyAwesomeBotToken")

	f.Start()
}
```

The webhook listener code is currently untested, and requires running a `net/http` server. 

```go
package main

import (
	"github.com/syfaro/finch"
	_ "github.com/syfaro/finch/commands/help"
	_ "github.com/syfaro/finch/commands/info"
	_ "github.com/syfaro/finch/commands/stats"
	"net/http"
)

func main() {
	f := finch.NewFinchWithClient("MyAwesomeBotToken", &http.Client{})

	f.StartWebhook()

	http.ListenAndServeTLS("0.0.0.0:8443", "cert.pem", "key.pem", nil)
}
```
