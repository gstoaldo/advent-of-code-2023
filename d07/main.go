package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/gstoaldo/advent-of-code-2023/utils"
)

var cardRanks = []rune{'A', 'K', 'Q', 'J', 'T', '9', '8', '7', '6', '5', '4', '3', '2'}

func getCardRanksMap() map[rune]int {
	result := map[rune]int{}

	for i := 0; i < len(cardRanks); i++ {
		result[cardRanks[i]] = len(cardRanks) - i
	}

	return result
}

var cardRanksMap = getCardRanksMap()

func parse(filepath string) ([]string, []int) {
	lines := utils.ReadLines(filepath)

	hands := []string{}
	bids := []int{}

	for _, line := range lines {
		parts := strings.Split(line, " ")
		hands = append(hands, parts[0])
		bid, _ := strconv.Atoi(parts[1])
		bids = append(bids, bid)
	}

	return hands, bids
}

type handMapT map[rune]int

func handMap(hand string) handMapT {
	result := handMapT{}
	for _, c := range hand {
		result[c] += 1
	}
	return result
}

func _5ok(handMap handMapT) bool {
	return len(handMap) == 1
}

func _4ok(handMap handMapT) bool {
	found := false
	for _, v := range handMap {
		found = found || v == 4
	}

	return len(handMap) == 2 && found
}

func _fh(handMap handMapT) bool {

	found := false
	for _, v := range handMap {
		found = found || v == 3
	}

	return len(handMap) == 2 && found
}

func _3ok(handMap handMapT) bool {
	found := false
	for _, v := range handMap {
		found = found || v == 3
	}

	return len(handMap) == 3 && found
}

func _2p(handMap handMapT) bool {
	found := false
	for _, v := range handMap {
		found = found || v == 2
	}

	return len(handMap) == 3 && found
}

func _1p(handMap handMapT) bool {
	return len(handMap) == 4
}

func _hc(handMap handMapT) bool {
	return len(handMap) == 5
}

func handRank(hand string) []int {
	types := []func(handMapT) bool{
		_5ok, _4ok, _fh, _3ok, _2p, _1p, _hc,
	}
	_handMap := handMap(hand)

	rank := []int{}

	for i := 0; i < len(types); i++ {
		if types[i](_handMap) {
			rank = append(rank, len(types)-i)
			break
		}
	}

	for _, r := range hand {
		rank = append(rank, cardRanksMap[r])
	}

	return rank
}

func isStronger(handA, handB string) bool {
	handARank := handRank(handA)
	handBRank := handRank(handB)

	for i := range handARank {
		if handARank[i] == handBRank[i] {
			continue
		}

		return handARank[i] > handBRank[i]
	}

	panic("hands are equal")
}

func part1(hands []string, bids []int) int {
	handToBid := map[string]int{}

	for i, hand := range hands {
		handToBid[hand] = bids[i]
	}

	sort.SliceStable(hands, func(i, j int) bool {
		return isStronger(hands[i], hands[j])
	})

	sum := 0

	for i, hand := range hands {
		sum += (len(hands) - i) * handToBid[hand]
	}

	return sum
}

func main() {
	hands, bids := parse(utils.Filepath())
	fmt.Println(part1(hands, bids))
}
