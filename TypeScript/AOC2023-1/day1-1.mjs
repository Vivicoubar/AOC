import * as fs from 'fs';

console.log(
    fs.readFileSync('./AOC2023-1/ex.txt', 'utf8')
        .trim()
        .split('\n')
        .map((l,i) =>
            l.split("")
                .reduce((acc,c) => acc + (c >= "0" && c <= "9" ? c : ""), ""))
        .map((n) => parseInt(n[0] + n[n.length - 1]))
        .reduce((acc, c) => acc + c, 0)
)

console.log(
    fs.readFileSync('./AOC2023-1/input.txt', 'utf8')
        .trim()
        .split('\n')
        .map((l,i) =>
            l.split("")
                .reduce((acc,c) => acc + (c >= "0" && c <= "9" ? c : ""), ""))
        .map((n) => parseInt(n[0] + n[n.length - 1]))
        .reduce((acc, c) => acc + c, 0)
)