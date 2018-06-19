package main

import (
    "fmt"
    "log"
    "time"
    "os"
    "github.com/sparrc/go-ping"
)

func main() {

    i := 0

    target := os.Args[1]
    
    fmt.Printf("%s ~> ", target)

    for ;; {

        pinger, err := ping.NewPinger(target)
        if err != nil {
            log.Fatal(err)
        }
        
        pinger.Count = 1        
        pinger.Run()
        stats := pinger.Statistics()
        
        if stats.PacketLoss == 0 {
            fmt.Printf("\033[32mo\033[0m")
        } else {
            fmt.Printf("\033[31mx\033[0m")
        }
        
        time.Sleep(0.5 * time.Second)
        if i % 10 == 0 && i != 0 {
            fmt.Printf("\b\b\b\b\b\b\b\b\b\b")
        }
        
        i+=1
    }
}
