var size = 12;
var color = "red";

function changeStyle() {
    var element = document.getElementById("target");

    // 文字の大きさを大きくする
    size = size + 1;
    element.style.fontSize = size + "px";

    // 文字色を切り替える
    switch(color) {
        case "red":
            color = "green";
            break;
        case "green":
            color = "red";
            break;
    }
    element.style.color = color;
}

