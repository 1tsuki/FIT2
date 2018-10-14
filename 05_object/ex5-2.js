var words = [];
for (var i = 0; i < 5; i++) {
    var word = prompt("文字列を入力してください");
    words.push(word);
}

words.sort();
while (words.length > 0) {
    alert(words.shift());
}