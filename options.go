package lotusdb

const (
	DefaultDBPath = "/tmp/lotusdb"
)

const (
	DefaultVLogBlockSize = 16 * 1024 * 1024 // 16MB
)

type Options struct {
	DBPath string
}

type ColumnFamilyOptions struct {
	// DirPath
	DirPath string

	// MemtableSize
	MemtableSize uint64

	MemtableNum int

	WalDir string

	WalMMap bool

	DisableWal bool

	ValueLogBlockSize int64
}
