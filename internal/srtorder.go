package internal

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

type ParsedTimestamp struct {
	Start time.Duration
	End   time.Duration
}

type Chunk struct {
	SeqNumber       string
	Timestamp       string
	ParsedTimestamp ParsedTimestamp
	Content         []string
}

type ByStart []Chunk

func (s ByStart) Len() int {
	return len(s)
}
func (s ByStart) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s ByStart) Less(i, j int) bool {
	return s[i].ParsedTimestamp.Start < s[j].ParsedTimestamp.Start
}

type Position string

const (
	SeqNumber Position = "SeqNumber"
	Timestamp Position = "Timestamp"
	Content   Position = "Content"
)

func Process(filepath string) error {
	file, err := os.Open(filepath)
	if err != nil {
		return err
	}
	defer file.Close()

	chunks := make([]Chunk, 0)
	scanner := bufio.NewScanner(file)

	var c Chunk
	currPosition := SeqNumber
	for scanner.Scan() {
		switch currPosition {
		case SeqNumber:
			{
				c = Chunk{}
				c.Content = make([]string, 0)

				c.SeqNumber = scanner.Text()
				currPosition = Timestamp
				continue
			}
		case Timestamp:
			{
				c.Timestamp = scanner.Text()

				sep := " --> "
				s := strings.Split(scanner.Text(), sep)
				if len(s) != 2 {
					return fmt.Errorf("broken line: '%s'", scanner.Text())
				}

				parsedFrom, err := parseTime(s[0])
				if err != nil {
					return err
				}

				parsedUntil, err := parseTime(s[0])
				if err != nil {
					return err
				}

				c.ParsedTimestamp.Start = parsedFrom
				c.ParsedTimestamp.End = parsedUntil
				currPosition = Content
				continue
			}
		case Content:
			{
				// TODO: check if blank line
				if len(scanner.Text()) <= 0 {
					// commit
					chunks = append(chunks, c)
					currPosition = SeqNumber
				} else {
					c.Content = append(c.Content, scanner.Text())
				}
			}
		}
	}

	// Sort based on Start
	sort.Sort(ByStart(chunks))

	f, err := os.CreateTemp("", "")
	if err != nil {
		return err
	}

	w := bufio.NewWriter(f)
	for i, c := range chunks {
		c.SeqNumber = strconv.Itoa(i + 1)

		w.WriteString(c.SeqNumber + "\n")
		w.WriteString(c.Timestamp + "\n")
		for _, content := range c.Content {
			w.WriteString(content + "\n")
		}
		w.WriteString("\n")
	}
	w.Flush()

	return nil
}
