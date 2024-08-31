package responses

import (
	"math/rand"
	"time"
)

// FunFrogFact generates a response to a message that contains 'fun', 'frog', or 'fact'
func FunFrogFact() string {

	// Create a new random number generator
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	facts := []string{"There is evidence that frogs have roamed the Earth for more than 200 million years. We're like little hoppy dinosaurs that lived through the unprovoked meteor attack on Chicxulub!",
		"The world's largest frog is the goliath frog of West Africa—it can grow to 15 inches and weigh up to 7 pounds (that's 32cm or 3.3kg for people who don't fear science). It comes out at night and dines on fish, crabs, baby turtles, young snakes, and other vertebrates along the river’s edge. Its average life span in the wild is up to 15 years.",
		"While the life spans of frogs in the wild are unknown, frogs in captivity have been known to live more than 20 years, or roughly 40 seasons of The Kardashians",
		"There are over 6000 species of frogs worldwide. Scientists continue to search for new ones, just as quickly as climate change eradicates us (we've lost over 200 species of frogs since the 1970s, ***go humans!!!***)",
		"Toads are frogs. The word \"toad\" is usually used for frogs that have warty and dry skin, as well as shorter hind legs. Except in the Frog & Toad books, and then clearly Toad is just an asshole (frog).",
		"Frogs have excellent night vision ([just like stoners!](https://pubmed.ncbi.nlm.nih.gov/15182912/)) and are very sensitive to movement. The bulging eyes of most frogs allow us to see in front, to the sides, and partially behind us. When a frog swallows food, we pull our eyes down into the roof of our mouth, to help push the food down our throat.",
		"Frogs were the first land animals with vocal cords. Male frogs have vocal sacs—pouches of skin that fill with air. These balloons resonate sounds like a megaphone, and some frog sounds can be heard from a mile away.",
		"Launched by our long legs, many frogs can leap more than 20 times our body length.",
		"The Costa Rican flying tree frog soars from branch to branch with the help of its feet. Webbing between the frog's fingers and toes extends out, helping the frog glide. Yes, that makes them the flying squirrels of the frog world.",
		"To blend into the environment, the Budgett's frog is muddy brown in color, while the Vietnamese mossy frog has spotty skin and bumps to make them look like little clumps of moss or lichen.",
		"Many poisonous frogs, such as the golden poison frog and dyeing poison frog, are boldly colored to warn predators of our dangerous toxic skins. Some colorful frogs, such as the Fort Randolph robber frog, have developed the same coloring as a coexisting poisonous species. Although their skins are not toxic, these mimics may gain protection from predators by looking dangerous. Some frogs wear leather jackets and ride motorcycles, but that's just because they're cool.",
		"Like all amphibians, frogs are cold-blooded, meaning our body temperatures change with the temperature of our surroundings, not that we have cold cold hearts. When temperatures drop, some frogs dig burrows underground or in the mud at the bottom of ponds. They hibernate in these burrows until spring, completely still and scarcely breathing.",
		"The wood frog can live north of the Arctic Circle, surviving for weeks with 65 percent of its body frozen. This frog uses glucose in its blood as a kind of antifreeze that concentrates in its vital organs, protecting them from damage while the rest of the body freezes solid. A frogsicle if you will.",
		"The Australian water-holding frog is a desert dweller that can wait up to seven years for rain. It burrows underground and surrounds itself in a transparent cocoon made of its own shed skin. Or: ***It rubs the cocoon of the skin until it gets the rain again***",
		"Frogs are freshwater creatures, although some frogs such as the Florida leopard frog are able to live in brackish or nearly completely salt waters. And yes, we definitely mock those Florida frogs.",
		"Almost all frogs fertilize the eggs outside of the female's body. The male holds the female around the waist in a mating hug called amplexus. He fertilizes the eggs as the female lays them. Amplexus can last hours or days. One pair of Andean toads stayed in amplexus for four months (he definitely took frog viagra).",
		"The marsupial frog keeps her eggs in a pouch like a kangaroo. When the eggs hatch into tadpoles, she opens the pouch with her toes and spills them into the water.",
		"Pipa pipa, the Suriname toad of South America, carries her young embedded in the skin of her back. After mating, the eggs sink gradually into the female's back, and a skin pad forms over the eggs. The developing juvenile frogs are visible inside our pockets for several days before hatching. They emerge over a period of days, thrusting their head and forelegs out first, then struggling free. It is 110% as gruesome and metal as that sounds",
		"The gastric brooding frog of Australia swallows her fertilized eggs. The tadpoles remain in her stomach for up to eight weeks, finally hopping out of her mouth as little frogs. During the brooding period, gastric secretions cease—otherwise she would digest her own offspring.",
		"Among Darwin frogs, it is the male who swallows and stores the developing tadpoles in his vocal sac until juvenile frogs emerge. Google doesn't mention if he ceases to produce gastric secretions like the female gastric brooding frog of Australia, but we can only hope, otherwise that's just a fancy way to say cannibalism.",
		"Frogs live around the world, on every continent, except Antarctica. (We've tried, too fucking cold!)",
		"A group of frogs is called an army. Maybe it’s because we wear army green camouflage! (You humans have never asked)",
		"Each frog species has its own special call from  \"brrr-ummmm\" to \"jug-o-rum\". Males croak during mating season to attract a female. The louder he croaks, the more likely he is to attract a mate.",
		"Frogs have teeth! The small teeth on the roof of our mouths are not typically used to bite or chew; they keep our dinner from escaping before we get a chance to swallow it. However, if a frog feels threatened, or a human hand-feeds a pet frog, certain species have been known to bite.",
		"Frogs don’t drink water. We absorb water through our skin.",
		"Not all frogs can jump (just like white men! I kid! Mostly!). While most long-legged species can jump a distance greater than 20 times our body length, those with shorter back legs can hop, crawl, or walk.",
		"The South African sharp-nosed frog holds the world’s record for the longest jump. It jumped 44 times its body length. This 3-inch species leaped more than 130 inches. To match that, a five-foot-tall human would have to jump 220 feet in one leap.",
		"The world’s tiniest frog is the Paedophryne amanuensis. It’s about the size of a common housefly. It lives in leaf litter in the rain forests of Papua New Guinea.",
		"The golden poison frog, native to Central and South American rainforests, has the distinction of being the most poisonous animal in the world, despite being about the length of a paper clip. Its skin secretes enough nerve toxin to kill 10 humans, and one gram of the toxin it produces could kill 100000 people. ***Mmmmm forbidden paperclip***",
		"A frog completely sheds its skin about once a week. After it pulls off the old, dead skin, the frog usually eats it.",
		"Jakai, the lungless frog of [TNBBBR](https://en.wikipedia.org/wiki/Bukit_Baka_Bukit_Raya_National_Park), breathes entirely through its skin.",
	}

	return facts[r1.Intn(len(facts))]
}
