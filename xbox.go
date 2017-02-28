package xbox

import (
	"bufio"
	"fmt"
	"io"
	"os/exec"
	"strconv"
	"strings"
)

type analogStick struct {
	X int
	Y int
}

type dpad struct {
	Left  bool
	Right bool
	Up    bool
	Down  bool
}

type xbox struct {
	LeftStick    analogStick
	RightStick   analogStick
	Dpad         dpad
	LeftBumper   bool
	RightBumper  bool
	LeftTrigger  int
	RightTrigger int
	Start        bool
	Back         bool
	Guide        bool
	A            bool
	B            bool
	X            bool
	Y            bool
}

// Xbox holds the state of the xbox controller inputs
var Xbox xbox

var stdout io.ReadCloser

func Connect() {
	fmt.Println("Connecting to Xbox Controller...")

	cmd := exec.Command("xboxdrv")
	stdout, _ = cmd.StdoutPipe()

	// TODO: error checking

	cmd.Start()

	fmt.Println("Connected to Xbox Controller.")
}

func Control() {
	fmt.Println("Reading Xbox controller input...")
	scanner := bufio.NewScanner(stdout)

	for scanner.Scan() {
		var line = scanner.Text()
		// Length 139 indicates data line, ignore anything else
		if len(line) == 139 {
			Xbox.LeftStick.X, _ = strconv.Atoi(strings.Trim(line[3:9], " "))
			Xbox.LeftStick.Y, _ = strconv.Atoi(strings.Trim(line[13:19], " "))
			Xbox.RightStick.X, _ = strconv.Atoi(strings.Trim(line[24:30], " "))
			Xbox.RightStick.Y, _ = strconv.Atoi(strings.Trim(line[34:40], " "))
		}
	}

	fmt.Println("Stopped reading Xbox controller input.")
}
