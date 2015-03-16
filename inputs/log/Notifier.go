package main

import (
	"fmt"
	"gopkg.in/fsnotify.v1"
	"os"
	"strings"
)

var watcher *fsnotify.Watcher
var basePath string

type Notifier struct {
	nc     chan string
	folder string
}

func NewNotifier(nc chan string, folder string) *Notifier {
	return &Notifier{nc, folder}
}

func getWatcher() *fsnotify.Watcher {
	fswatcher, err := fsnotify.NewWatcher()
	if err != nil {
		return nil
	}
	return fswatcher
}

func (notifier *Notifier) Watch() {
	basePath = notifier.folder
	watcher = getWatcher()
	defer func() {
		watcher.Close()
		str := recover()
		if str != nil {
			fmt.Println(str)
		}
	}()
	fmt.Println("Watching folder ", basePath)
	watcher.Add(basePath)
	for {
		defer func() {
			str := recover()
			fmt.Println(str)
		}()
		select {
		case event := <-watcher.Events:
			if event.Op == fsnotify.Create {
				fi, _ := os.Stat(event.Name)
				if fi.Mode().IsDir() {
					fmt.Println(fi.Name() + " directory created. Adding to watcher...")
					fmt.Println("Event Name: " + event.Name)
					watcher.Add(event.Name)
				} else if strings.HasPrefix(fi.Name(), ".") == false {
					fmt.Println("Notified of the new file: " + fi.Name())
					notifier.nc <- event.Name
				}
			} else if event.Op == fsnotify.Remove {
				fmt.Println("Directory ", event.Name, " Removed...Removing from watcher")
				watcher.Remove(event.Name)
			}
		case err := <-watcher.Errors:
			fmt.Println("error:", err)
		}
	}

}

/*func main() {
	var input string
	watch("/home/suresh/go-projects/file")
	fmt.Scanln(&input)
	fmt.Println("Exit")
}*/
