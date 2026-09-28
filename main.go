package main

import (
	"fmt"
	"musicplaylist/music"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage:")
		fmt.Println("  go run . add")
		fmt.Println("  go run . list")
		fmt.Println("  go run . find T001")
		fmt.Println("  go run . play T001")
		fmt.Println("  go run . delete T001")
		return
	}

	switch os.Args[1] {
	case "add":
		track := music.Track{
			TrackID:     "T001",
			Title:       "Perfect",
			Artist:      "Ed Sheeran",
			DurationSec: 263,
			PlayCount:   0,
		}

		err := music.AddTrack(track)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		fmt.Println("Track added successfully")

	case "list":
		tracks, err := music.ListTracks()
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		if len(tracks) == 0 {
			fmt.Println("No tracks found.")
			return
		}

		fmt.Println("===== SAVED TRACKS =====")
		for _, track := range tracks {
			fmt.Printf("ID: %s | Title: %s | Artist: %s | Duration: %d sec | PlayCount: %d\n",
				track.TrackID, track.Title, track.Artist, track.DurationSec, track.PlayCount)
		}

	case "find":
		if len(os.Args) != 3 {
			fmt.Println("Usage: go run . find T001")
			return
		}

		track, err := music.FindTrack(os.Args[2])
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		fmt.Println("Track found:")
		fmt.Printf("ID: %s | Title: %s | Artist: %s | Duration: %d sec | PlayCount: %d\n",
			track.TrackID, track.Title, track.Artist, track.DurationSec, track.PlayCount)

	case "play":
		if len(os.Args) != 3 {
			fmt.Println("Usage: go run . play T001")
			return
		}

		err := music.IncrementPlayCount(os.Args[2])
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		track, _ := music.FindTrack(os.Args[2])
		fmt.Printf("Playing: %s - %s | PlayCount: %d\n", track.Title, track.Artist, track.PlayCount)

	case "delete":
		if len(os.Args) != 3 {
			fmt.Println("Usage: go run . delete T001")
			return
		}

		err := music.DeleteTrack(os.Args[2])
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		fmt.Println("Track deleted successfully")

	default:
		fmt.Println("Unknown command:", os.Args[1])
	}
}
