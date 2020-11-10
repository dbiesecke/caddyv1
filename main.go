package main

import (
      "github.com/caddyserver/caddy/caddy/caddymain"

//      "github.com/mastercactapus/caddy-proxyprotocol"       //http.proxyprotocol    
//   _ "github.com/mahrud/caddy-altonions"  
    _ "github.com/hyperion-hyn/caddy-cron"
    _ "github.com/hacdias/caddy-minify"                     //http.minify
    _ "github.com/hacdias/caddy-webdav"                     //http.webdav
    _ "github.com/jung-kurt/caddy-cgi"                      //http.cgi
    _ "github.com/jung-kurt/caddy-pubsub"                   //http.pubsub
    _ "github.com/xuqingfeng/caddy-rate-limit"              //http.ratelimit
//      _ "go.okkur.org/gomods"                                 //http.gomods
        _ "go.okkur.org/torproxy"                               //http.torproxy  
//     _ "github.com/pieterlouw/caddy-net"
//      "github.com/caddyserver/caddy"
//      "github.com/caddyserver/caddy/caddyhttp/header"
//      "github.com/caddyserver/caddy/caddyhttp/httpserver"  
)

func main() {
  caddymain.Run()
}
