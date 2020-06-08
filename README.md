# flags

Flags is partially an alternative, partially an extension to the golang's standard `flags` package.

**Flags** aims to allow a more unix-style development of application's arguments, providing the following features:

- Double dash (`--option`) for long syntax options
- Single dash (`-o`) for short syntax options
- Possibility to accumulate short syntax options (like `-opx`)
- An automatic and opt-in `--help` (or `-h`) familiar print (column-based, see examples)