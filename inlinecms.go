// Package inlinecms is the Inline CMS togo plugin: Turn any public site into its own CMS: schema-driven edit-in-place, admin bar, inline fields and pickers, reorder — no separate dashboard.
//
// It self-registers a provider on blank-import and mounts its routes onto the
// kernel. The concrete implementation is ported from the fadymondy.com app under
// internal/server — this scaffold wires the provider + a health route.
package inlinecms

import (
	"net/http"

	"github.com/togo-framework/togo"
)

func init() {
	togo.RegisterProviderFunc("inline-cms", togo.PriorityService, func(k *togo.Kernel) error {
		k.Router.Get("/api/inline-cms/health", func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(`{"plugin":"inline-cms","status":"ok"}`))
		})
		return nil
	})
}
