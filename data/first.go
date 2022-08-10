package data

import (
	"github.com/jaroddev/textadventure/model"
)

func CreateStory() *model.StoryNode {
	return firstday()
}

//  - wake up
//  - wp => freezing (BE) + gorilla (BE) + warmdoor
func firstday() *model.StoryNode {
	next := warmDoorStory()

	doNothing := model.AddChoiceArgs{
		Cmd:         "none",
		Description: "You stay here, doing nothing",
		NextNode:    model.NewFinalNode(BadEnd().Why()),
	}

	coldDoor := model.AddChoiceArgs{
		Cmd:         "cold",
		Description: "You go to the cold door",
		NextNode:    model.NewFinalNode(BadEnd().Gorilla()),
	}

	warmDoor := model.AddChoiceArgs{
		Cmd:         "warm",
		Description: "You go to the warm door",
		NextNode:    model.NewNode(next),
	}

	firstStory := model.CreatNodeArgs{
		Text: `You are in a large chamber, deep underground.
		You are freezing cold and in a near death state.
		You can see two doors:
		 - behind the first one, you feel cold air 
		 - behind the second one, you feel warm air
		As you tremble, and begin to loose your conciousness,
		you take a decision.		
		`,
		Choices: []model.AddChoiceArgs{
			doNothing,
			coldDoor,
			warmDoor,
		},
	}

	adv := model.NewNode(firstStory)
	return adv
}

//  - wd => bath + walk
func warmDoorStory() model.CreatNodeArgs {
	takeABath := model.AddChoiceArgs{
		Cmd:         "dip",
		Description: "You take a bath",
		NextNode:    model.NewNode(bathStory()),
	}

	walk := model.AddChoiceArgs{
		Cmd:         "walk",
		Description: "You don't want to take a bath so you continue walking your way",
		NextNode:    model.NewNode(walk()),
	}

	story := model.CreatNodeArgs{
		Text: `Being cold, naturally, you end up being attracted to the warm
		and sweet air your skin feels, in hope that you will be saved.
		
		You touch the square-shaped knob and turn it from left to right, it is not opening
		you hurry and turn it from right to left this time, knowing that you might die
		anytime soon.
		
		The door finally opens, you walk a few step, 
		the stone you walk on is hot and confortable, your lungs are heating up too.

		Tearing up, you continue walking and see in front of you a stone panel,

		<<Please take a dip in the hot thermal bath to recover>>
		`,
		Choices: []model.AddChoiceArgs{
			takeABath,
			walk,
		},
	}

	return story
}

//  - bath => rest a bit more and sleep (BE) + leave and hear a sound
func bathStory() model.CreatNodeArgs {
	sleep := model.AddChoiceArgs{
		Cmd:         "sleep",
		Description: "You sleep",
		NextNode:    model.NewFinalNode(BadEnd().BathSleep()),
	}

	leave := model.AddChoiceArgs{
		Cmd:         "leave",
		Description: "You leave the bath to see where the sound is from",
		NextNode:    model.NewNode(sound()),
	}

	story := model.CreatNodeArgs{
		Text: `A confortable sensation runs through your whole body.
		The previously tensed and rigid muscles gradually soften
		while your mind begin to finally work.

		Who am i ? where is this ? Why was i here ? alone ?
		What should i do ?

		You realize that you have a lot to think about. what is the first
		thing to do...

		You still are not in the best state yourself, waiting a bit to rest
		to make sure you are healed and ready to go could be an option.

		You are hungry and you found no food, you don't know where you are
		so exploring this place might be a good idea.
		
		Deep in your thought, you grab your chin with your right hand,
		examine the room you are in.
		
		Both options seems to be important and could be done right now.
		But the bath you are in is exquisite, the thought of leaving affect your
		thoughts and you decide to rest and sleep in the bath.
		
		rrrrrr, tch tch tch...
		
		You hear a sound, you glance to you right... to your left...
		Nothing, you wait for a minute, still nothing.
		
		Was it your mind playing you a tour ? Should i leave ?`,
		Choices: []model.AddChoiceArgs{
			sleep,
			leave,
		},
	}

	return story
}

//  - walk => hear a sound * bonus
func walk() model.CreatNodeArgs {
	story := model.CreatNodeArgs{
		Text: `The sensations in your limbs, begin to come back.
		It was as if, thousands of ants, that ran in your limbs after
		suddenly awaking.

		But still, you understand that the worst situation passed,
		you feel alive again.

		Kkkk Bump...

		It hurts, a wall, in front of you, you probably bumped into it,
		you notice that something was written on it:

		<< Don't forget that ====D don't fight the tiger !!! >>.

		This makes no sense, what did it mean ? you decide to not forget it,
		at least for now... 
		
		Well nothing really hard for you, a person whose memory was gone.

		Who am i ? where is this ? Why was i here ? alone ?
		What should i do ?

		You realize that you have a lot to think about. what is the first
		thing to do...
		
		Deep in your thought, you grab your chin with your right hand
		and examine the room you are in.
		
		rrrrrr, tch tch tch...
		
		You hear a sound, you glance to you right... to your left...
		Nothing, you wait for a minute, still nothing.

		RRRRRRRRRRRr!!!!!! woof! woof!!

		a dog ?

		You run to his direction`,
		Choices: []model.AddChoiceArgs{
			{
				Cmd:         "go",
				Description: "You have no choice other than going there",
				NextNode:    model.NewNode(sound()),
			},
		},
	}

	return story
}

//  - sound => (take the weapon * kill dog * remember) eat the dog or eat nothing (BE)
func sound() model.CreatNodeArgs {
	return model.CreatNodeArgs{
		Text:    `You see a`,
		Choices: []model.AddChoiceArgs{},
	}
}
