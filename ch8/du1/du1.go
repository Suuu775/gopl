package du1

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

func dirents(dir string) []os.DirEntry {
    entries, err := os.ReadDir(dir)
    if err != nil {
        fmt.Fprintf(os.Stderr, "du1: %v\n", err)
        return nil
    }
    return entries
}

func walkDir(dir string,fileSizes chan<-int64){
	for _,entry := range(dirents(dir)){
		if entry.IsDir(){
			subdir := filepath.Join(dir, entry.Name())
			walkDir(subdir,fileSizes)
		} else {
			entryinfo,err  := entry.Info()
			if err  != nil {
				return
			}
			fileSizes<-entryinfo.Size()
		}
	}
}

func Du1(){
	flag.Parse()
	roots := flag.Args()
	if len(roots) ==0{
		roots = []string{"."}
	}
	fileSizes := make(chan int64)
	go func(){
		for _,root := range roots{
			walkDir(root,fileSizes)
		}
		close(fileSizes)
	}()
	var nfiles,nbytes int64
	for size:= range fileSizes{
		nfiles++
		nbytes += size
	}
	printDiskUsage(nfiles,nbytes)
}

func printDiskUsage(nfiles,nbytes int64){
	fmt.Printf("%d files  %.1f GB\n", nfiles, float64(nbytes)/1e9)}
