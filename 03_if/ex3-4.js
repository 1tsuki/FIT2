function checkClass() {
    var weekday = prompt("曜日は？");
    var period = prompt("時限は？");

    switch (weekday) {
        case "火曜":
            switch (period) {
                case "1":
                case "2":
                    alert("情報基礎2");
                    break;
                default:
                    alert("何も履修していません");
            }
            break;
        case "月曜":
        case "水曜":
        case "木曜":
        case "金曜":
        default:
            alert("何も履修していません");
    }
}