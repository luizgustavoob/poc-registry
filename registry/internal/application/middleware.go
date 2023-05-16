package application

import (
	"context"
	"errors"
	"net/http"
	"strings"

	"github.com/luizgustavoob/registry/internal/consts"
	"github.com/luizgustavoob/registry/internal/entities"
)

type remoteServiceFinder interface {
	FindService(serviceName string) (entities.RemoteService, error)
}

func remoteServiceMiddleware(finder remoteServiceFinder) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()

			if strings.HasPrefix(r.URL.Path, "/commands") {
				processName := r.Header.Get("x-process-name")
				if processName == "" {
					internalHandleError(w, errors.New("header x-process-name not informed"), http.StatusBadRequest)
					return
				}

				remoteService, err := finder.FindService(processName)
				if err != nil {
					internalHandleError(w, err, http.StatusInternalServerError)
					return
				}

				ctx = context.WithValue(r.Context(), consts.RemoteServiceKey, remoteService)
			}

			next.ServeHTTP(w, r.Clone(ctx))
		})
	}
}
