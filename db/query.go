package db

const (
	GetLatestVersionOfAppQuery = `
		select app, major, minor, patch
		from res 
		where app = ?
		ORDER BY major DESC, minor DESC, patch DESC
		LIMIT 1;
	`
	UpdateLatestVersionOfAppQuery = `INSERT INTO res(app, major, minor, patch) values (?, ?, ?, ?);`
)
