package model

import "time"

type Song struct {
	ID       int64
	Title    string
	Artist   string
	TimeSong time.Duration
	Prev     *Song
	Next     *Song
}

type Playlist struct {
	User    string
	Current *Song
	Playing bool
}

type Command struct {
	User   string
	Action string // action: "pause", "play", "next", "prev"
	Song   *Song  // maybe empty for play/pause
}
