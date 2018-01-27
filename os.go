package uaparser

var (
	linux = &ItemSpec{
		Name:            "Linux",
		mustContains:    []string{"Linux"},
		mustNotContains: []string{"Android"},
	}

	macOS = &ItemSpec{
		Name:            "Mac OS",
		mustContains:    []string{"Mac OS", "Macintosh"},
		mustNotContains: []string{"iPad", "iPhone", "iPod"},
		versionSplitters: [][]string{
			[]string{"Mac OS ", ";"},
			[]string{"Mac OS ", ")"},
		},
	}

	windows = &ItemSpec{
		Name:            "Windows",
		mustContains:    []string{"Windows"},
		mustNotContains: []string{"Windows Phone"},
		versionSplitters: [][]string{
			[]string{"Windows ", ";"},
		},
	}

	android = &ItemSpec{
		Name:            "Android",
		mustContains:    []string{"Android"},
		mustNotContains: []string{"Windows Phone"},
		versionSplitters: [][]string{
			[]string{"Android ", ";"},
			[]string{"Android-", " "},
		},
	}

	iOS = &ItemSpec{
		Name:            "iOS",
		mustContains:    []string{"CPU", "OS", "like Mac OS X", "iphone os"},
		mustNotContains: []string{"Windows Phone"},
		versionSplitters: [][]string{
			[]string{"CPU iPhone OS ", " "},
			[]string{"CPU OS ", " "},
		},
	}

	wpOS = &ItemSpec{
		Name:            "Windows Phone OS",
		mustContains:    []string{"Windows Phone OS", "Windows Phone"},
		mustNotContains: []string{},
		versionSplitters: [][]string{
			[]string{"Windows Phone OS ", ";"},
		},
	}

	playstationOS = &ItemSpec{
		Name:         "PlayStation OS",
		mustContains: []string{"PlayStation", "PLAYSTATION"},
		versionSplitters: [][]string{
			[]string{"PLAYSTATION", "3"},
			[]string{"PlayStation", "4"},
		},
	}

	OS = []*ItemSpec{
		linux,
		macOS,
		windows,
		android,
		iOS,
		wpOS,
		playstationOS,
	}
)
