package helpers

import (
	"fmt"
	"path/filepath"
)

// Include, belirtilen yolu temsil eden HTML şablon dosyalarını içeren bir dilim döndürür.
// Bu işlev, belirli bir dizindeki tüm HTML şablon dosyalarını birleştirir.
func Include(path string) []string {
	// admin/views/templates dizinindeki tüm .html dosyalarını alır.
	files, err := filepath.Glob("admin/views/templates/*.html")
	if err != nil {
		fmt.Println(err)
	}

	// Belirtilen yoldaki tüm .html dosyalarını alır.
	path_files, _ := filepath.Glob("admin/views/" + path + "/*.html")

	// path_files dilimini files dilimine ekler.
	for _, file := range path_files {
		files = append(files, file)
	}

	return files
	//fmt.Println(files)
}
