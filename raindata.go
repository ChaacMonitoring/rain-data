package main

import (
  "fmt"
  "flag"
  "os"
  "log"

  "github.com/abhishekkr/goshare"
)

var (
  logfile     = flag.String("log-file", "rain-data.log", "to dump run-logs to")
)

func log_it(msg string){
    f, err := os.OpenFile(*logfile, os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
    if err != nil {
        fmt.Println("error opening file: %v", err)
    }
    defer f.Close()

    log.SetOutput(f)
    log.Println(msg)
}

func banner(){
  fmt.Println("**************************************************")
  fmt.Println("   _____       _             _____        _        ")
  fmt.Println("  |  __ \\     (_)           |  __ \\      | |       ")
  fmt.Println("  | |__) |__ _ _ _ __ ______| |  | | __ _| |_ __ _ ")
  fmt.Println("  |  _  // _` | | '_ \\______| |  | |/ _` | __/ _` |")
  fmt.Println("  | | \\ \\ (_| | | | | |     | |__| | (_| | || (_| |")
  fmt.Println("  |_|  \\_\\__,_|_|_| |_|     |_____/ \\__,_|\\__\\__,_|")
  fmt.Println("                                                  ")
  fmt.Println("**************************************************")
}

func main(){
  banner()
  goshare.GoShare()
}
