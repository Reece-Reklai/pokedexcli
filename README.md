# PokedexCLI

A command-line Pokémon adventure game built in Go. Explore the Pokémon world, catch Pokémon, and build your personal Pokédex through an interactive REPL interface.

![Go](https://img.shields.io/badge/Go-1.24.3+-00ADD8?style=flat&logo=go))

## Quick Start

```bash
# Clone the repository
git clone https://github.com/Reece-Reklai/pokedexcli.git
cd pokedexcli

# Run the game
go run .
```

## How to Play

Once launched, you'll see the `Pokedex > ` prompt. Type commands to explore and interact:

### Navigation
- `help` - Display all available commands
- `map` - Show current locations
- `nmap` - Navigate to next locations
- `nmapb` - Navigate to previous locations

### Pokémon Actions
- `explore <location>` - Find wild Pokémon in an area
- `catch <pokemon>` - Attempt to catch a Pokémon
- `inspect <pokemon>` - View detailed stats of caught Pokémon
- `pokedex` - View all Pokémon you've caught

### Other
- `exit` - Leave the Pokédex

## Technical Architecture

### Core Components
- **REPL Engine**: Interactive command-line interface with input validation
- **Caching System**: Dual-layer caching for locations and Pokémon encounters with TTL support
- **API Integration**: Seamless integration with PokeAPI.co for real-time Pokémon data
- **State Management**: Player progress tracking with persistent Pokédex storage

### Project Structure
```
├── main.go              # Entry point and command router
├── commands.go          # Command implementations
├── internal/
│   ├── player/          # Player state and Pokémon management
│   ├── location/        # Location data and navigation
│   ├── explore/         # Encounter system
│   └── pokecache/       # Caching layer implementation
└── test/               # Testing utilities
```

## Game Example

```
$ go run .
Pokedex > map
canalave-city
eterna-city
pastoria-city
...
Pokedex > explore canalave-city
Exploring canalave-city...
pikachu
gyarados
magikarp
...
Pokedex > catch pikachu
Throwing a Pokeball at pikachu...
pikachu was caught!
Pokedex > inspect pikachu
Name: pikachu
Height: 4
Weight: 60
Stats:
- hp: 35
- attack: 55
- defense: 40
- special-attack: 50
- special-defense: 50
- speed: 90
Types:
- electric
Pokedex > pokedex
- pikachu
Pokedex > exit
Closing the Pokedex... Goodbye!
```
