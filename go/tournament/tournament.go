package tournament

import (
	"bufio"
	"fmt"
	"io"
	"sort"
	"strings"
)

func Tally(in io.Reader, out io.Writer) error {
	teams, err := readRecords(in)
	if err != nil {
		return err
	}

	out.Write([]byte("Team                           | MP |  W |  D |  L |  P\n"))
	for _, t := range ranked(teams) {
		out.Write([]byte(t.String()))
	}

	return nil
}

func readRecords(in io.Reader) (map[string]*record, error) {
	teams := make(map[string]*record)
	scanner := bufio.NewScanner(in)
	for scanner.Scan() {
		text := scanner.Text()
		if len(text) == 0 || text[0] == '#' {
			continue
		}

		tokens := strings.Split(scanner.Text(), ";")
		if len(tokens) != 3 {
			return nil, fmt.Errorf("Invalid record: %v", tokens)
		}

		switch tokens[2] {
		case "win":
			match(teams, tokens[0]).won()
			match(teams, tokens[1]).lost()
		case "loss":
			match(teams, tokens[0]).lost()
			match(teams, tokens[1]).won()
		case "draw":
			match(teams, tokens[0]).tied()
			match(teams, tokens[1]).tied()
		default:
			return nil, fmt.Errorf("Invalid match result")
		}
	}
	return teams, nil
}

func ranked(records map[string]*record) []*record {
	teams := make([]*record, 0, len(records))
	for _, v := range records {
		teams = append(teams, v)
	}
	sort.Slice(teams, func(a, b int) bool {
		if teams[a].points > teams[b].points {
			return true
		}
		if teams[a].points < teams[b].points {
			return false
		}
		return strings.Compare(teams[a].name, teams[b].name) < 0
	})

	return teams
}

func match(teams map[string]*record, name string) *record {
	team, found := teams[name]
	if !found {
		team = &record{name: name}
		teams[name] = team
	}
	return team
}

type record struct {
	name   string
	wins   int
	losses int
	draws  int
	points int
	played int
}

func (r *record) won() {
	r.wins++
	r.points += 3
	r.played++
}

func (r *record) lost() {
	r.losses++
	r.played++
}

func (r *record) tied() {
	r.draws++
	r.points++
	r.played++
}

func (r *record) String() string {
	return fmt.Sprintf("%-30s |%3d |%3d |%3d |%3d |%3d\n",
		r.name, r.played, r.wins, r.draws, r.losses, r.points)
}
