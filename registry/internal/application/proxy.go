package application

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/luizgustavoob/registry/internal/consts"
	"github.com/luizgustavoob/registry/internal/entities"
)

func proxy(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	remoteService := ctx.Value(consts.RemoteServiceKey).(entities.RemoteService)

	log.Printf("calling concrete address of process %s. (%s)", remoteService.ProcessName, remoteService.FinalAddress)

	url, _ := url.Parse(fmt.Sprintf("%s%s", remoteService.Address, r.URL.Path))

	proxy := httputil.NewSingleHostReverseProxy(url)
	proxy.Director = func(req *http.Request) {
		req.URL.Host = url.Host
		req.URL.Scheme = url.Scheme
		req.Header.Set("X-Forwarded-Host", req.Header.Get("Host"))
		req.Host = url.Host
	}
	proxy.ServeHTTP(w, r)
}
