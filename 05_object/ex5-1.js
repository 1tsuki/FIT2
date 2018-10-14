var words = [];
for (var i = 0; i < 6; i++) {
    var word = prompt("名前を入力してください");
    words.push(word);
}

var counter = 0;
var message = "";
while (words.length > 0) {
    counter++;
    message += words.shift();
    if (counter >= 3) {
        alert(message);
        counter = 0;
        message = "";
    }
}