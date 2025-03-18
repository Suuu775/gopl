package ex89

import (
	"flag"
    "fmt"
	"os"
	"path/filepath"
	"sync"
	"time"
)

var sema = make(chan struct{}, 20)

type DirSize struct {
	Path string
	Size int64
}

func dirents(dir string) []os.DirEntry {
	sema <- struct{}{}
	defer func() { <-sema }()
	entries, err := os.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du3: %v\n", err)
		return nil
	}
	return entries
}

func walkDir(dir string, wg *sync.WaitGroup, dirSizes chan<- DirSize) {
	defer wg.Done()

	dirSize := int64(0)
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			wg.Add(1)
			subdir := filepath.Join(dir, entry.Name())
			go walkDir(subdir, wg, dirSizes)
		} else {
			entryInfo, err := entry.Info()
			if err != nil {
				return
			}
			dirSize += entryInfo.Size()
		}
	}
	dirSizes <- DirSize{Path: dir, Size: dirSize}
}

func Du() {
	var verbose = flag.Bool("v", false, "show me verbose progress message")
	var tick <-chan time.Time
	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}

	dirSizes := make(chan DirSize)

	var wg sync.WaitGroup
	for _, root := range roots {
		wg.Add(1)
		go walkDir(root, &wg, dirSizes)
	}

	go func() {
		wg.Wait()
		close(dirSizes)
	}()

	if *verbose {
		tick = time.Tick(500 * time.Millisecond)
	}

	var totalFiles, totalBytes int64
	directorySizes := make(map[string]int64)

loop:
	for {
		select {
		case sizeInfo, ok := <-dirSizes:
			if !ok {
				break loop
			}
			totalFiles++
			totalBytes += sizeInfo.Size
			directorySizes[sizeInfo.Path] = sizeInfo.Size
		case <-tick:
			printDirectoryUsage(directorySizes, totalFiles, totalBytes)
		}
	}
	printDirectoryUsage(directorySizes, totalFiles, totalBytes)
}

func printDirectoryUsage(directorySizes map[string]int64, totalFiles, totalBytes int64) {
	fmt.Println("\nDirectory Sizes:")
	for path, size := range directorySizes {
		fmt.Printf("%s: %.1f GB\n", path, float64(size)/1e9)
	}
	fmt.Printf("\nTotal Files: %d\nTotal Size: %.1f GB\n", totalFiles, float64(totalBytes)/1e9)
}
