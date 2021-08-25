# NBT-CLI
Simple command-line tool for working with Minecraft's NBT format

## Usage
You can obtain help on usage with `nbt --help`

### View data stored in file
You can omit `-l` flag if default value (50) satisfies you. This flag allows to limit count of displayed entries in arrays.

```bash
nbt -f some.dat tree -l 100
```
