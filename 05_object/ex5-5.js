
function reset() {
    var ptags = document.getElementsByTagName("p");
    for (var i = 0; i < ptags.length; i++) {
        ptags[i].innerHTML = "要素" + (i + 1);
    }
}

function changeCcc() {
    var cccs = document.getElementsByClassName("ccc");
    for (var i = 0; i < cccs.length; i++) {
        cccs[i].innerHTML = getValue();
    }
}

function changeAaaBbb() {
    var aaabbbs = document.getElementsByClassName("aaa bbb");
    for (var i = 0; i < aaabbbs.length; i++) {
        aaabbbs[i].innerHTML = getValue();
    }
}

function getValue() {
    var elem = document.getElementById('val');
    return elem.value;
}