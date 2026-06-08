package game

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"math/rand"
	"errors"
	"strings"
	"71cj34/tos2/internal/game/abrev"
)

func readRoleFile(rolefile ...string) ([]string, error) {
	// todo: rolelist switching
	_, filename, _, _ := runtime.Caller(0)
	file, err := os.Open(filepath.Join(filepath.Dir(filename), "..\\..", "roles", "hideandseek.txt"))
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var roles []string

	fmt.Println("== ROLELIST ==")
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
		roles = append(roles, line)
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %v\n", err)
	}

	if len(roles) != 15 {return nil, errors.New("Rolelist does not have 15 items")}

	return roles, nil
}

var temp []string

func AssignRoles(roles []string) {
	for i := 0; i < len(roles); i++ {
		var cr string
		var role string
		var ok bool
		cr = strings.TrimSpace(strings.ToLower(roles[i]))
		if strings.Contains(cr, "|") {
			switch strings.Split(cr, "|")[0] {
				// we can't just set role and align here since could be complex, eg na|tk/tpow
				case "t":
					Players[i].team = "Town"
				case "c":
					Players[i].team = "Coven"
				case "na":
					Players[i].team = "Neutral Apocalypse"
				default:
					fmt.Println("Error parsing rolelist: Unknown alignment found: ", cr[:1])
			}

			cr = strings.Split(cr, "|")[1]
		}

		crsplit := strings.Split(cr, "/")
		crrandom := crsplit[rand.Intn(len(crsplit))]

		// this whole if block is about setting the role
		// we can get alignment/team afterwards

		// crrandom will be a role literal if we chopped off the _| part of it already above
		_, _, ok = abrev.GetUpperInfo(crrandom)
		if ok {
			role = crrandom
		} else if crrandom[:1] == "~" {
			switch crrandom[1:3] {
				case "rt":
					role, ok = abrev.GetRandomRoleFromAlignments([]string{"tpow", "ti", "tp", "ts", "tk"})
					if !ok {fmt.Println("Error in GetRandomRoleFromAlignments with case ",crrandom[1:3])}
				case "ct":
					role, ok = abrev.GetRandomRoleFromAlignments([]string{"ti", "tp", "ts", "tk"})
					if !ok {fmt.Println("Error in GetRandomRoleFromAlignments with case ",crrandom[1:3])}
				case "rc":
					role, ok = abrev.GetRandomRoleFromAlignments([]string{"cpow", "ck", "cd", "cu"})
					if !ok {fmt.Println("Error in GetRandomRoleFromAlignments with case ",crrandom[1:3])}
				case "cc":
					role, ok = abrev.GetRandomRoleFromAlignments([]string{"ck", "cd", "cu"})
					if !ok {fmt.Println("Error in GetRandomRoleFromAlignments with case ",crrandom[1:3])}
				default:
					fmt.Println("Error parsing rolelist: Incorrect characters after ~, got ",crrandom[1:3])
			}
		} else {
			role, ok = abrev.GetRandomRoleFromAlignment(crrandom)
			if !ok {fmt.Println("Error in GetRandomRoleFromAlignment with case ",crrandom)}
		}
		Players[i].role = role

		team, align, ok := abrev.GetUpperInfo(role)
		temp = append(temp, role)
		if !ok {fmt.Println("Error getting upper info for role ", role)}
		if Players[i].team == "" {
			Players[i].team = team
		}
		Players[i].alignment = align

	}
	fmt.Println(roles)
	fmt.Println("Success!! Here are all the roles: ")
	for i := 0; i < len(roles); i++ {
		tempv, _ := abrev.UnpackRole(Players[i].role)
		fmt.Printf("%-2d: %-20s %-20s %-20s (%-8s)\n", i+1, tempv, Players[i].alignment, Players[i].team, temp[i])
	}
}

func GenRoles() {
	roles, _ := readRoleFile()

	shuffledroles := make([]string, len(roles))
	copy(shuffledroles, roles)
	for i := len(shuffledroles) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		shuffledroles[i], shuffledroles[j] = shuffledroles[j], shuffledroles[i]
	}

	AssignRoles(shuffledroles)
}