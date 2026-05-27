package storage

type Storager interface {
	SaveMasterKey(entry MasterEntry) error
	ReadMasterKey() (MasterEntry, error)

	SavePassword(entry PasswordEntry) error
	GetPassword(filter EntryFilter) ([]PasswordEntry, error)
	DeletePassword(filter EntryFilter) error
	Close() error
}

type PasswordEntry struct {
	Username string
	Resource string
	Password []byte
}

type MasterEntry struct {
	Hash []byte
	Salt []byte
}

type EntryFilter func(username string, resource string) bool

func New() (Storager, error) {
	return newSqliteStorager()
}
