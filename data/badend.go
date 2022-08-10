package data

import "github.com/jaroddev/textadventure/model"

func BadEnd() end {
	return end{}
}

type end struct{}

// First bad end
func (end end) Why() model.CreateFinalNodeArgs {
	story := model.CreateFinalNodeArgs{
		Text: `Your whole body feels gradually colder and colder.
		Then suddenly, your eyes began to feel really heavy.
		
		You died... END`,
	}

	return story
}

// Cold door
func (end end) Gorilla() model.CreateFinalNodeArgs {
	story := model.CreateFinalNodeArgs{
		Text: `Being cold, naturally, you end up being attracted to the warm
		and sweet air your skin feels, in hope that you will be saved.
		
		huh, wait a minute, for what not reason, you go to the cold door.
		You push the frozen door with all the strenth that remains, the door opens.

		You take your first step in what seems to be snow, the air is freezing a lot more
		than your former place, your skin began to seem like burning, and you don't have any strength 
		remaining but you manage to turn around.
		
		You stop, you hear a horryfying sound from above your head.

		A giant gorilla like creature, is looking at you, atop what seems to be 
		the rooftop of the same you seem to came from.

		Soon, the landscape you see changes, there is red snow everywhere,
		but that is not the only things that changed, you are aching from everywhere,
		and the ground is next your head...

		You manage to use the last bit of strength to look toward your body.

		AHHHHHHHHHHHHHHHHHHHH, there is a hole, ah ah ah ah, ghhhh
		after a quick scream, and a single phrase, as if reporting your state to someone.

		You take your last breath... END
		`,
	}

	return story
}

// sleep in the bath
func (end end) BathSleep() model.CreateFinalNodeArgs {
	story := model.CreateFinalNodeArgs{
		Text: `rrrrrrrrrr...

		RRRRRRRRR...

		you hear something but your sleep, while not deep,
		still prevents you from understanding the danger you are in.

		RRRRRRRRRRRRRRRRRRRR !! Woof! Woof!!

		The sounds woke you up, what's that !
		A dog ? you look around you in a surprise, a dog, in front of you.

		With a bloody appearance and sharp fangs, the more he comes
		in your direction, the bigger he appears to be.
		
		Afraid and without proper way to defend yourself, you decide to run
		you leave the bath, run it the oposite direction at your best speed,
		you look behind to check the dog's behavior.
		
		Woof !!!! Woof !!!

		It runs at full speed in your direction, faster than you are,
		the gap quickly lessen, you try to speed up but you are already at full speed.

		Your feet begin wet from the bath, you slip, face to the ground,
		when you try to get up, you see blood coming from your dripping node.

		RRRRRRRRRRR!

		A shadow, whose ? you already know it, but you still have to check.
		You turn your head behind you, a huge black dog, opens his mouth,
		his saliva falls on top of your head.
		
		Woof ! Woof !!

		You don't move, you want but your body is not moving, the dog
		jump on your throat and rip it...

		The last thing you see, the dog dragging your body somewhere... END
		`,
	}

	return story
}

func (end end) templateEnd() model.CreateFinalNodeArgs {
	story := model.CreateFinalNodeArgs{
		Text: `
		You take your last breath... END
		`,
	}

	return story
}
