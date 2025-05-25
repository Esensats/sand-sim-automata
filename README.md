# Sand Simulation Automata

A real-time sand simulation using cellular automata that runs in your terminal console. Watch as sand particles fall and pile up naturally, following physics-based rules!

![Sand Simulation Demo](https://img.shields.io/badge/demo-terminal-brightgreen)

## 🌟 Features

- **Real-time simulation**: Smooth 30 FPS sand particle animation
- **Physics-based movement**: Sand falls down and spreads naturally
- **Terminal-based**: Runs entirely in your console with ASCII graphics
- **Randomized spawning**: Sand particles appear randomly at the top
- **Cellular automata**: Each sand grain follows simple local rules

## 🛠️ Tech Stack

- **Language**: Go 1.23.1
- **Platform**: Linux x86_64
- **Build System**: Makefile
- **Dependencies**: None (uses only Go standard library)

## 🚀 Quick Start

### Prerequisites

- Go 1.23.1 or higher
- Linux terminal (or compatible terminal emulator)

### Installation & Running

1. **Clone the repository**:
   ```bash
   git clone https://github.com/esensats/sand-sim-automata.git
   cd sand-sim-automata
   ```

2. **Build and run**:
   ```bash
   make run
   ```

   Or build separately:
   ```bash
   make build
   ./bin/sand-sim
   ```

3. **Exit**: Press `Ctrl+C` to stop the simulation

## 🎮 How It Works

The simulation creates a 40×20 grid where:

- **Empty spaces** are represented by spaces (` `)
- **Sand particles** are represented by blocks (`▒`)
- Sand spawns randomly at the top row (10% chance per frame)
- Each sand particle follows these rules:
  1. Try to fall straight down
  2. If blocked, try to fall diagonally (left or right, randomly chosen)
  3. If all paths are blocked, stay in place

## 📁 Project Structure

```
.
├── main.go           # Main application entry point
├── go.mod           # Go module definition
├── Makefile         # Build automation
├── spec.md          # Project specification
├── bin/             # Compiled binaries
│   └── sand-sim
├── grid/            # Grid data structure and cell types
│   └── grid.go
├── renderer/        # Terminal rendering logic
│   ├── renderer.go
│   └── renderer_test.go
└── simulation/      # Simulation logic and physics
    └── simulation.go
```

## 🔧 Build Commands

| Command      | Description                              |
| ------------ | ---------------------------------------- |
| `make build` | Compile the binary to `./bin/sand-sim`   |
| `make run`   | Build and immediately run the simulation |
| `make clean` | Remove compiled binaries                 |

## 🎯 Configuration

You can modify these constants in `main.go` to customize the simulation:

```go
const (
    WIDTH  = 40    // Grid width
    HEIGHT = 20    // Grid height  
    FPS    = 30    // Frames per second
)
```

The sand spawn rate can be adjusted by changing the probability in the main loop:
```go
if randSource.Float32() < 0.1 {  // 10% chance per frame
    // Spawn sand
}
```

## 🧪 Testing

Run the tests with:
```bash
go test ./...
```

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 📝 License

This project is open source. Feel free to use, modify, and distribute as needed.

## 🎨 Future Ideas

- [ ] Add different particle types (water, stone, etc.)
- [ ] Interactive mode (place sand with mouse/keyboard)
- [ ] Color support for different materials
- [ ] Save/load simulation states
- [ ] Adjustable simulation speed
- [ ] Larger grid sizes with scrolling
- [ ] Wind effects and other environmental factors

---

**Enjoy watching the sand fall!** 🏖️
