# TUI Package

This package implements the Miller Columns TUI for the `scd` tool.

## Structure
- `tui.go`: Main entry point (`RunTUI`).
- `model.go`: Bubble Tea model lifecycle (`Init`).
- `update.go`: Keybinding and state navigation logic (`Update`).
- `view.go`: Rendering logic (`View`, `renderColumn`).
- `types.go`: Internal data structures.
- `styling.go`: UI styling using Lip Gloss.
- `utils.go`: File system interaction utilities.

## Features
- Miller Columns (3-column layout: Parent, Current, Preview).
- Navigates using `h/j/k/l` or arrow keys.
- Filters out files by default; toggle with `.`.
- Shell integration safe: draws to `Stderr` to allow `Stdout` capture.
