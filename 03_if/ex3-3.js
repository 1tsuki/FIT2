function dinnerIf() {
    var appetite = ask();

    if (appetite == "洋食") {
        alert("ジョエル・ロブション");
    } else if (appetite == "和食") {
        alert("数寄屋橋次郎");
    } else if (appetite == "中華") {
        alert("四川飯店");
    } else {
        alert("学食");
    }
}

function dinnerSwitch() {
    var appetite = ask();

    switch (appetite) {
        case "洋食":
            alert("ジョエル・ロブション");
            break;
        case "和食":
            alert("数寄屋橋次郎");
            break;
        case "中華":
            alert("四川飯店");
            break;
        default:
            alert("学食");
    }
}

function ask() {
    return prompt("今日の夕飯の気分を 洋食/和食/中華/その他 で答えてください");
}