package storage

import (
	"database/sql"

	_ "modernc.org/sqlite"
)

func newSqliteStorager() (Storager, error) {
	s := sqliteStorager{path: "database.db"}
	d, err := sql.Open("sqlite", s.path)
	if err != nil {
		s.db = d
		return s, s.initStorage()
	}
	return nil, err
}

type sqliteStorager struct {
	path string
	db   *sql.DB
}

func (s sqliteStorager) initStorage() error {
	_, err := s.db.Exec(
		`CREATE TABLE IF NOT EXISTS passwords (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			username TEXT NOT NULL,
			resource TEXT NOT NULL,
			password BLOB NOT NULL,
			created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
			UNIQUE(username, resource));
			
		CREATE TABLE IF NOT EXISTS master (
            id INTEGER PRIMARY KEY CHECK (id = 1),
            master_key BLOB NOT NULL,      
            salt BLOB NOT NULL, 
            created_at DATETIME DEFAULT CURRENT_TIMESTAMP);`)
	return err
}

func (s sqliteStorager) SaveMasterKey(entry MasterEntry) error {
	_, err := s.db.Exec(
		`INSERT OR REPLACE INTO master (id, master_key, salt, created_at) 
        VALUES (1, ?, ?, CURRENT_TIMESTAMP)`,
		entry.Hash, entry.Salt)
	return err
}

func (s sqliteStorager) ReadMasterKey() (MasterEntry, error) {
	var entry MasterEntry
	err := s.db.QueryRow(
		`SELECT master_key, salt FROM master WHERE id = 1
		`).Scan(&entry.Hash, &entry.Salt)
	if err == sql.ErrNoRows {
		err = nil
	}
	return entry, err
}

func (s sqliteStorager) GetPassword(filter EntryFilter) ([]PasswordEntry, error) {
	rows, err := s.db.Query(
		`SELECT username, resource, password FROM passwords`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var results []PasswordEntry
	for rows.Next() {
		var p PasswordEntry
		err := rows.Scan(&p.Username, &p.Resource, &p.Password)
		if err != nil {
			return nil, err
		}
		if filter(p.Username, p.Resource) {
			results = append(results, p)
		}
	}
	return results, nil
}

func (s sqliteStorager) SavePassword(entry PasswordEntry) error {
	_, err := s.db.Exec(
		`INSERT OR REPLACE INTO passwords 
		(username, resource, password, created_at) 
        VALUES (?, ?, ?, CURRENT_TIMESTAMP)`,
		entry.Username, entry.Resource, entry.Password)
	return err
}

func (s sqliteStorager) DeletePassword(filter EntryFilter) error {
	entries, err := s.GetPassword(filter)
	if err != nil {
		return err
	}
	for _, entry := range entries {
		_, err = s.db.Exec(
			`DELETE FROM passwords 
            WHERE username = ? AND resource = ?`,
			entry.Username, entry.Resource)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s sqliteStorager) Close() error {
	return s.db.Close()
}
