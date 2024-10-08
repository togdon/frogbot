package responses

import (
	"fmt"
	"math/rand"
	"regexp"
	"time"
)

// MentionsResponse generates a response to a message that mentions the frog's user
func MentionsResponse(message string) string {

	// Create a new random number generator
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	// Look for the message to match the regex "what is best in life"
	bil := regexp.MustCompile(`(?i)what is best in life`)

	// If the message matches the regex then respond with:
	//To [verb] your [noun], to see them [pparticiple] [preposition] you, and
	//to [sense] the [singular] of their [kinfolk].
	if bil.MatchString(message) {

		verb := []string{"crush", "decimate", "demolish", "eradicate", "erase", "exterminate", "extinguish", "massacre", "obliterate", "quash", "quell", "raze", "slaughter", "smush", "tickle", "wipe out"}
		noun := []string{"adversaries", "antagonists", "assailants", "attackers", "competitors", "detractors", "enemies", "foes", "invaders", "opponents", "opposition", "rivals"}
		pparticiple := []string{"arisen", "awoken", "beaten", "bespoken", "bitten", "blown", "broken", "cast", "chosen", "drawn", "drawn", "driven", "eaten", "forgotten", "frozen", "hidden", "ridden", "shaken", "smitten", "thrust"}
		preposition := []string{"above", "across", "among", "at", "before", "beneath", "beside", "between", "by", "down", "in", "in front of", "over", "through", "with"}
		sense := []string{"hear", "see", "smell", "taste", "touch"}
		singular := []string{"lamentations", "complaints", "dirges", "elegies", "keenings", "laments", "moanings", "mournings", "requiems", "sobs", "tears", "ululations", "wails"}
		kinfolk := []string{"fuckbois", "women", "children", "cats", "dogs", "Great Aunt Mildred", "Crazy Uncle Ernie"}

		response := fmt.Sprintf("To %s your %s, to see them %s %s you, and to %s the %s of their %s.", verb[r1.Intn(len(verb))], noun[r1.Intn(len(noun))], pparticiple[r1.Intn(len(pparticiple))], preposition[r1.Intn(len(preposition))], sense[r1.Intn(len(sense))], singular[r1.Intn(len(singular))], kinfolk[r1.Intn(len(kinfolk))])

		return response
	} else {

		// Otherwise respond with a random line from Warcraft II
		replies := []string{"A noble quest.",
			"Ach?",
			"Ahoy!",
			"Are you opposing Shadowmoon?",
			"Are you still here?",
			"Are you still touching me?",
			"Arg?",
			"Awaiting orders.",
			"Aye, captain?",
			"Aye, matey?",
			"Aye?",
			"Bombs are great!",
			"By your command.",
			"Captain on the bridge.",
			"Do that again and you'll pull back a stump!",
			"Do you feel lucky, punk?",
			"Do you know who I am?",
			"Do you like fire?",
			"Do you need assistance?",
			"Don't anger me.",
			"Don't force me to run you through!",
			"Don't push it.",
			"Don't waste my time.",
			"Don't you have a kingdom to run?",
			"Enjoying yourself?",
			"Even elder races get tired of waiting.",
			"Exalted one?",
			"Excuse me!",
			"For the king!",
			"Glor'duk.",
			"Greetings!",
			"Hands off, grubber.",
			"Hee hee hee! That tickles.",
			"Hello?",
			"Here I am.",
			"Huh?",
			"I can see my house!",
			"I challenge you, heathen!",
			"I come to serve.",
			"I do have work to do.",
			"I got axe for you!",
			"I love blowin' things up!",
			"I wish I had a weapon!",
			"I would not do such things if I were you!",
			"I'd be delighted.",
			"I'd rather be sailing.",
			"I'll fling a hammer at you.",
			"I'm a busy frogbot.",
			"I'm alive!",
			"I'm at your service.",
			"I'm full of it!",
			"I'm growing impatient.",
			"I'm not listening.",
			"I'm not ready.",
			"I'm on it.",
			"I'm very busy.",
			"I'm waiting.",
			"I've got a flying machine!",
			"I've got the brain.",
			"In your name.",
			"It doesn't get any better than this!",
			"Job's done.",
			"'Join the army,' they said. 'See the world,' they said.",
			"KABOOOOOM!",
			"Leave me alone!",
			"Lok'tar.",
			"Look out!",
			"Make it quick.",
			"Make up your mind.",
			"Master?",
			"Milord?",
			"More work?",
			"My lord?",
			"My sovereign?",
			"My tummy feels funny.",
			"Need a hand?",
			"Need something?",
			"Now I'm hungry!",
			"Now what?",
			"Nuh-uh!",
			"Of course, my king.",
			"Of course.",
			"Of course... master...",
			"Oh, what?",
			"Orders, sire?",
			"Ready to serve, my lord.",
			"Ready to serve.",
			"Ready to work.",
			"Right away.",
			"Right-o.",
			"Say hello to my little friend!",
			"Set sail?",
			"Sire?",
			"Skipper?",
			"Slice and dice!",
			"Stop it.",
			"Stop rocking the boat!",
			"Stop that incessant clicking!",
			"Swobu.",
			"Take it.",
			"Tell me what to do.",
			"That doesn't hurt.",
			"This is the reason I ended it all.",
			"This way!",
			"Tilt one back with me, dog!",
			"Time is of the essence.",
			"Time to die!",
			"Under way.",
			"Ur-ur-ur!",
			"Very well.",
			"Want me to fly?",
			"We are under attack!",
			"We don't understand.",
			"We move.",
			"We must take action.",
			"We're being attacked!",
			"We're not brainless anymore.",
			"We're on our way.",
			"We're ready, master.",
			"We're ready.",
			"We've got explosives!",
			"Welcome to my nightmaaaare!",
			"Well?",
			"Whaaat?!",
			"Whaaat?",
			"What do you want?",
			"What ho!",
			"What is it?",
			"What you wan'me kill?",
			"What?",
			"Whatever you wish.",
			"When my work is finished, I'm coming back for you.",
			"Who is it?",
			"Who summoned me?",
			"Who wants to sing?",
			"Who's there?",
			"Why must you torment me?",
			"Yeah?",
			"Yeees?",
			"Yep.",
			"Yes, boss?",
			"Yes, Captain?",
			"Yes, lad?",
			"Yes, master?",
			"Yes, milord?",
			"Yes, my lord?",
			"Yes, sir?",
			"Yes, sire?",
			"Yes, what is it?",
			"Yes?",
			"You da boss.",
			"You got a belly full of haggis.",
			"You never touch the other bots like that.",
			"You talking to me?",
			"You think Lothar's death was my fault, don't you?",
			"You!",
			"You're making me seasick!",
			"You're the captain.",
			"Your command, master?",
			"Your command?",
			"Your Eminence?",
			"Your majesty?",
			"Your orders?",
			"Your request?",
			"Your wish?",
			"Y—uh-huh?",
			"Zoktok.",
			"Zug zug."}

		return replies[r1.Intn(len(replies))]
	}
}
