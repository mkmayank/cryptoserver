package config

import (
	"fmt"
	"os"
	"path"
)

// paths config
type paths struct {
	WorkDir string `json:"work_dir"`
}

func (p paths) checkSanity() {

	if p.WorkDir == "" {
		fmt.Println("work_dir is not defined in given config file")
		os.Exit(2)
	}

	err := os.MkdirAll(p.WorkDir, 0755)
	if err != nil {
		fmt.Printf("error while making work dir: %s, error: %s", rawConfig.Paths.WorkDir, err.Error())
		os.Exit(2)
	}
}

// PathLogFile returns path of log file
func PathLogFile(tag string) string {
	return path.Join(rawConfig.Paths.WorkDir, fmt.Sprintf("logfile.%s", tag))
}
