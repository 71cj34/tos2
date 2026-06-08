package abrev

import (
	"strings"
	"math/rand"
	"slices"
)

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

// get a role's alignment/team
func GetUpperInfo(key string) (string, string, bool) {
	for _, cat := range roleData {
		for _, align := range cat.Alignments {
			if _, exists := align.Roles[key]; exists {
				return cat.Name, align.Name, true
			}
		}
	}
	return "", "", false
}

// value -> key
func RepackRole(value string) (string, bool) {
	searchVal := strings.ToLower(value)
	for _, cat := range roleData {
		for _, align := range cat.Alignments {
			for k, v := range align.Roles {
				if strings.ToLower(v) == searchVal {
					return k, true
				}
			}
		}
	}
	return "", false
}

// key -> value
func UnpackRole(key string) (string, bool) {
	for _, cat := range roleData {
		for _, align := range cat.Alignments {
			if val, exists := align.Roles[key]; exists {
				return val, true
			}
		}
	}
	return "", false
}

func GetRandomRoleFromAlignment(key string) (string, bool) {
    var result string
    // the hackiest way ever to stop >4 na
    for i := 0; i < 150; i++ {
        for _, cat := range roleData {
            if align, exists := cat.Alignments[key]; exists {
                keys := make([]string, 0, len(align.Roles))
                for k := range align.Roles {
                    keys = append(keys, k)
                }

                result = keys[rand.Intn(len(keys))]

                if !slices.Contains(banned, result) {
                    // if it's not banned and not already used, return it
                    if !slices.Contains(uniqueRoles, result) {
                        return result, true
                    }
                    // if it's not banned but already used, add to banned list
                    banned = append(banned, result)
                    return result, true
                }

            }
        }
    }
    return "", false
}

func GetRandomRoleFromAlignments(keys []string) (string, bool) {
	if len(keys) == 0 {
		return "", false
	}

	target := keys[rand.Intn(len(keys))]

	return GetRandomRoleFromAlignment(target)
}