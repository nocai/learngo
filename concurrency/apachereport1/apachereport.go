package apachereport1

import (
	"bufio"
	"io"
	"log"
	"os"
	"sync"
)

func readLines(filename string, lines chan<- string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal("failed to open the file : ", err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if line != "" {
			lines <- line
		}
		if err != nil {
			if err != io.EOF {
				log.Println("failed to finish reading the file:", err)
			}
			break
		}
	}
	close(lines)
}

//func processLines(done chan<- struct{}, pageMap )

type pageMap struct {
	countForPae map[string]int
	mutex       *sync.RWMutex
}

func NewPageMap() *pageMap {
	return &pageMap{make(map[string]int), new(sync.RWMutex)}
}

func (pm *pageMap) Increment(page string) {
	pm.mutex.Lock()
	defer pm.mutex.Unlock()
	pm.countForPae[page]++
}

func (pm *pageMap) Len() int {
	pm.mutex.RLock()
	defer pm.mutex.RUnlock()
	return len(pm.countForPae)
}
