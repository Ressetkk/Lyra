# Windblume Lyre Player

I really love the new event in Genshin Impact. I really love the new toy that they added. WHat did I tell you that you can just paste your favourite tune and then listen to everything you want to!

This simple application parses the score you provide, and then sends mapped virtual key presses to play the Windblume Lyre - the new limited toy from Genshin Impact.

## How does it work?

The application reads encoded base64 string with a score and then plays mapped notes to the keys by sending key events to the OS. If you leave the app run in the background ant you maximize your game with Windblume Lyre on - it starts to play by itself! If the application does not work then try to run it with the administrative privileges.

## Creating a score

Each score is a set of notes and tempo which after encoding gets mapped to the proper keypresses. [See the keymap in score.go]. Each note is a set of commands which application decodes and then performs specific action. Example note goes as follows: `[mod:]e4/4`, where `mod:` is an optional modifier which can be omitted, `e4` is a note which is then mapped to the key and `/d4` which defines the length of a note (Quarter note).

Score is written as a set of notes separated by space, which is shown in the example below, and a song's tempo:
```
a4/8 c5/8 d:f3a4d5/6 e5/16 d5/4 h3g3/8 e5/16 d5/16 c5/8 h4/8 d:a3c4e4/6 h4/16 h4/4 g3h3/4 a4/8 c5/8 d:f3a4d5/6 e5/16 d5/4 h3g3/8 e5/16 d5/16 c5/8 h4/8 d:c3g3e4h4/6 g5/16 g5/4 e3/8 h3/8 g4a5/8 g5/8 d:d3f3c4e5/6 a4/16 e5/4 h3g3/8 d5/16 c5/16 h4/8 g4/8 d:a3c4e4/8 c4/16 h4/16 e4h4/8 g4/8 h3g3/4 a4/8 c5/8 d:f3a4d5/8 d4/16 e5/16 a4d5/8 f3c5/8 g3h4/8 e5/16 d5/16 h4/8 g4/8 d:a3c4e4a4/8 c4/8 e4/8 d4/8 a3e4c5a5/8
```

If you want to encode score to a lyra-compatible base64 format use built-in command `encode`:
```shell
./lyra encode --tempo 40 --author "frostyardor" --name "Winding River" --notes="a4/8 c5/8 +:f3a4d5/6 e5/16 d5/4 h3g3/8 e5/16 d5/16 c5/8 h4/8 +:a3c4e4/6 h4/16 h4/4 g3h3/4 a4/8 c5/8 +:f3a4d5/6 e5/16 d5/4 h3g3/8 e5/16 d5/16 c5/8 h4/8 +:c3g3e4h4/6 g5/16 g5/4 e3/8 h3/8 g4a5/8 g5/8 +:d3f3c4e5/6 a4/16 e5/4 h3g3/8 d5/16 c5/16 h4/8 g4/8 +:a3c4e4/8 c4/16 h4/16 e4h4/8 g4/8 h3g3/4 a4/8 c5/8 +:f3a4d5/8 d4/16 e5/16 a4d5/8 f3c5/8 g3h4/8 e5/16 d5/16 h4/8 g4/8 +:a3c4e4a4/8 c4/8 e4/8 d4/8 a3e4c5a5/8"
```

Additionally you can add `name` and `author` fields to the score. They will be encoded in the score too. Once you encode the string you can share it to others. They will be able to play your tune with no issues.

## Example
Lyra command:
```
Lyra is a simple application that plays custom scors on the Windblume Lyre - toy from limited Genshin Impact event Windblume Festival.
The application sends mapped key events that simulate pressing lyre notes.

Usage:
  lyra [command]

Available Commands:
  encode      Encode raw score to base64 format
  help        Help about any command
  play        Play the score.

Flags:
      --debug   Print debug information.
  -h, --help    help for lyra

Use "lyra [command] --help" for more information about a command.

```
```
./lyra play H4sIAAAAAAAA/+yUsW60MBCEX+XXtjcFthfwz2OkSRFdQQRJKMARRyJFp3v3iAo7srXmRLq06NMynhnNlaZ27Kmhx2Hqhun138Pw2c8Eaj+WNzdTQy+zuyxf7dy59fPSj++OGi5Ak1v6CzVPV3ru24Uau30y5Rk0um49TDdECFVFiWojuIYpoeqNO3mg8khlo7fYI2pJT1mA42cy/hQg4q8ST/ftk835D1PAaNmcxC3eQ7BFWYii/zL/5cwZbGE0fDTlj47nxSLhSWJRdEYrGFpJUEKK//YSazcK+GkkKx9vIouZ7u9GPPgAuTf5wEQxqvw9iHudMRm+oLCEu1RzjslH7Iq9Z1dMvKwZgxBIhjwJXCNDuE0ZffhEHVlD+KGdEtfk7dAiIU7LqkhDVT9W6Hz7DgAA///HnJvV+QgAAA==
```

## DISCLAIMER
Use this software at your own risk! Theoretically any 3rd party software that interacts with the game is against MiHoYo Terms of Service. I don't take any responsibility if your account gets banned after using this app!
