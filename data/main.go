// Contains the data to create the game
// Each story from the scenario/storynode is created here
// Contains the scenario by day and the bad end are separated from all the scenario

// A game where the player has to survive.
// The player controls a man, with no memory, recovers with time.

// Former hunter who came to save a person
// and kill the gorilla

// Meet no one but some beast (dog, snow tiger, snow bear, snake, gorilla)
// And dead persons, from whom he gets clues to who he is (memory artifacts and weapon).

// One or two good routes
// A dozen of bad ends

// Routes:
//  - wake up
//  - wp => freezing (BE) + gorilla (BE) + warmdoor
//  - wd => bath + walk
//  - bath => rest a bit more and sleep (BE) + leave and hear a sound
//  - walk => hear a sound * bonus
//  - sound => (take the weapon * kill dog * remember) eat the dog or eat nothing (BE)

//  - second wake up => left + right
//  - left => dead person * take weapon * read journal * key
//  - right => not unlocked * has to go to left
//  - right second => coat * new weapon * boot
//  - coat => sleep + bath
//  - bath => snake 1 (BE)

//  - sleep => wakeup and decide to leave this place and meet bear (cold door)
//  - cd => run + bear 1 (BE)
//  - run => see a house * enters it * read journal * eat final dog * gun * sleep

//  - sleep * encounter tiger => shoot + not shoot
//  - shoot => avalanche (BE)
//  - not shoot => mansion * bear
//  - bear => shoot + not shoot
//  - not shoot => bear 2 (BE)
//  - shoot => eat bear * enter mansion
//  - mansion => read journal * sleep * code and map

//  - sleep => go to the place thanks to the map
//  - place => look + not look
//  - not look => no bell
//  - look => bell
//  - no bell => gorilla fight with gun
//  - bell => gorilla fight with gun + gorilla ring bell
//  - gun => gori kill person * gori die
//  - bell on gori => person saved * gori die (Good End TRUE)
//  - person dir => despair (BE) + burial (good end 2)

// Ends:
// Meet the person dead but kills the gorilla and survive
// Meet the person alive and kill the gorilla
// Dies => bad end

// Bad ends:
// Nothing
// gorilla 1
// (tired) then killed by a dog in the bath
// hunger
// snake 1
// bear 1
// avalanche (tiger seems to be not harmful)
// bear 2
// despair and deathly self harm
// gorilla 2
package data
