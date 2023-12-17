import * as fs from 'fs';

console.log(
    fs.readFileSync('./AOC2023-1/ex.txt', 'utf8')
        .trim()
        .split('\n')
        //Replace one by 1e
        .map((l,i) => l.replace(/one/g, "o1e").replace(/two/g, "t2o").replace(/three/g, "t3e").replace(/four/g, "f4").replace(/five/g, "f5e").replace(/six/g, "s6").replace(/seven/g, "s7n").replace(/eight/g, "e8t").replace(/nine/g, "n9e"))
        .map((l,i) => l.split("").reduce((acc,c) => acc + (c >= "0" && c <= "9" ? c : ""), ""))
        .map((n) => parseInt(n[0] + n[n.length - 1]))
        .reduce((acc, c) => acc + c, 0))

console.log(
    fs.readFileSync('./AOC2023-1/input.txt', 'utf8')
        .trim()
        .split('\n')
        //Replace one by 1e
        .map((l,i) => l.replace(/one/g, "o1e").replace(/two/g, "t2o").replace(/three/g, "t3e").replace(/four/g, "f4").replace(/five/g, "f5e").replace(/six/g, "s6").replace(/seven/g, "s7n").replace(/eight/g, "e8t").replace(/nine/g, "n9e"))
        .map((l,i) => l.split("").reduce((acc,c) => acc + (c >= "0" && c <= "9" ? c : ""), ""))
        .map((n) => parseInt(n[0] + n[n.length - 1]))
        .reduce((acc, c) => acc + c, 0))