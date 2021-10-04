# Windsong Lyre Player

I really love the new event in Genshin Impact. I really love the new toy that they added. What if I told you that you can just paste your favourite tune and then listen to everything you want to!

This simple application parses the score you provide, and then sends mapped virtual key presses to play the Windsong Lyre - the new limited toy from Genshin Impact.

## How does it work?

The application reads MIDI file from local directory or directly from URL and then plays mapped notes to the keys by sending key events to the OS. If you leave the app run in the background ant you maximize your game with Windsong Lyre on - it starts to play by itself! If the application does not work then try to run it with the administrative privileges.

## Example
Lyra command:
```
Lyra is a simple application that plays custom scors on the Windsong Lyre - toy from limited Genshin Impact event Windblume Festival.
The application sends mapped key events that simulate pressing lyre notes.

Usage:
  lyra [command]

Available Commands:
  help        Help about any command
  play        Play the score.

Flags:
      --debug   Print debug information.
  -h, --help    help for lyra

Use "lyra [command] --help" for more information about a command.

```

## DISCLAIMER
Use this software at your own risk! Theoretically any 3rd party software that interacts with the game is against MiHoYo Terms of Service. I don't take any responsibility if your account gets banned after using this app!

## Key Map

|Note|Note (in-game)|Key|
|:---:|:---:|:---:|
|c3|do1|z|
|d3|re1|x|
|e3|mi1|c|
|f3|fa1|v|
|g3|so1|b|
|a3|la1|n|
|h3|ti1|m|
|c4|do2|a|
|d4|re2|s|
|e4|mi2|d|
|f4|fa2|f|
|g4|so2|g|
|a4|la2|h|
|h4|ti2|j|
|c5|do3|q|
|d5|re3|w|
|e5|mi3|e|
|f5|fa3|r|
|g5|so3|t|
|a5|la3|y|
|h5|ti3|u|