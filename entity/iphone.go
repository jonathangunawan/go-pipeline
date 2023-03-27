package entity

type Phone struct {
	Numbering          int
	SerialNumber       string
	IsDisplayInstalled bool
	IsBatteryInstalled bool
	IsFunctional       bool
	IsAlreadyPacked    bool
}

type Box []Phone
