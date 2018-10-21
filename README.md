# dym

## did you mean X?

A simple tool for checking one or two character typos, useful for detecting input errors.

## Warning

This library is pretty much untested, under development and incomplete. Use at your own risk.

## Usage

- provide your own frequency table in `assets/dict.csv`, it must be a comma separated file with: 

`word`, `number of occurrences`

- `go build .`

```bash
./dym correct pabl
Did you mean 'pablo' ?
```

