#Sugoku

```bash
git clone https://github.com/Virtual-Machine/sugoku.git
cd sugoku
go build
./sugoku #solves puzzle.csv with detail
```

Solves most puzzles by first eliminating possibilities, and then recursively solving for exclusive and hidden singles, pairs, and triples.

Could be improved further by adding some more advanced solving algorithms such as x-wing, swordfish, and walking connected nodes of values to eliminate further possibilities.