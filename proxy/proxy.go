package proxy

import (
	"io"
	"net/http"

	"github.com/mjthecoder65/load-balancer/backend"
)

func ForwardRequest(backend *backend.Backend, writer http.ResponseWriter, request *http.Request) {
	url := backend.URL + request.RequestURI

	req, err := http.NewRequest(request.Method, url, request.Body)

	if err != nil {
		http.Error(writer, "Error creating request", http.StatusInternalServerError)
		return
	}

	req.Header = request.Header
	client := &http.Client{}
	res, err := client.Do(req)

	if err != nil {
		http.Error(writer, "Error forwarding request", http.StatusInternalServerError)
		return
	}

	defer res.Body.Close()

	for key, value := range res.Header {
		writer.Header()[key] = value
	}

	writer.WriteHeader(res.StatusCode)
	io.Copy(writer, res.Body)
}
