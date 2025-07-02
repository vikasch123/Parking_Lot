package vehicle

type Vehicle struct {
	Plate      string
	Color      string
	Make       string
	IsHandicap bool
	Size       string
}

func New(pl string, col string, make string, handicap bool, size string) *Vehicle {
	return &Vehicle{
		Plate:      pl,
		Color:      col,
		Make:       make,
		IsHandicap: handicap,
		Size:       size,
	}
}
