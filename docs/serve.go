package docs

import (
	"embed"
	"net/http"
)

var (
	//go:embed swagger/*
	docsFS embed.FS

	pages map[string]string
)


func DocsHandler() http.Handler {
  return http.FileServer(http.FS(docsFS))
}


// func init() {
// 	pages = make(map[string]string)
// 	docs, err := embed.FS(docsFS).ReadDir("swagger")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	for _, page := range docs {
// 		pages["/swagger/"+page.Name()] = "swagger/" + page.Name()
// 	}
// 	pages["/swagger/"] = "swagger/index.html"
// }

// func DocsHandler() http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		page, ok := pages[r.URL.Path]
// 		if !ok {
// 			w.WriteHeader(http.StatusNotFound)
// 			return
// 		}
//
// 		file, err := docsFS.Open(page)
// 		defer file.Close()
// 		if err != nil {
// 			log.Fatal(err)
// 		}
//
// 		fileInfo, err := file.Stat()
// 		if err != nil {
// 			http.Error(w, "Unable to get file information", http.StatusInternalServerError)
// 			return
// 		}
//
// 		// Set the Content-Type header based on the file's extension
// 		contentType := "application/octet-stream"
// 		switch {
// 		case len(fileInfo.Name()) > 2 && fileInfo.Name()[len(fileInfo.Name())-3:] == ".js":
// 			contentType = "application/javascript"
// 		case len(fileInfo.Name()) > 3 && fileInfo.Name()[len(fileInfo.Name())-4:] == ".css":
// 			contentType = "text/css"
// 		case len(fileInfo.Name()) > 4 && (fileInfo.Name()[len(fileInfo.Name())-5:] == ".html" || fileInfo.Name()[len(fileInfo.Name())-5:] == ".htm"):
// 			contentType = "text/html"
// 		}
//
// 		w.Header().Set("Content-Type", contentType)
//
// 		http.ServeContent(w, r, fileInfo.Name(), fileInfo.ModTime(), file.(io.ReadSeeker))
// 	})
// }
