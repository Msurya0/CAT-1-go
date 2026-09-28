package music

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Track struct {
	TrackID     string
	Title       string
	Artist      string
	DurationSec int
	PlayCount   int
}

const dataFile = "tracks.txt"

var playlist = make(map[string]Track)

func loadTracks() error {
	playlist = make(map[string]Track)

	file, err := os.Open(dataFile)
	if os.IsNotExist(err) {
		return nil
	}
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		parts := strings.Split(line, "|")
		if len(parts) != 5 {
			continue
		}

		duration, err1 := strconv.Atoi(parts[3])
		playCount, err2 := strconv.Atoi(parts[4])
		if err1 != nil || err2 != nil {
			continue
		}

		playlist[parts[0]] = Track{
			TrackID:     parts[0],
			Title:       parts[1],
			Artist:      parts[2],
			DurationSec: duration,
			PlayCount:   playCount,
		}
	}

	return scanner.Err()
}

func saveTracks() error {
	file, err := os.Create(dataFile)
	if err != nil {
		return err
	}
	defer file.Close()

	for _, t := range playlist {
		_, err = fmt.Fprintf(file, "%s|%s|%s|%d|%d\n",
			t.TrackID, t.Title, t.Artist, t.DurationSec, t.PlayCount)
		if err != nil {
			return err
		}
	}

	return nil
}

func AddTrack(t Track) error {
	if err := loadTracks(); err != nil {
		return err
	}

	if t.TrackID == "" {
		return fmt.Errorf("track ID cannot be empty")
	}

	if t.DurationSec <= 0 {
		return fmt.Errorf("duration must be positive")
	}

	if _, exists := playlist[t.TrackID]; exists {
		return fmt.Errorf("track already exists")
	}

	playlist[t.TrackID] = t
	return saveTracks()
}

func FindTrack(id string) (*Track, error) {
	if err := loadTracks(); err != nil {
		return nil, err
	}

	t, exists := playlist[id]
	if !exists {
		return nil, fmt.Errorf("track not found")
	}

	return &t, nil
}

func IncrementPlayCount(id string) error {
	if err := loadTracks(); err != nil {
		return err
	}

	t, exists := playlist[id]
	if !exists {
		return fmt.Errorf("track not found")
	}

	t.PlayCount++
	playlist[id] = t

	return saveTracks()
}

func DeleteTrack(id string) error {
	if err := loadTracks(); err != nil {
		return err
	}

	if _, exists := playlist[id]; !exists {
		return fmt.Errorf("track not found")
	}

	delete(playlist, id)
	return saveTracks()
}

func ListTracks() ([]Track, error) {
	if err := loadTracks(); err != nil {
		return nil, err
	}

	tracks := make([]Track, 0, len(playlist))
	for _, t := range playlist {
		tracks = append(tracks, t)
	}

	return tracks, nil
}
