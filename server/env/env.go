package env

import (
	"log"
	"os"
	"strings"
)

func Init() {
  bytes, err := os.ReadFile(".env")

  if err != nil {
    log.Println("ERROR: Reading .env file failed.")
    log.Fatalln(err)
  }

  file := string(bytes) 

  for _, line := range strings.Split(file, "\n") {

    keyValue := strings.SplitN(line, "=", 2)

    os.Setenv(keyValue[0], keyValue[1])
  }
}
