# caddyv1
-------------------

`xgo -out "cadd" -buildmode=exe -pkg caddy -ldflags='-linkmode external -extldflags "-static"' -tags=release -targets "linux/arm" -branch v1 github.com/dbiesecke/caddyv1 `


`xgo -buildmode=exe -ldflags='-linkmode external -extldflags "-static"' -tags=release -targets "linux/amd64" github.com/dbiesecke/caddyv1  `

caddyv1
