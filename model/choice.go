package model

type choices struct {
	cmd         string
	description string
	nextNode    *StoryNode
	nextChoice  *choices
}

type AddChoiceArgs struct {
	Cmd,
	Description string
	NextNode *StoryNode
}

func newChoice(args AddChoiceArgs) *choices {
	return &choices{
		cmd:         args.Cmd,
		description: args.Description,
		nextNode:    args.NextNode,
		nextChoice:  nil,
	}
}
