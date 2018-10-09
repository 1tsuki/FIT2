var n = Number(prompt("試行回数"));

var m = 0;
for (var i = 0; i < n; i++) {
    var x = Math.random();
    var y = Math.random();

    if (Math.sqrt(Math.pow(x, 2) + Math.pow(y, 2)) < 1) {
        m += 1;
    }
}


var pi = 4 * m / n;
alert(pi);



if (m % 2 == 0) {
    alert("偶数です");
} else {
    alert("奇数です");
}


