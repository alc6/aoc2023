package main

import (
	"fmt"
	"github.com/alc6/aoc2023/util"
	"log"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	lines, err := util.ReadFileLines("d2/input.txt")
	if err != nil {
		log.Fatalf("error reading file: %v", err)
	}

	games, err := parseGames(lines)
	if err != nil {
		log.Fatalf("error parsing games: %v", err)
	}

	pt1(games)
	pt2(games)
}

func pt1(games []Game) {
	limits := ColorsSet{
		"red":   12,
		"green": 13,
		"blue":  14,
	}

	compliantGames := possibleGames(games, limits)
	var sumIDs int
	for _, cg := range compliantGames {
		sumIDs += cg.ID
	}

	log.Println("pt1 sum of the ids of compliant games: ", sumIDs)
}

func pt2(games []Game) {
	var summedPow int
	for _, game := range games {
		summedPow += game.minColorSetRequired().Pow()
	}

	log.Println("pt2 sum of the pow of the min color set required: ", summedPow)
}

type ColorsSet map[string]int

func (cs ColorsSet) Pow() int {
	var pow = 1

	for _, value := range cs {
		pow *= value
	}

	return pow
}

type Game struct {
	ID       int
	Sessions []ColorsSet
}

func (g Game) minColorSetRequired() ColorsSet {
	colorSet := ColorsSet{
		"red":   0,
		"green": 0,
		"blue":  0,
	}

	for _, session := range g.Sessions {
		for color, value := range session {
			if colorSet[color] < value {
				colorSet[color] = value
			}
		}
	}

	return colorSet
}

func parseGames(lines []string) ([]Game, error) {
	var (
		games = make([]Game, 0, len(lines))
		re    = regexp.MustCompile(`Game (\d+): (.+)`)
	)

	for _, l := range lines {
		for _, match := range re.FindAllStringSubmatch(l, -1) {
			gameID, err := strconv.Atoi(match[1])
			if err != nil {
				return nil, err
			}

			sessions, err := parseSessions(match[2])
			if err != nil {
				return nil, fmt.Errorf("failed to parse sessions: %v", err)
			}

			games = append(games, Game{
				ID:       gameID,
				Sessions: sessions,
			})
		}
	}

	return games, nil
}

func parseSessions(data string) ([]ColorsSet, error) {
	var parsedSessions = make([]ColorsSet, 0)

	for _, session := range strings.Split(data, ";") {
		parsedSession, err := parseSession(session)
		if err != nil {
			return nil, err
		}

		parsedSessions = append(parsedSessions, parsedSession)
	}

	return parsedSessions, nil
}

func parseSession(data string) (ColorsSet, error) {
	var (
		session = make(ColorsSet)
		re      = regexp.MustCompile(`(\d+)([a-zA-Z]+)`)
	)

	for _, totalOnColor := range strings.Split(strings.Replace(data, " ", "", -1), ",") {
		matches := re.FindStringSubmatch(totalOnColor)
		if matches == nil {
			return nil, fmt.Errorf("could not parse session data: %s", data)
		}

		number, err := strconv.Atoi(matches[1])
		if err != nil {
			return nil, err
		}

		session[matches[2]] = number
	}

	return session, nil
}

func possibleGames(games []Game, limits ColorsSet) []Game {
	var compliantGames = make([]Game, 0)

LoopGame:
	for _, game := range games {
		// For each game
		for _, session := range game.Sessions {
			// And for each session
			for color, limit := range limits {
				// All limits are checked, if the limit value for a color is exceeded, the game is not possible
				if session[color] > limit {
					continue LoopGame
				}
			}
		}

		compliantGames = append(compliantGames, game)
	}

	return compliantGames
}
