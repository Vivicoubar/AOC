import * as fs from 'fs' ;

let red = 12;
let green = 13;
let blue = 14;
let p1 = 0;
let gameCounter = 1;

let input = fs.readFileSync('./AOC2023-2/input.txt', 'utf8').trim().split('\n')
    .map(l => l.split(": "));

for (let i = 0; i < input.length; i++) {
    let games = input[i][1];
    let possible = true;
    games.trim().split(";").forEach((game) => {
            game.trim().split(", ").forEach((box) =>{
                    if (box.trim().split(" ")[1] === "red") {
                        if (parseInt(box.trim().split(" ")[0]) > red) {
                            possible = false;
                        }
                    }
                    if (box.trim().split(" ")[1] === "green") {
                        if (parseInt(box.trim().split(" ")[0]) > green) {
                            possible = false;
                        }
                    }
                    if (box.trim().split(" ")[1] === "blue") {
                        if (parseInt(box.trim().split(" ")[0]) > blue) {
                            possible = false;
                        }
                    }
            })
        }
    )
    if (possible) {
        p1+=gameCounter;
    }
    gameCounter++
}
console.log(p1)
