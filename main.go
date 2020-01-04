package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

type Direction string

const(
	NORTH Direction = "N"
	EAST = "E"
	SOUTH = "S"
	WEST = "W"
)

type RoverPos struct {
	X int
	Y int
	D Direction
}

type Rover struct {
	RoverPos
}

func main() {
	var ymax, xmax int
	fmt.Scanf("%d %d", &xmax, &ymax)

	in := bufio.NewReader(os.Stdin)
	idx := 0

	var rover Rover
	for {
		s, err := in.ReadString('\n')
		if s == "\n" {
			break
		}
		if err != nil {
			// io.EOF is expected, anything else
			// should be handled/reported
			if err != io.EOF {
				log.Fatal(err)
			}
			break
		}
		if idx%2 == 0 {
			roverPos := strings.Split(strings.TrimSpace(s), " ")
			rover.RoverPos.X, _ = strconv.Atoi(roverPos[0])
			rover.RoverPos.Y,_ = strconv.Atoi(roverPos[1])
			rover.RoverPos.D = Direction(roverPos[2])
		} else {
			roverMovement := s
			rover.MoveRover(roverMovement)
			fmt.Println(rover.X, rover.Y, rover.D)
		}
		idx++
	}

}

func (rover *Rover)TurnLeft() {
	if rover.RoverPos.D == NORTH {
		rover.RoverPos.D = WEST
	} else if rover.RoverPos.D == WEST {
		rover.RoverPos.D = SOUTH
	} else if rover.RoverPos.D == SOUTH {
		rover.RoverPos.D = EAST
	}else if rover.RoverPos.D == EAST {
		rover.RoverPos.D = NORTH
	}
}

func (rover *Rover)TurnRight() {
	if rover.RoverPos.D == NORTH {
		rover.RoverPos.D = EAST
	}else if rover.RoverPos.D == EAST {
		rover.RoverPos.D = SOUTH
	}else if rover.RoverPos.D == SOUTH {
		rover.RoverPos.D = WEST
	}else if rover.RoverPos.D == WEST {
		rover.RoverPos.D = NORTH
	}
}

func (rover *Rover)Move() {
	if rover.RoverPos.D == NORTH {
		rover.RoverPos.Y++
	}else if rover.RoverPos.D == EAST {
		rover.RoverPos.X++
	}else if rover.RoverPos.D == SOUTH {
		rover.RoverPos.Y--
	}else if rover.RoverPos.D == WEST {
		rover.RoverPos.X--
	}
}

func (rover *Rover)MoveRover(movements string) {
	for _, movement := range movements{
		if movement == 'L'{
			rover.TurnLeft()
		}
		if movement == 'R'{
			rover.TurnRight()
		}
		if movement == 'M'{
			rover.Move()
		}
	}
}
