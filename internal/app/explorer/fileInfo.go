package explorer

import (
	"fmt"
	"io/fs"
	"time"

	"gofm/pkg/utils"
)
// TODO uid, gid
type FileInfo struct {
	Mode       fs.FileMode
	ModeTime   time.Time
	Size       float64
	MemoryUnit utils.MemoryUnit
}

func (f FileInfo) String() string {
	mode := f.Mode.String()
	modeTime := f.ModeTime.Format("Mon Jan 2 15:04:05 2006")

	return fmt.Sprintf("%s\t%.1f %s\t%s", mode, f.Size, f.MemoryUnit, modeTime)
}
