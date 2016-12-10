package main

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/jbarratt/goadvent2016/advent"
)

var botProgram = regexp.MustCompile(`bot (\d+) gives low to (output|bot) (\d+) and high to (output|bot) (\d+)`)
var valueDelivery = regexp.MustCompile(`value (\d+) goes to bot (\d+)`)

// Chip represents a chip ID
type Chip int

// Bot represents a bot
type Bot struct {
	id      int
	chips   []Chip
	lowVal  Receiver
	highVal Receiver
}

// Deliver gives a chip to a bot
func (b *Bot) Deliver(chip Chip) {
	b.chips = append(b.chips, chip)
	b.Distribute()
}

// Distribute is run when changes are made to the bot.
// If it is ready to distribute, it does.
// - Both chips in hand
// - Both Receivers non-nil
func (b *Bot) Distribute() {
	if len(b.chips) == 2 && b.lowVal != nil && b.highVal != nil {
		fmt.Printf("Bot %d comparing values %d and %d\n", b.id, int(b.chips[0]), int(b.chips[1]))
		fmt.Printf("%s", b.Status())
		if b.chips[0] < b.chips[1] {
			b.lowVal.Deliver(b.chips[0])
			b.highVal.Deliver(b.chips[1])
		} else {
			b.lowVal.Deliver(b.chips[1])
			b.highVal.Deliver(b.chips[0])
		}
		b.chips = []Chip{}
	}
}

// Deliver gives a chip to an output bin
func (b *Bin) Deliver(chip Chip) {
	b.chips = append(b.chips, chip)
}

// Status Return a string version of the status
func (b *Bin) Status() string {
	return fmt.Sprintf("Bin %d: Chips: %v\n", b.id, b.chips)
}

// Status Return a string version of the status
func (b *Bot) Status() string {
	return fmt.Sprintf("Bot %d: Chips: %v lowVal: %v highVal %v\n", b.id, b.chips, b.lowVal, b.highVal)
}

// Bin is a bin
type Bin struct {
	id    int
	chips []Chip
}

// Receiver is implemented by bin and Bot
// Allows bots to be programmed to dispatch to either
type Receiver interface {
	Deliver(Chip)
	Status() string
}

func getBot(botID int) *Bot {
	if bot, ok := bots[botID]; ok {
		return bot
	}
	bot := &Bot{}
	bot.chips = make([]Chip, 0)
	bot.id = botID
	bots[botID] = bot
	return bot
}

func getBin(binID int) *Bin {
	if bin, ok := bins[binID]; ok {
		return bin
	}
	bin := &Bin{}
	bin.id = binID
	bin.chips = make([]Chip, 0)
	bins[binID] = bin
	return bin
}

func dispatchProgram(cmd string) {
	matches := valueDelivery.FindStringSubmatch(cmd)
	if len(matches) > 0 {
		chipID, _ := strconv.Atoi(matches[1])
		botID, _ := strconv.Atoi(matches[2])
		bot := getBot(botID)
		bot.Deliver(Chip(chipID))
		return
	}
	matches = botProgram.FindStringSubmatch(cmd)
	if len(matches) > 0 {
		botID, _ := strconv.Atoi(matches[1])
		lowID, _ := strconv.Atoi(matches[3])
		highID, _ := strconv.Atoi(matches[5])
		bot := getBot(botID)
		if matches[2] == "bot" {
			bot.lowVal = getBot(lowID)
		} else {
			bot.lowVal = getBin(lowID)
		}
		if matches[4] == "bot" {
			bot.highVal = getBot(highID)
		} else {
			bot.highVal = getBin(highID)
		}
		// If a bot has a payload and was just waiting for a program
		// Allow it to deliver the chips
		bot.Distribute()
	}
}

var bots = map[int]*Bot{}
var bins = map[int]*Bin{}

func printStatus() {
	for _, bot := range bots {
		fmt.Printf("%s", bot.Status())
	}
}

func main() {
	lines, err := advent.ReadFile(advent.FilenameArg())
	if err != nil {
		panic(err)
	}
	for _, line := range lines {
		fmt.Printf("  > %s\n", line)
		dispatchProgram(line)
	}
	secretCode := int(getBin(0).chips[0])
	for binID := 1; binID < 3; binID++ {
		secretCode *= int(getBin(binID).chips[0])
	}
	fmt.Printf("Secret code: %d\n", secretCode)
}
