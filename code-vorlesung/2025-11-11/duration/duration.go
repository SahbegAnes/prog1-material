package duration

type Duration int

func FromSeconds(s int) Duration {
	return Duration(s)
}

func FromMinutes(m int) Duration {
	return Duration(m * 60)
}

func (d Duration) Seconds() int {
	return int(d)
}
