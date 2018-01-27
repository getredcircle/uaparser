package uaparser

var (
	desktop = &ItemSpec{
		Name:            "Desktop",
		mustContains:    []string{"Windows", "Linux", "Mac OS X", "CrOS", "Macintosh"},
		mustNotContains: []string{"Windows Phone", "Android", "ARM"},
	}

	tablet = &ItemSpec{
		Name:         "Tablet",
		mustContains: []string{"iPad", "Android", "ARM", "PlayBook"},
		mustNotContains: []string{
			"Mobile ",
			"C6802", // Xperia Z Android (which is a phone) does not include Mobile in UA-string so without this is seen as a tablet
			"Windows Phone",
		},
	}

	phone = &ItemSpec{
		Name: "Phone",
		mustContains: []string{
			"iPhone",
			"Android",
			"Windows Phone",
			"BB10",
			"BlackBerry",
		},
		mustNotContains: []string{},
	}

	car = &ItemSpec{
		Name: "Car",
		mustContains: []string{
			"QtCarBrowser",
		},
		mustNotContains: []string{},
	}

	smartTv = &ItemSpec{
		Name: "SmartTV",
		mustContains: []string{
			"SMART-TV",
			"AppleTV",
			"CrKey",
			"Large Screen",
			"HbbTV",
			"LG Browser",
			"PhilipsTV",
			"Opera TV",
			"PlayStation",
		},
		mustNotContains: []string{},
	}

	DEVICETYPES = []*ItemSpec{
		tablet,
		phone,
		car,
		smartTv,
		desktop,
	}
)
