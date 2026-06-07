package abrev

import "strings"

type RoleCategory struct {
	Name  string
	Roles map[string]string
}

var roleCategories = []RoleCategory{
	{
		Name: "t",
		Roles: map[string]string{
			"admi": "admirer", "amne": "amnesiac", "bg": "bodyguard", "coro": "coroner",
			"crus": "crusader", "cleric": "cleric", "dep": "deputy", "invest": "investigator",
			"jailor": "jailor", "lo": "lookout", "marsh": "marshal", "mayor": "mayor",
			"mon": "monarch", "oracle": "oracle", "pros": "prosecutor", "psy": "psychic",
			"ret": "retributionist", "seer": "seer", "soc": "socialite", "spy": "spy",
			"tav": "tavern keeper", "track": "tracker", "trapper": "trapper",
			"trick": "trickster", "sheriff": "sheriff", "vet": "veteran", "vig": "vigilante",
		},
	},
	{
		Name: "c",
		Roles: map[string]string{
			"cl": "archmage", "conj": "conjurer", "dusa": "medusa", "dw": "dreamweaver",
			"ench": "enchanter", "hex": "hex master", "illu": "illusionist", "jinx": "jinx",
			"necro": "necromancer", "pois": "poisoner", "pm": "potion master", "rit": "ritualist",
			"vm": "voodoo master", "witch": "witch", "wild": "wildling",
		},
	},
	{
		Name: "na",
		Roles: map[string]string{
			"sc": "soul collector", "pb": "plaguebearer", "bers": "berserker",
			"baker": "baker", "death": "death", "famine": "famine",
			"pest": "pestilence", "war": "war",
		},
	},
	{
		Name: "n",
		Roles: map[string]string{
			"arso": "arsonist", "doom": "doomsayer", "exe": "executioner",
			"jest": "jester", "pirate": "pirate", "sk": "serial killer",
			"shroud": "shroud", "ww": "werewolf",
		},
	},
}

func GetCategory(key string) string {
	for _, cat := range roleCategories {
		if _, exists := cat.Roles[key]; exists {
			return cat.Name
		}
	}
	return "Unknown"
}

func RepackRole(value string) (string, bool) {
	searchVal := strings.ToLower(value)
	for _, cat := range roleCategories {
		for k, v := range cat.Roles {
			if strings.ToLower(v) == searchVal {
				return k, true
			}
		}
	}
	return "", false
}

func UnpackRole(key string) (string, bool) {
	for _, cat := range roleCategories {
		if val, exists := cat.Roles[key]; exists {
			return val, true
		}
	}
	return "", false
}