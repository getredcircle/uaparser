package uaparser

var (
	iPad = &ItemSpec{
		Name:         "iPad",
		mustContains: []string{"iPad"},
	}

	iPhone = &ItemSpec{
		Name:         "iPhone",
		mustContains: []string{"iPhone"},
	}

	iPod = &ItemSpec{
		Name:         "iPod",
		mustContains: []string{"iPod"},
	}

	mac = &ItemSpec{
		Name:         "Macintosh",
		mustContains: []string{"Macintosh"},
	}

	DEVICES = []*ItemSpec{
		iPad,
		iPhone,
		iPod,
		mac,
	}
)
