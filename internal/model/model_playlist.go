package model

import "time"

type Song struct {
	ID       int64
	Title    string
	Article  string
	TimeSong time.Duration
	Prev     int64
	Next     int64
	Playnig  bool
	Duration time.Duration
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

type SongRequest struct {
	Login    string
	Title    string
	Article  string
	Duration time.Duration
}
