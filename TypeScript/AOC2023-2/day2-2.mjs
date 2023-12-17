import * as fs from 'fs' ;

let p2 = 0;
let input = fs.readFileSync('./AOC2023-2/input.txt', 'utf8').trim().split('\n').map(l => l.split(": "));
for (let i = 0; i < input.length; i++) {
    let games = input[i][1];
    let neededRed = 0;
    let neededGreen = 0;
    let neededBlue = 0;
    games.trim().split(";").forEach((game) => {
            game.trim().split(", ").forEach((box) =>{
                    if (box.trim().split(" ")[1] === "red") {
                        if(neededRed < parseInt(box.trim().split(" ")[0])) {
                            neededRed = parseInt(box.trim().split(" ")[0]);
                        }
                    }
                    if (box.trim().split(" ")[1] === "green") {
                        if(neededGreen < parseInt(box.trim().split(" ")[0])) {
                            neededGreen = parseInt(box.trim().split(" ")[0]);
                        }
                    }
                    if (box.trim().split(" ")[1] === "blue") {
                        if(neededBlue < parseInt(box.trim().split(" ")[0])) {
                            neededBlue = parseInt(box.trim().split(" ")[0]);
                        }
                    }
            });
        });
    p2 += neededRed * neededGreen * neededBlue;
}
console.log(p2);
