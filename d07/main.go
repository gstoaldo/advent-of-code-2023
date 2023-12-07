package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/gstoaldo/advent-of-code-2023/utils"
)

var cardRanksP1 = []rune{'A', 'K', 'Q', 'J', 'T', '9', '8', '7', '6', '5', '4', '3', '2'}
var cardRanksP2 = []rune{'A', 'K', 'Q', 'T', '9', '8', '7', '6', '5', '4', '3', '2', 'J'}

func getCardRanksMap(cardRanks []rune) map[rune]int {
	result := map[rune]int{}

	for i := 0; i < len(cardRanks); i++ {
		result[cardRanks[i]] = len(cardRanks) - i
	}

	return result
}

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

func typeRank(hand string) int {
	types := []func(handMapT) bool{
		_5ok, _4ok, _fh, _3ok, _2p, _1p, _hc,
	}
	_handMap := handMap(hand)

	for i := 0; i < len(types); i++ {
		if types[i](_handMap) {
			return len(types) - i
		}
	}

	panic("type rank not found")
}

func handRank(hand string, cardRanksMap map[rune]int) []int {
	rank := []int{}
	rank = append(rank, typeRank(hand))

	for _, r := range hand {
		rank = append(rank, cardRanksMap[r])
	}

	return rank
}

func isStronger(handA, handB string) bool {
	cardRanksMap := getCardRanksMap(cardRanksP1)
	handARank := handRank(handA, cardRanksMap)
	handBRank := handRank(handB, cardRanksMap)

	for i := range handARank {
		if handARank[i] == handBRank[i] {
			continue
		}

		return handARank[i] > handBRank[i]
	}

	panic("hands are equal")
}

func isStrongerWithJoker(handA, handB string) bool {
	cardRanksMap := getCardRanksMap(cardRanksP2)
	handARank := handRank(handA, cardRanksMap)
	handBRank := handRank(handB, cardRanksMap)

	handARank[0] = maxPossible(handA)
	handBRank[0] = maxPossible(handB)

	for i := range handARank {
		if handARank[i] == handBRank[i] {
			continue
		}

		return handARank[i] > handBRank[i]
	}

	panic("hands are equal")
}

func maxPossible(hand string) int {
	max := 0

	for _, c := range cardRanksP2 {
		max = utils.Max(max, typeRank(strings.Replace(hand, "J", string(c), -1)))
	}

	return max
}

func totalWinnings(hands []string, bids []int, less func(i, j string) bool) int {
	handToBid := map[string]int{}

	for i, hand := range hands {
		handToBid[hand] = bids[i]
	}

	handsSorted := append([]string{}, hands...)

	sort.SliceStable(handsSorted, func(i, j int) bool {
		return less(handsSorted[i], handsSorted[j])
	})

	sum := 0
	for i, handsSorted := range handsSorted {
		sum += (len(hands) - i) * handToBid[handsSorted]
	}

	return sum
}

func part1(hands []string, bids []int) int {
	return totalWinnings(hands, bids, isStronger)
}

func part2(hands []string, bids []int) int {
	return totalWinnings(hands, bids, isStrongerWithJoker)
}

func main() {
	hands, bids := parse(utils.Filepath())
	fmt.Println(part1(hands, bids))
	fmt.Println(part2(hands, bids))
}
