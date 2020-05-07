# Marriage Card Game Score Calculator
My family loves playing the [Marriage Card Game](https://en.wikipedia.org/wiki/Marriage_(card_game)) but we hate having to keep track of the weird scoring system and determining who won at the end of the game

I will only do something manually once. The second time I need a program. Here's the program to handle that and let us focus on purely playing the game and leave the scoring to the dumb machine

# PSA
This is my very first Golang project so the code is probably not very idiomatic. I am looking to improve it as much as possible so if by some chance you find yourself here, do take a look at the code and suggest improvements I can make

# Install
Simply download the relevant binary from the `Releases` tab for your operating system

# Sample usage
Naviate the [examples](./examples) directory to see examples. Let's consider the following example `config.yml` file:

```yaml
players:
  - Goku
  - Vegeta
  - Gohan
  - Trunks
rounds: 1
```
This is the bare minimum a configuration file needs. Running `marriage start -f config.yml` will yield a `game.yml` file that looks like:

```yaml
rounds:
- roundnum: 1
  players:
  - name: Goku
    score: 0
    winner: false
    pachayo: false
  - name: Vegeta
    score: 0
    winner: false
    pachayo: false
  - name: Gohan
    score: 0
    winner: false
    pachayo: false
  - name: Trunks
    score: 0
    winner: false
    pachayo: false
```
This is the game configuration file that will be used to keep score. However you choose to keep score is up to you and you can modify the `score` value freely. The only enforced rule here is:
> The player who ends the game will get 10 points from each of the players who have not completed primary/pure rounds, and 3 points from each of the players who have done so
-- <cite>[Wikipedia](https://en.wikipedia.org/wiki/Marriage_(card_game))</cite>

If a player has not completed the primary round, mark the `pachayo` key as false for that player for that round. Mark `winner: true` for the player that won the round. That will be sufficient to handle the enforced scoring rule. You can choose to compute the `score` field however you want besides that

To score the game, simply run:
```bash
marriage score
```
or
```bash
marriage score -g <other_file>.yml
```
if using some other game yml file besides the default `game.yml` file

You will see an output like this:
<img src="assets/gameResults.jp">

# TODOS
- Figure out where to put all the consts
- Logic to determine only one winner still doesn't work
- More tests
