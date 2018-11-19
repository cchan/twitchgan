package main

import (
  "net/http"
  "io"
  "os"
	"strconv"
	"log"
)

const minI int = 1
const maxI int = 1617254
const nWorkers int = 200

type WorkerStats struct {
	success int
	exists int
}

func downloadWorker(done chan<- WorkerStats, tasks <-chan int) {
	var stats WorkerStats

	for i := range tasks {
		func(i int) {
			stat, err := os.Stat("dataset/" + strconv.Itoa(i) + ".png")
			if err == nil && stat.Size() > 0 { stats.exists += 1; return }
			if err != nil && !os.IsNotExist(err) { log.Println("err(", i, "): os.Stat:", err); return }

			resp, err := http.Get("https://static-cdn.jtvnw.net/emoticons/v1/" + strconv.Itoa(i) + "/1.0")
			if err != nil { log.Println("err(", i, "): http.Get:", err); return }
			defer resp.Body.Close()
			if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNotFound { log.Println("err(", i, "): resp.Status:", resp.Status); return }

			out, err := os.Create("dataset/" + strconv.Itoa(i) + ".png")
			if err != nil { log.Println("err(", i, "): os.Create:", err); return }
			defer out.Close()

			_, err = io.Copy(out, resp.Body)
			if err != nil { log.Println("err(", i, "): io.Copy:", err); return }

			// log.Println("success(", i, ")")
			stats.success += 1
		}(i)
	}

	done <- stats
}

func main() {
	total := new(WorkerStats)
	done := make(chan WorkerStats)
	tasks := make(chan int)
	go func(){
		for i := minI; i <= maxI; i++ {
			tasks <- i
			if i % 1000 == 0 {
				log.Println(i, "/", maxI - minI + 1)
			}
		}
		close(tasks)
	}()
	for i := 0; i < nWorkers; i++ {
		go downloadWorker(done, tasks)
	}
	for i := 0; i < nWorkers; i++ {
		stats := <- done
		total.success += stats.success
		total.exists += stats.exists
	}
	log.Println("success: ", total.success, " exists: ", total.exists, " other: ", maxI - minI + 1 - total.success - total.exists, " | total: ", maxI - minI + 1)
}
