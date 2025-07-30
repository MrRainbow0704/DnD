package types

// Coontext key for storing values in context
type CtxKey string

// Alias for map[string]any, used to represent any JSON-like structure
type AnyMap map[string]any

// Alias for map[string]error, used to represent a collection of errors
type ErrorMap map[string]error

type ClassLevel struct {
	Name   string `json:"name"`
	Level  int64  `json:"level"`
	HitDie int64  `json:"hitdie"`
}

type Misc struct {
	ClassLevel []ClassLevel `json:"classlevel"`
	Background string       `json:"background"`
	PlayerName string       `json:"playername"`
	Race       string       `json:"race"`
	Alignment  string       `json:"alignment"`
	Experience int64        `json:"experience"`
}

type Scores struct {
	Strength     int64 `json:"strength"`
	Dexterity    int64 `json:"dexterity"`
	Constitution int64 `json:"constitution"`
	Intelligence int64 `json:"intelligence"`
	Wisdom       int64 `json:"wisdom"`
	Charisma     int64 `json:"charisma"`
}

type Saves struct {
	Strength     bool `json:"strength"`
	Dexterity    bool `json:"dexterity"`
	Constitution bool `json:"constitution"`
	Intelligence bool `json:"intelligence"`
	Wisdom       bool `json:"wisdom"`
	Charisma     bool `json:"charisma"`
}

type Skill struct {
	Proficiency bool   `json:"proficency"`
	Expertise   bool   `json:"expertise"`
	Score       string `json:"score"`
}

type Skills struct {
	Athletics      Skill `json:"athletics"`
	Acrobatics     Skill `json:"acrobatics"`
	SleightOfHand  Skill `json:"sleightofhand"`
	Stealth        Skill `json:"stealth"`
	Arcana         Skill `json:"arcana"`
	Historical     Skill `json:"history"`
	Investigation  Skill `json:"investigation"`
	Nature         Skill `json:"nature"`
	Religion       Skill `json:"religion"`
	AnimalHandling Skill `json:"animalhandling"`
	Insight        Skill `json:"insight"`
	Medicine       Skill `json:"medicine"`
	Perception     Skill `json:"perception"`
	Survival       Skill `json:"survival"`
	Deception      Skill `json:"deception"`
	Intimidation   Skill `json:"intimidation"`
	Performance    Skill `json:"performance"`
	Persuasion     Skill `json:"persuasion"`
}

type Speed struct {
	Type  string `json:"type"`
	Value int64  `json:"value"`
}

type Attack struct {
	Name   string `json:"name"`
	Bonus  int64  `json:"bonus"`
	Damage string `json:"damage"`
}

type Combat struct {
	AC           int64    `json:"ac"`
	MaxHP        int64    `json:"maxhp"`
	CurrentHP    int64    `json:"currenthp"`
	TempHP       int64    `json:"temphp"`
	DeathSuccess int64    `json:"deathsuccess"`
	DeathFail    int64    `json:"deathfail"`
	Speeds       []Speed  `json:"speeds"`
	Attacks      []Attack `json:"attacks"`
}

type Money struct {
	Copper   int64 `json:"copper"`
	Silver   int64 `json:"silver"`
	Electrum int64 `json:"electrum"`
	Gold     int64 `json:"gold"`
	Platinum int64 `json:"platinum"`
}

type Equipment struct {
	Other string `json:"other"`
	Money Money  `json:"money"`
}

type Trait struct {
	Name string `json:"name"`
	Desc string `json:"desc"`
}

type Flavor struct {
	Personality string `json:"personality"`
	Ideals      string `json:"ideals"`
	Bonds       string `json:"bonds"`
	Flaws       string `json:"flaws"`
}

// Represents the full structure of a character
type Character struct {
	Owner      int64     `json:"owner"`
	CharName   string    `json:"charname"`
	OtherProfs string    `json:"otherprofs"`
	Misc       Misc      `json:"misc"`
	Scores     Scores    `json:"scores"`
	Saves      Saves     `json:"saves"`
	Skills     Skills    `json:"skills"`
	Combat     Combat    `json:"combat"`
	Equipment  Equipment `json:"equipment"`
	Flavor     Flavor    `json:"flavor"`
	Traits     []Trait   `json:"traits"`
}
