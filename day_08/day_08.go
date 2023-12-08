package day_08

import (
	"fmt"
	"strings"
)

type Node struct {
	Name  string
	Left  string
	Right string
}

func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

func CountStepsToDestinationNode(startNode Node, allNodes map[string]Node, directions string, allZ bool) int {
	steps := 0

	node := startNode
	reachedDestination := false
	for {
		directionToTake := string(directions[steps%len(directions)])

		var targetNodeName string
		switch directionToTake {
		case "L":
			targetNodeName = node.Left
		case "R":
			targetNodeName = node.Right
		}

		// fmt.Printf("Currently at node %s, following direction %s so next node will be %s\n", node.Name, directionToTake, targetNodeName)

		if allZ {
			if targetNodeName == "ZZZ" {
				reachedDestination = true
			}
		} else {
			if strings.HasSuffix(targetNodeName, "Z") {
				reachedDestination = true
			}
		}

		if reachedDestination {
			steps += 1
			break
		}

		steps += 1
		node = allNodes[targetNodeName]
	}

	return steps
}

func Solution(inputLines []string, partTwo bool) int {
	var directions string
	allNodes := map[string]Node{}

	for lineIndex, line := range inputLines {
		if lineIndex == 0 {
			directions = line
		}

		if lineIndex >= 2 {
			lineSplit := strings.Split(line, " = ")
			nodeName := lineSplit[0]
			targetNodesSplit := strings.Split(lineSplit[1][1:len(lineSplit[1])-1], ", ")
			allNodes[nodeName] = Node{Name: nodeName, Left: targetNodesSplit[0], Right: targetNodesSplit[1]}
		}
	}

	fmt.Printf("nodes maps:\n%v\n", allNodes)

	steps := 0
	if !partTwo {
		startNode := allNodes["AAA"]
		steps = CountStepsToDestinationNode(startNode, allNodes, directions, true)
	} else {
		var startNodes []Node
		stepsForNodes := []int{}
		for nodeName, node := range allNodes {
			if strings.HasSuffix(nodeName, "A") {
				startNodes = append(startNodes, node)
			}
		}

		for _, node := range startNodes {
			stepsForNode := CountStepsToDestinationNode(node, allNodes, directions, false)
			stepsForNodes = append(stepsForNodes, stepsForNode)
			fmt.Printf("From node %s to the first node ending in Z, there are %d steps\n", node.Name, stepsForNode)
		}

		// steps = 1
		// for _, stepCount := range stepsForNodes {
		// 	steps *= stepCount
		// }
		steps = LCM(stepsForNodes[0], stepsForNodes[1], stepsForNodes[2:]...)
	}

	fmt.Printf("Total steps from AAA to ZZZ are %d\n", steps)
	return steps
}
