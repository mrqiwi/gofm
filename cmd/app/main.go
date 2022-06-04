package main

import (
	"log"
	"log/syslog"

	"gofm/internal/app/explorer"
	"gofm/internal/app/tui"
)

const appName = "gofm"

// TODO
// 1) SFTP
// 2) FTP
// 3) SMB
// 4) YANDEX DISK

func main() {
	logger, err := syslog.New(syslog.LOG_DEBUG, appName)
	if err != nil {
		log.Fatal(err)
	}

	newExplorer := explorer.NewFileExplorer(logger)
	app := tui.NewTUI(newExplorer, logger)

	err = app.Run()
	if err != nil {
		log.Fatal(err)
	}
}
