package tokenizer

type Parent int
type Children []int

type Rules struct {
	parent      Parent
	allChildren []Children
}

var colRules = Rules{COLLECTION,
	[]Children{
		{DATABASE, DOT, DBMETHOD},
	},
}

var expRules = Rules{EXP,
	[]Children{
		{SECONDBRACKETLEFT, KEY, COLON, VALUE, SECONDBRACKETRIGHT},
	},
}

var resultsRules = Rules{RESULTS,
	[]Children{
		{COLLECTION, DOT, COLMETHOD},
	},
}

var dbmethodRules = Rules{DBMETHOD,
	[]Children{
		{DBMETHODNAME, FIRSTBRACKETLEFT, ID, FIRSTBRACKETRIGHT},
	},
}

var colmethodRules = Rules{COLMETHOD,
	[]Children{
		{COLMETHODNAME, FIRSTBRACKETLEFT, EXP, FIRSTBRACKETRIGHT},
	},
}

var allRules = []Rules{}

func init() {
	allRules = append(allRules,
		colRules,
		expRules,
		resultsRules,
		dbmethodRules,
		colmethodRules,
	)
}
