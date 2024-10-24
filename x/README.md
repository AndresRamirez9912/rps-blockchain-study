# `/x/{modules}`

This directory contains the code for your chain custom modules.

## State objects

### Game

- gameNumber (uint)
- playerA (address)
- playerB (address)
- status (string)
- rounds (uint)
- playerAMoves ([]string)
- playerBMoves ([]string)
- score ([2]uint) -> score[0] = scoreA / score[1] = scoreB
- expirationHeight (uint64)

### Params (empty)

- ttl (uint64): game's life in blocks

## Message service

### Mesage Create game

- creator (address - signer)
- oponent (address)
- rounds (uint)

### Message Make move

- player (address - signer)
- gameIndex (uint)
- move (string - Rock, Paper, Scissors): SHA256 of the move and salt (random string)

### MsgRevealMove

- player (address -signer)
- gameIndex (uint)
- revealedMove (string - Rock, Paper, Scissors)
- salt

hashedMove == sha256(revealedMoved + salt)

## Query Services

### GetGame

- gameNumber (uint)

### GetParams

## Events

### EventCreateGame

- gameNumber (uint)
- playerA (address)
- playerB (address)

### EventEndGame

- gameNumber (uint)
- status (string)

### EventMakeMove

- gameNumber (uint)
- player (address)
- move (string)

### EventRevealed

- gameNumber (uint)
- player (address)
- reveal_move (string)
