package input

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type SysInputProvider struct {
}

func (s *SysInputProvider) TakeInput() PlayerInput {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Please provide your input:\n")

	line, _ := reader.ReadString('\n')
	splitLine := strings.Split(line, " ")

	var err error
	var playerId int
	var targetX int
	var targetY int

	if playerId, err = strconv.Atoi(splitLine[0]); err == nil {
		fmt.Printf("i=%d, type: %T\n", playerId, playerId)
	}
	if targetX, err = strconv.Atoi(splitLine[1]); err == nil {
		fmt.Printf("i=%d, type: %T\n", targetX, targetX)
	}
	if targetY, err = strconv.Atoi(splitLine[2]); err == nil {
		fmt.Printf("i=%d, type: %T\n", targetY, targetY)
	}

	return PlayerInput{playerId, targetX, targetY}
}
