package main

import (
	"fmt"
	"os"
)

// import (
// 	"fmt"
// )

func main() {
	if !IsIntialized() {
		_, err := InitConfig()
		if err != nil {
			fmt.Println(WriteConfError)
			return
		}
	}

	config, err := readconfig()

}

func readconfig() (*config, *error) {
	configFile, err := os.ReadFile(ConfigDirDef)
}
