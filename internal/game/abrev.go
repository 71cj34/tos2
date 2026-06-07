package game

var rolesMap = map[string]string{
	"admi":      "admirer",
	"amne":       "amnesiac",
	"bg":       "bodyguard",
	"coro":       "coroner",
	"crus":       "crusader",
	"cleric": "cleric",
	"dep":        "deputy",
	"invest":     "investigator",
	"jailor":       "jailor",
	"lo":        "lookout",
	"marsh":      "marshal",
	"mayor":      "mayor",
	"mon":    "monarch",
	"oracle": "oracle",
	"pros":       "prosecutor",
	"psy":      "psychic",
	"ret":        "retributionist",
	"seer":       "seer",
	"soc":        "socialite",
	"spy":        "spy",
	"tav":        "tavern keeper",
	"track":         "tracker",
	"trapper":        "trapper",
	"trick":        "trickster",
	"sheriff":         "sheriff",
	"vet":        "veteran",
	"vig":        "vigilante",

	"cl":       "archmage",
	"conj":       "conjurer",
	"dusa":       "medusa",
	"dw":      "dreamweaver",
	"ench":     "enchanter",
	"hex":        "hex master",
	"illu":       "illusionist",
	"jinx":       "jinx",
	"necro":        "necromancer",
	"pois":      "poisoner",
	"pm":       "potion master",
	"rit":       "ritualist",
	"vm":        "voodoo master",
	"witch":       "witch",
	"wild":       "wildling",

	"sc": "soul collector",
	"pb": "plaguebearer",
	"bers": "berserker",
	"baker": "baker",
	"death":        "death",
	"famine":     "famine",
	"pest":       "pestilence",
	"war":        "war",

	"arso":       "arsonist",
	"doom":       "doomsayer",
	"exe":        "executioner",
	"jest":       "jester",
	"pirate":        "pirate",
	"sk":         "serial killer",
	"shroud":       "shroud",
	"ww":       "werewolf",
}

func UnpackRole(key string) (string, bool) {
	val, ok := rolesMap[key]
    return val, ok
}

func RepackRole (target string) (string, bool) {
    for key, value := range rolesMap {
        if value == target {
            return key, true
        }
    }
    return "", false
}

