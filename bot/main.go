package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"regexp"
	"strings"
	"syscall"
	"time"

	"github.com/bwmarrin/discordgo"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

var frogs []string

func main() {

	token := os.Getenv("BOT_TOKEN")
	if token == "" {
		panic("No token provided.")
	}

	// Create a new Discord session using the provided bot token.
	bot, err := discordgo.New("Bot " + token)
	if err != nil {
		panic(fmt.Sprintf("Error creating Discord session: %v", err))
	}

	// We need information about guilds and messages
	bot.Identify.Intents = discordgo.IntentsAll

	// Open the websocket and begin listening
	err = bot.Open()
	if err != nil {
		panic(fmt.Sprintf("Error opening Discord session: %v", err))
	}

	files, err := os.ReadDir("frogs/")
	if err != nil {
		panic(fmt.Sprintf("Error opening folder: %v", err))
	}

	for _, file := range files {
		frogs = append(frogs, file.Name())
	}

	bot.AddHandler(messageCreate)

	// Wait here until CTRL-C or other term signal is received.
	fmt.Printf("Frogbot is now running; I currently know about %v frogs. Press CTRL-C to exit.\n", len(frogs))
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	bot.Close()
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore all messages created by the bot itself
	if m.Author.ID == s.State.User.ID {
		return
	}

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

	for _, user := range m.Mentions {
		if user.ID == s.State.User.ID {

			s1 := rand.NewSource(time.Now().UnixNano())
			r1 := rand.New(s1)

			bil := regexp.MustCompile(`(?i)what is best in life`)
			if bil.MatchString(m.Content) {

				//To [verb] your [noun], to see them [pparticiple] [preposition] you, and
				//to [sense] the [singular] of their [kinfolk].

				verb := []string{"crush", "decimate", "demolish", "eradicate", "erase", "exterminate", "extinguish", "massacre", "obliterate", "quash", "quell", "raze", "slaughter", "smush", "tickle", "wipe out"}
				noun := []string{"adversaries", "antagonists", "assailants", "attackers", "competitors", "detractors", "enemies", "foes", "invaders", "opponents", "opposition", "rivals"}
				pparticiple := []string{"arisen", "awoken", "beaten", "bespoken", "bitten", "blown", "broken", "cast", "chosen", "drawn", "drawn", "driven", "eaten", "forgotten", "frozen", "hidden", "ridden", "shaken", "smitten", "thrust"}
				preposition := []string{"above", "across", "among", "at", "before", "beneath", "beside", "between", "by", "down", "in", "in front of", "over", "through", "with"}
				sense := []string{"hear", "see", "smell", "taste", "touch"}
				singular := []string{"lamentations", "complaints", "dirges", "elegies", "keenings", "laments", "moanings", "mournings", "requiems", "sobs", "tears", "ululations", "wails"}
				kinfolk := []string{"fuckbois", "women", "children", "cats", "dogs", "Great Aunt Mildred", "Crazy Uncle Ernie"}
				response := fmt.Sprintf("To %s your %s, to see them %s %s you, and to %s the %s of their %s.", verb[r1.Intn(len(verb))], noun[r1.Intn(len(noun))], pparticiple[r1.Intn(len(pparticiple))], preposition[r1.Intn(len(preposition))], sense[r1.Intn(len(sense))], singular[r1.Intn(len(singular))], kinfolk[r1.Intn(len(kinfolk))])

				s.ChannelMessageSend(m.ChannelID, response)
			} else {
				s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("%s", replies[r1.Intn(len(replies))]))
			}
		}
	}

	fmt.Printf("%v (%v) wrote: %v\n", m.Author, m.Author.ID, m.Content)

	frogme := regexp.MustCompile(`(?i)Frog me`)

	if frogme.MatchString(m.Content) {
		s1 := rand.NewSource(time.Now().UnixNano())
		r1 := rand.New(s1)

		frog_file := frogs[r1.Intn(len(frogs))]
		f, err := os.Open("frogs/" + frog_file)
		if err != nil {
			panic(fmt.Sprintf("Error opening file: %v", err))
		}

		s.ChannelFileSend(m.ChannelID, frog_file, f)
	}

	// Upper case only
	yelling := regexp.MustCompile(`^\P{L}*\p{Lu}\P{Ll}*$`)

	if yelling.MatchString(m.Content) && m.Content != "LOL" && m.Content != "WTF" {
		// s.ChannelMessageSend(m.ChannelID, shh(m.Content, m.Author.ID))
		// Markov in the future, for now nothing
	}

	// If the message is "ping" reply with "Ping: " and a random digit between 1 and 10
	// if m.Content == "ping" {
	// 	s1 := rand.NewSource(time.Now().UnixNano())
	// 	r1 := rand.New(s1)
	// 	s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Ping: %d", r1.Intn(10)+1))
	// }
	// If the message is "pong" reply with "Ping!"
	// if m.Content == "pong" {
	// 	s.ChannelMessageSend(m.ChannelID, "Ping!")
	// }

	// More thoughts:
	// https://huggingface.co/facebook/blenderbot-400M-distill?text=Hey+my+name+is+Julien%21+How+are+you%3F
	// https://github.com/bwmarrin/discordgo/blob/master/examples/slash_commands/main.go
	// https://github.com/montanaflynn/meme-generator/blob/master/main.go
}

func shh(message string, author string) string {
	c := cases.Title(language.English)

	return fmt.Sprintf("Hey <@%s>, there's no need to yell. \"%s\" works just as well", author, c.String(strings.ToLower(message)))
}
