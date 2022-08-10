package model

import (
	"fmt"
	"strings"
)

type StoryNode struct {
	text    string
	choices *choices
}

type CreatNodeArgs struct {
	Text    string
	Choices []AddChoiceArgs
}

func NewNode(args CreatNodeArgs) *StoryNode {
	node := StoryNode{
		text: args.Text,
	}

	for _, choice := range args.Choices {
		node.AddChoice(choice)
	}

	return &node
}

type CreateFinalNodeArgs struct {
	Text string
}

func NewFinalNode(args CreateFinalNodeArgs) *StoryNode {
	newArgs := CreatNodeArgs{
		Text:    args.Text,
		Choices: []AddChoiceArgs{},
	}

	return NewNode(newArgs)
}

var storySeparator = "///////////////////////////////"

func (node *StoryNode) Play() {
	node.render()
	if node.choices != nil {
		var cmd string
		fmt.Scanln(&cmd)
		newNode := node.executeCmd(cmd)
		if newNode != nil {
			fmt.Println(storySeparator)
			newNode.Play()
		}
	}
}

func (node *StoryNode) AddChoice(args AddChoiceArgs) {
	choice := newChoice(args)

	if node.choices == nil {
		node.choices = choice
	} else {
		currentChoice := node.choices
		for currentChoice.nextChoice != nil {
			currentChoice = currentChoice.nextChoice
		}
		currentChoice.nextChoice = choice
	}
}

func (node *StoryNode) applyToChoice(fn func(choices *choices)) {
	currentChoice := node.choices

	for currentChoice != nil {
		fn(currentChoice)
		currentChoice = currentChoice.nextChoice
	}
}

func (node *StoryNode) render() {
	fmt.Println(node.text)
	node.applyToChoice(func(choices *choices) {
		fmt.Println(choices.cmd, ":", choices.description)
	})
}

func (node *StoryNode) executeCmd(cmd string) *StoryNode {
	var match *choices

	node.applyToChoice(func(choice *choices) {
		if strings.EqualFold(choice.cmd, cmd) {
			match = choice
		}
	})

	if match == nil {
		fmt.Println("Sorry, I did not understand")
		return node
	}

	return match.nextNode
}
