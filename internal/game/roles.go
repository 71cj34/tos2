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
)

func readRoleFile(rolefile ...string) ([]string, error) {
	// todo: rolelist switching
	_, filename, _, _ := runtime.Caller(0)
	file, err := os.Open(filepath.Join(filepath.Dir(filename), "..\\..", "roles", "default.txt"))
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

func HandleRoles(roles []string) {
	var cr string
	for i := 0; i < len(roles); i++ {
		cr = strings.TrimSpace(strings.ToLower(roles[i]))
		if strings.Contains(cr, "|") {
			switch cr[:1] {
				case "t", "c", "n", "na":
					Players[i].alignment = cr[:1]
				default:
					fmt.Println("Unknown alignment found: ", cr[:1])
			}




		}
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

	HandleRoles(shuffledroles)
}