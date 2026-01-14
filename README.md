# Conway’s Game of Life — Rules

## Universe

- The universe is an infinite two-dimensional grid of square cells.
- Each cell is in one of two states:
  - **Live**
  - **Dead**
- Each cell has **8 neighbors**:
  - Horizontally adjacent
  - Vertically adjacent
  - Diagonally adjacent

## Update Cycle

Time advances in discrete steps called **ticks**.  
All cells are updated **simultaneously** each tick.

Each new generation is computed only from the previous one.

## Rules

### For Live Cells

- **Underpopulation**
  - If a live cell has **fewer than 2 live neighbors**, it becomes **dead**.

- **Survival**
  - If a live cell has **2 or 3 live neighbors**, it **stays alive**.

- **Overpopulation**
  - If a live cell has **more than 3 live neighbors**, it becomes **dead**.

### For Dead Cells

- **Reproduction**
  - If a dead cell has **exactly 3 live neighbors**, it becomes **alive**.

## Summary Table

| Current State | Live Neighbors | Next State |
| ------------- | -------------- | ---------- |
| Live          | < 2            | Dead       |
| Live          | 2 or 3         | Live       |
| Live          | > 3            | Dead       |
| Dead          | = 3            | Live       |
| Dead          | Any other      | Dead       |
