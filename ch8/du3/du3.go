package du3

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"
)

var sema = make(chan struct{},20)


func dirents(dir string) []os.DirEntry {
	sema <- struct{}{}
	defer func(){<-sema}()
    entries, err := os.ReadDir(dir)
    if err != nil {
        fmt.Fprintf(os.Stderr, "du1: %v\n", err)
        return nil
    }
    return entries
}

func walkDir(dir string,n *sync.WaitGroup,fileSizes chan<-int64){
	defer n.Done()
	for _,entry := range(dirents(dir)){
		if entry.IsDir(){
			n.Add(1)
			subdir := filepath.Join(dir, entry.Name())
			go walkDir(subdir,n,fileSizes)
		} else {
			entryinfo,err  := entry.Info()
			if err  != nil {
				return
			}
			fileSizes<-entryinfo.Size()
		}
	}
}

func Du3(){
	var verbose = flag.Bool("v",false,"show me verbose progress message")
	var tick <-chan time.Time
	flag.Parse()
	roots := flag.Args()
	if len(roots) ==0{
		roots = []string{"."}
	}
	fileSizes := make(chan int64)

	var n sync.WaitGroup
	for _,root := range roots{
		n.Add(1)
		go walkDir(root,&n,fileSizes)
	}
	go func(){
		n.Wait()
		close(fileSizes)
	}()

	if *verbose {
		tick = time.Tick(500*time.Millisecond)
	}

	var nfiles,nbytes int64
	loop:
	    for {
			select {
			case size,ok := <-fileSizes:
				if !ok {
					break loop
				}
				nfiles++
				nbytes+=size
            case <-tick:
                printDiskUsage(nfiles,nbytes)
			}
		}
	printDiskUsage(nfiles,nbytes)
}

func printDiskUsage(nfiles,nbytes int64){
	fmt.Printf("%d files  %.1f GB\n", nfiles, float64(nbytes)/1e9)}
