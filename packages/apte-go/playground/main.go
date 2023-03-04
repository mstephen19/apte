package main

import (
	"fmt"
	"time"

	"github.com/mstephen19/apte"
	"github.com/mstephen19/apte/lib"
)

func main() {
	server := apte.NewEventServer(&apte.EventServerOptions{
		Endpoint: "/events",
	})
	messages := server.Namespace("messages")

	messages.Stream(func(ctx *lib.StreamHandlerContext) {
		ticker := time.NewTicker(time.Second)

		stream := ctx.Continuous(func() bool {
			<-ticker.C
			ctx.Dispatch("msg", []byte("hello world"))
			return true
		})

		<-stream
	})

	messages.Receive("msg", func(ctx *lib.ReceiveHandlerContext) {
		data := &map[string]any{}
		ctx.GetJSONData(data)
		fmt.Println(data)
	})

	fmt.Println("Server starting...")
	server.Start(apte.DefaultStartConfig())
}
