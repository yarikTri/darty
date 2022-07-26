package darty

import "time"

// Event includes information about any event
type Event struct {
	ID          int
	Title       string
	Organizer   User
	Description string
	Date        time.Time
	Site        string // Link
	// Location ???    // coordinates (later)
	// Photo    ???
}
