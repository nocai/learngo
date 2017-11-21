package filter

import (
	"fmt"
	"path/filepath"
	"strings"
)

func source(files []string) <-chan string {
	out := make(chan string, 10000)
	go func() {
		for _, filename := range files {
			out <- filename
		}
		close(out)
	}()
	return out
}

func filterSuffixes(suffixes []string, in <-chan string) <-chan string {
	out := make(chan string, cap(in))
	go func() {
		for filename := range in {
			if len(suffixes) == 0 {
				out <- filename
				continue
			}
			ext := strings.ToLower(filepath.Ext(filename))
			for _, suffix := range suffixes {
				if ext == suffix {
					out <- filename
					break
				}
			}
		}
		close(out)
	}()
	return out
}

func filterSize(min, max int, in <-chan string) <-chan string {
	out := make(chan string, cap(in))
	go func() {
		for filename := range in {
			if min < len(filename) && len(filename) < max {
				out <- filename
			}
		}
		close(out)
	}()
	return out
}

func sink(in <-chan string) {
	for filename := range in {
		fmt.Println(filename)
	}
}
