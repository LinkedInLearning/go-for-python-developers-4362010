package killer

import (
	"fmt"
	"os"

	"github.com/353solutions/killer/log"
)

var (
	logger = log.New("killer")
)

// KillServer kill server from pid in pidFile.
func KillServer(pidFile string) error {
	file, err := os.Open(pidFile)
	if err != nil {
		logger.Errorw("can't open", "file", pidFile, "error", err)
		return err
	}
	defer file.Close()

	var pid int
	if _, err := fmt.Fscanf(file, "%d", &pid); err != nil {
		logger.Errorw("bad pid", "file", pidFile, "error", err)
		return err

	}

	logger.Infow("killing", "pid", pid)
	// TODO: Actual kill

	if err := os.Remove(pidFile); err != nil {
		logger.Warnw("can't delete", "file", pidFile, "error", err)
	}

	return nil
}
