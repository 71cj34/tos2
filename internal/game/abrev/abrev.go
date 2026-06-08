package abrev

import "strings"

type RoleCategory struct {
	Name       string
	Alignments map[string]AlignmentCategory
}

type AlignmentCategory struct {
	Name  string
	Roles map[string]string
}

var roleData = []RoleCategory{
	{
		Name: "Town",
		Alignments: map[string]AlignmentCategory{
			"tp": {Name: "Town Protective", Roles: map[string]string{
				"bg":      "bodyguard",
				"crus":    "crusader",
				"cleric":  "cleric",
				"trapper": "trapper",
				"oracle":  "oracle",
			}},
			"ti": {Name: "Town Investigative", Roles: map[string]string{
				"coro":    "coroner",
				"invest":  "investigator",
				"lo":      "lookout",
				"psy":     "psychic",
				"seer":    "seer",
				"spy":     "spy",
				"track":   "tracker",
				"sheriff": "sheriff",
			}},
			"tpow": {Name: "Town Power", Roles: map[string]string{
				"jailor": "jailor",
				"mayor":  "mayor",
				"mon":    "monarch",
				"pros":   "prosecutor",
				"marsh":  "marshal",
			}},
			"ts": {Name: "Town Support", Roles: map[string]string{
				"admi": "admirer",
				"amne": "amnesiac",
				"ret":  "retributionist",
				"soc":  "socialite",
				"tav":  "tavern keeper",
			}},
			"tk": {Name: "Town Killing", Roles: map[string]string{
				"dep":   "deputy",
				"trick": "trickster",
				"vet":   "veteran",
				"vig":   "vigilante",
			}},
		},
	},
	{
		Name: "Coven",
		Alignments: map[string]AlignmentCategory{
			"ck": {Name: "Coven Killing", Roles: map[string]string{
				"conj": "conjurer",
				"rit":  "ritualist",
				"jinx": "jinx",
			}},
			"cpow": {Name: "Coven Power", Roles: map[string]string{
				"cl":    "archmage",
				"hex":   "hex master",
				"witch": "witch",
			}},
			"cd": {Name: "Coven Deception", Roles: map[string]string{
				"dw":   "dreamweaver",
				"ench": "enchanter",
				"dusa": "medusa",
				"illu": "illusionist",
			}},
			"cu": {Name: "Coven Utility", Roles: map[string]string{
				"necro": "necromancer",
				"pois":  "poisoner",
				"pm":    "potion master",
				"vm":    "voodoo master",
				"wild":  "wildling",
			}},
		},
	},
	{
		Name: "Neutral",
		Alignments: map[string]AlignmentCategory{
			"nk": {Name: "Neutral Killing", Roles: map[string]string{"arso": "arsonist", "sk": "serial killer", "shroud": "shroud", "ww": "werewolf"}},
			"ne": {Name: "Neutral Evil", Roles: map[string]string{"exe": "executioner", "jest": "jester", "pirate": "pirate", "doom": "doomsayer"}},
		},
	},
	{
		Name: "Neutral Apocalypse",
		Alignments: map[string]AlignmentCategory{
			"na": {Name: "Neutral Apocalypse", Roles: map[string]string{"sc": "soul collector", "pb": "plaguebearer", "bers": "berserker", "baker": "baker", "death": "death", "famine": "famine", "pest": "pestilence", "war": "war"}},
		},
	},
}

// get a role's alignment
func GetCategory(key string) string {
	for _, cat := range roleCategories {
		if _, exists := cat.Roles[key]; exists {
			return cat.Name
		}
	}
	return "Unknown"
}

// value -> key
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

// key -> value
func UnpackRole(key string) (string, bool) {
	for _, cat := range roleCategories {
		if val, exists := cat.Roles[key]; exists {
			return val, true
		}
	}
	return "", false
}

func GetRandomRoleFromAlignment(key string) (string, bool) {

}

func GetRandomRoleFromAlignments(keys []string) (string, bool) {

}