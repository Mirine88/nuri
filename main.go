package nuri

// version 1.0.0

import (
	"mime/multipart"
	"net/http"
)

type methodType string

type HandlerFuncType func(c Context) (int, string)

// Context context
type Context struct {
	// Request is the same with *http.Request
	Request *http.Request
	// ToText makes html to text/plain
	ToText func(int, string) (int, string)
	// SetHeader sets the header value corresponding to the key
	SetHeader func(string, string)
	// GetFormValue gets the form value corresponding to the key
	GetFormValue func(string) string
	// GetFormFile gets the form file corresponding to the key
	GetFormFile func(string) (multipart.File, *multipart.FileHeader, error)
}

var (
	handlers                        = []handler{}
	notFoundHandler handlerFuncType = func(Context) (int, string) {
		return http.StatusNotFound, "404 Not Found"
	}
)

// methods
var (
	// MethodGET GET
	MethodGET methodType = "GET"
	// MethodHEAD HEAD
	MethodHEAD methodType = "HEAD"
	// MethodPOST POST
	MethodPOST methodType = "POST"
	// MethodPUT PUT
	MethodPUT methodType = "PUT"
	// MethodDELETE DELETE
	MethodDELETE methodType = "DELETE"
	// MethodCONNECT CONNECT
	MethodCONNECT methodType = "CONNECT"
	// MethodOPTIONS OPTIONS
	MethodOPTIONS methodType = "OPTIONS"
	// MethodTRACE TRACE
	MethodTRACE methodType = "TRACE"
	// MethodPATCH PATCH
	MethodPATCH methodType = "PATCH"
)

type handler struct {
	path    string
	method  methodType
	handler handlerFuncType
}

func handlerMiddleware(path string, method methodType, h handlerFuncType) {
	handlers = append(handlers, handler{
		path:    path,
		method:  method,
		handler: h,
	})
}

// Set404 shows a page that shows what you programmed when the path is not declared
func Set404(h handlerFuncType) {
	notFoundHandler = h
}

// SetNotFoundHandler shows a page that shows what you programmed when the path is not declared
// It the same with Set404 function.
func SetNotFoundHandler(h handlerFuncType) {
	Set404(h)
}

// GET handles a page when the method is GET
func GET(path string, h handlerFuncType) {
	handlerMiddleware(path, MethodGET, h)
}

// HEAD handles a page when the method is HEAD
func HEAD(path string, h handlerFuncType) {
	handlerMiddleware(path, MethodHEAD, h)
}

// POST handles a page when the method is POST
func POST(path string, h handlerFuncType) {
	handlerMiddleware(path, MethodPOST, h)
}

// PUT handles a page when the method is PUT
func PUT(path string, h handlerFuncType) {
	handlerMiddleware(path, MethodPUT, h)
}

// DELETE handles a page when the method is DELETE
func DELETE(path string, h handlerFuncType) {
	handlerMiddleware(path, MethodDELETE, h)
}

// CONNECT handles a page when the method is CONNECT
func CONNECT(path string, h handlerFuncType) {
	handlerMiddleware(path, MethodCONNECT, h)
}

// OPTIONS handles a page when the method is OPTIONS
func OPTIONS(path string, h handlerFuncType) {
	handlerMiddleware(path, MethodOPTIONS, h)
}

// TRACE handles a page when the method is TRACE
func TRACE(path string, h handlerFuncType) {
	handlerMiddleware(path, MethodTRACE, h)
}

// PATCH handles a page when the method is PATCH
func PATCH(path string, h handlerFuncType) {
	handlerMiddleware(path, MethodTRACE, h)
}

// Run runs pages that you make
func Run(addr string) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// context
		c := Context{
			Request: r,
			ToText: func(statusCode int, str string) (int, string) {
				w.Header().Add("Content-Type", "text/plain")
				return statusCode, str
			},
			SetHeader: func(key, value string) {
				w.Header().Add(key, value)
			},
			GetFormValue: func(key string) string {
				return r.FormValue(key)
			},
			GetFormFile: func(key string) (multipart.File, *multipart.FileHeader, error) {
				return r.FormFile(key)
			},
		}

		show404 := true
		for _, h := range handlers {
			if h.path == r.URL.Path && h.method == methodType(r.Method) {
				show404 = false

				statusCode, result := h.handler(c)
				w.WriteHeader(statusCode)
				w.Write([]byte(result))
			}
		}

		if show404 {
			_, result := notFoundHandler(c)
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(result))
		}
	})
	http.ListenAndServe(addr, nil)
}
