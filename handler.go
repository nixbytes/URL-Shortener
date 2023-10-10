package urlshort

import (
	"net/http"
)

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	// Return an anonymous function (http.HandlerFunc)
	return func(w http.ResponseWriter, r *http.Request) {
		//Extract the path from the incoming request URL
		path := r.URL.Path

		// Check if the path exists in the pathsToUrls map
		if dest, ok := pathsToUrls[path]; ok {

			// If it exists, perform a 302 (Found) HTTP redirect to the destination URL
			http.Redirect(w, r, dest, http.StatusFound)

			return // Exit the function to prevent further processing

		}

		// If the path is not found in the map, call the fallback handler
		fallback.ServeHTTP(w, r)
	}
}

// YAMLHandler will parse the provided YAML and then return
// an http.HandlerFunc (which also implements http.Handler)
// that will attempt to map any paths to their corresponding
// URL. If the path is not provided in the YAML, then the
// fallback http.Handler will be called instead.
//
// YAML is expected to be in the format:
//
//   - path: /some-path
//     url: https://www.some-url.com/demo
//
// The only errors that can be returned all related to having
// invalid YAML data.
//
// See MapHandler to create a similar http.HandlerFunc via
// a mapping of paths to urls.
func YAMLHandler(yamlbytes []byte, fallback http.Handler) (http.HandlerFunc, error) {
	// TODO: Implement this...
	return ""
}

type pathUrl struct {
	Path string `yaml:"path"`
}
