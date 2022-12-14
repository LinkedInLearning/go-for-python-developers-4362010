package shutdown

import (
	"fmt"
	"os"

	"github.com/353solutions/shutdown/log"
)

var (
	logger = log.New("shutdown")
)

// ShutdownServer shuts down server from pid in pidFile.
func ShutdownServer(pidFile string) error {
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

	logger.Infow("shutdown", "pid", pid)
	// TODO: Actual shutdown

	if err := os.Remove(pidFile); err != nil {
		logger.Warnw("can't delete", "file", pidFile, "error", err)
	}

	return nil
}
