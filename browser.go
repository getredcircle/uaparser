package uaparser

var (
	ie = &ItemSpec{
		Name:         "IE",
		mustContains: []string{"MSIE", "rv:11.0", "Edge/12.0", "IEMobile"},
		mustNotContains: []string{
			"360SE",
			"Maxthon",
			"qihu",
			"QIHU",
			"QQBrowser",
			"QQDownload",
			"SE ",
			"MetaSr",
			"TencentTraveler",
			"Firefox",
			"Opera",
			"MAXTHON",
		},
		versionSplitters: [][]string{[]string{"MSIE ", ";"}},
	}

	firefox = &ItemSpec{
		Name:             "Firefox",
		mustContains:     []string{"Firefox", "firefox"},
		mustNotContains:  []string{"Seamonkey", "Opera"},
		versionSplitters: [][]string{[]string{"Firefox/", " "}},
	}

	safari = &ItemSpec{
		Name:         "Safari",
		mustContains: []string{"Safari", "AppleWebKit"},
		mustNotContains: []string{
			"Chrome",
			"Chromium",
			"CoolNovo",
			"Maxthon",
			"LBBROWSER",
			"QIHU",
			"QQBrowser",
			"SE ",
			"MetaSr",
			"TaoBrowser",
			"PlayStation",
			"PLAYSTATION",
			"Opera",
		},
		versionSplitters: [][]string{
			[]string{"Version/", " "},
		},
	}

	chrome = &ItemSpec{
		Name:         "Chrome",
		mustContains: []string{"Chrome"},
		mustNotContains: []string{
			"CoolNovo",
			"LBBROWSER",
			"Maxthon",
			"QIHU",
			"QQBrowser",
			"SE ",
			"Edge",
			"MetaSr",
			"TaoBrowser",
		},
		versionSplitters: [][]string{[]string{"Chrome/", " "}},
	}

	edge = &ItemSpec{
		Name:         "Edge",
		mustContains: []string{"Chrome", "Edge"},
		mustNotContains: []string{
			"CoolNovo",
			"LBBROWSER",
			"Maxthon",
			"QIHU",
			"QQBrowser",
			"SE ",
			"MetaSr",
			"TaoBrowser",
		},
		versionSplitters: [][]string{[]string{"Edge/", " "}},
	}

	opera = &ItemSpec{
		Name:            "Opera",
		mustContains:    []string{"Opera"},
		mustNotContains: []string{},
		versionSplitters: [][]string{
			[]string{"Version/", " "},
			[]string{"Opera/", " "},
		},
	}

	_360se = &ItemSpec{
		Name:            "360SE",
		mustContains:    []string{"360SE", "qihu", "QIHU"},
		mustNotContains: []string{},
	}

	sougou = &ItemSpec{
		Name:             "Sougou",
		mustContains:     []string{"SE ", "MetaSr"},
		mustNotContains:  []string{},
		versionSplitters: [][]string{[]string{"SE ", " "}},
	}

	tencent = &ItemSpec{
		Name:             "Tencent Traveler",
		mustContains:     []string{"TencentTraveler"},
		mustNotContains:  []string{"SE ", "MetaSr"},
		versionSplitters: [][]string{[]string{"TencentTraveler ", " "}},
	}

	qq = &ItemSpec{
		Name:             "QQ Browser",
		mustContains:     []string{"QQBrowser"},
		mustNotContains:  []string{},
		versionSplitters: [][]string{[]string{"QQBrowser/", " "}},
	}

	maxthon = &ItemSpec{
		Name:            "Maxthon",
		mustContains:    []string{"Maxthon", "MAXTHON"},
		mustNotContains: []string{},
		versionSplitters: [][]string{
			[]string{"Maxthon/", " "},
			[]string{"Maxthon ", " "},
		},
	}

	playstation = &ItemSpec{
		Name:         "PlayStation",
		mustContains: []string{"PlayStation", "PLAYSTATION"},
		versionSplitters: [][]string{
			[]string{"PLAYSTATION 3", ")"},
			[]string{"PlayStation 4", ""},
		},
	}

	BROWSERS = []*ItemSpec{
		ie,
		firefox,
		safari,
		chrome,
		edge,
		opera,
		_360se,
		sougou,
		tencent,
		qq,
		maxthon,
		playstation,
	}
)
