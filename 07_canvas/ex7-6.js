var canvas = document.getElementById("canvas")
var ctx = canvas.getContext("2d");
//
// ctx.fillStyle = "lightgray";
// for(var y = 0; y < 200; y = y + 20) {
//     for (var x = 0; x < 200; x = x + 20) {
//         if ((x + y) % 40 == 0) {
//             ctx.fillRect(x, y, 20, 20);
//         }
//     }
// }

ctx.fillStyle = "black";
for(var y = 0; y < 200; y = y + 20) {
    for (var x = 0; x < 200; x = x + 20) {
        if ((x + y) % 80 == 0) {
            ctx.fillRect(x, y, 20, 20);
        }
    }
}


