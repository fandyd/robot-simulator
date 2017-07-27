package robot

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	//NORTH constant to be used for facing
	NORTH = "NORTH"
	//SOUTH constant to be used for facing
	SOUTH = "SOUTH"
	//EAST constant to be used for facing
	EAST = "EAST"
	//WEST constant to be used for facing
	WEST = "WEST"

	errRobotNotPlaced  = "Robot is not placed on the table yet."
	errRobotOverBoard  = "Command ignored, Robot will fall."
	errMovementIgnored = "Invalid Movement so it is ignored."
	errInvalidCmd      = "Sorry i dont get that!"
	errFailToInitiate  = "Failed to initiate"
)

//Robot hold all the information about the robot
type Robot struct {
	X, Y          int
	F             string
	IsRobotPlaced bool
	Table         Table
}

//Perform is to receive a command and perform the command
func (r *Robot) Perform(cmd string) (string, error) {
	var err error
	command, x, y, f := parseCommand(cmd)
	switch command {
	case "PLACE":
		err = r.Place(x, y, f)
	case "MOVE":
		err = r.Move()
	case "LEFT":
		err = r.Left()
	case "RIGHT":
		err = r.Right()
	case "REPORT":
		return r.Report()
	default:
		return "", fmt.Errorf(errInvalidCmd)
	}
	if err != nil {
		return "", err
	}
	return "Command Executed Successfully", nil
}

//Place will put the toy robot on the table in position X,Y and facing NORTH, SOUTH, EAST or WEST.
func (r *Robot) Place(x int, y int, f string) error {
	direction := strings.ToUpper(f)
	//IF its not on the table
	if !isStillOnTheTable(x, y, r.Table) {
		return fmt.Errorf(errRobotOverBoard)
	}
	if !isValidFacing(direction) {
		return fmt.Errorf(errInvalidCmd)
	}
	r.X = x
	r.Y = y
	r.F = direction
	r.IsRobotPlaced = true
	return nil
}

//Move will move the toy robot one unit forward in the direction it is currently facing.
func (r *Robot) Move() error {
	if !r.IsRobotPlaced {
		return fmt.Errorf(errRobotNotPlaced)
	}
	switch r.F {
	case NORTH:
		//IF after move and still in the table
		if isStillOnTheTable(r.X, r.Y+1, r.Table) {
			r.Y++
			return nil
		}
	case SOUTH:
		//IF after move and still in the table
		if isStillOnTheTable(r.X, r.Y-1, r.Table) {
			r.Y--
			return nil
		}
	case EAST:
		//IF after move and still in the table
		if isStillOnTheTable(r.X+1, r.Y, r.Table) {
			r.X++
			return nil
		}
	case WEST:
		//IF after move and still in the table
		if isStillOnTheTable(r.X-1, r.Y, r.Table) {
			r.X--
			return nil
		}
	}
	return fmt.Errorf(errMovementIgnored)
}

//Left will rotate the robot 90 degrees to left in the specified direction without changing the position of the robot.
func (r *Robot) Left() error {
	if !r.IsRobotPlaced {
		return fmt.Errorf(errRobotNotPlaced)
	}
	switch r.F {
	case NORTH:
		r.F = WEST
	case SOUTH:
		r.F = EAST
	case EAST:
		r.F = NORTH
	case WEST:
		r.F = SOUTH
	}
	return nil
}

//Right will rotate the robot 90 degrees to right in the specified direction without changing the position of the robot.
func (r *Robot) Right() error {
	if !r.IsRobotPlaced {
		return fmt.Errorf(errRobotNotPlaced)
	}
	switch r.F {
	case NORTH:
		r.F = EAST
	case SOUTH:
		r.F = WEST
	case EAST:
		r.F = SOUTH
	case WEST:
		r.F = NORTH
	}
	return nil
}

//Report will announce the X,Y and F of the robot. This can be in any form, but standard output is sufficient
func (r *Robot) Report() (string, error) {
	if !r.IsRobotPlaced {
		return "", fmt.Errorf(errRobotNotPlaced)
	}
	return fmt.Sprintf("%d,%d,%s", r.X, r.Y, r.F), nil
}

//NewRobot is to instantiate a new robot object
func NewRobot(tableWidth int, tableLength int) (Robot, error) {
	table, err := NewTable(tableWidth, tableLength)
	if err != nil {
		return Robot{}, fmt.Errorf(errFailToInitiate)
	}
	return Robot{Table: table}, nil
}

//private methods

func isStillOnTheTable(x int, y int, table Table) bool {
	return x < table.Width && x >= 0 && y < table.Length && y >= 0
}

func isValidFacing(f string) bool {
	return f == NORTH || f == SOUTH || f == EAST || f == WEST
}

func parseCommand(line string) (string, int, int, string) {
	var f, command string
	var x, y int
	var args []string
	commands := strings.Split(strings.ToUpper(line), " ")
	command = commands[0]
	if len(commands) == 2 {
		args = strings.Split(commands[1], ",")
	}
	if len(args) == 3 {
		x, _ = strconv.Atoi(args[0])
		y, _ = strconv.Atoi(args[1])
		f = args[2]
	}
	return command, x, y, f
}
