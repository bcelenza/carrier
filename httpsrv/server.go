package httpsrv

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"

	log "github.com/sirupsen/logrus"
)

const targetHeader = "X-Carrier-Target"

type HttpServer struct {
	Port               string
	TLSCertificateFile string
	TLSKeyFile         string
}

func New(port, cert, key string) *HttpServer {
	return &HttpServer{
		Port:               port,
		TLSCertificateFile: cert,
		TLSKeyFile:         key,
	}
}

func (h *HttpServer) Start() error {
	http.HandleFunc("/", h.handleRequest)
	http.HandleFunc("/ping", h.handlePing)
	if h.TLSCertificateFile != "" && h.TLSKeyFile != "" {
		log.Info("Certificate and key files provided, will serve TLS")
		return http.ListenAndServeTLS(fmt.Sprintf(":%s", h.Port), h.TLSCertificateFile, h.TLSKeyFile, nil)
	}
	log.Info("No certificate or key file provided, will serve plain text")
	return http.ListenAndServe(fmt.Sprintf(":%s", h.Port), nil)
}

func (h *HttpServer) handleRequest(res http.ResponseWriter, req *http.Request) {
	log.Info("Handling request: ", req)

	// Parse the target
	target := req.Header.Get(targetHeader)
	parsedTarget, _ := url.Parse(target)

	// Modify the request to the new target, then delete the
	// target header
	req.URL.Host = parsedTarget.Host
	req.URL.Scheme = parsedTarget.Scheme
	req.Host = parsedTarget.Host
	req.Header.Del(targetHeader)

	log.Debug("Upstream request: ", req)

	// Proxy the request
	proxy := httputil.NewSingleHostReverseProxy(parsedTarget)
	proxy.ServeHTTP(res, req)
}

func (h *HttpServer) handlePing(res http.ResponseWriter, req *http.Request) {
	res.WriteHeader(200)
}
