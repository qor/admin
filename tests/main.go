package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/TykTechnologies/raava-admin/tests/dummy"
	"github.com/qor/qor/utils"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	fmt.Printf("Listening on: %s\n", port)

	mux := http.NewServeMux()
	mux.Handle("/system/", utils.FileServer(http.Dir("public")))
	dummy.NewDummyAdmin(true).MountTo("/admin", mux)
	http.ListenAndServe(fmt.Sprintf(":%s", port), mux)
}
