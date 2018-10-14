function checkClass() {
    var weekday = prompt("曜日は？ (月曜,火曜...)");
    var period = prompt("時限は？ (1,2...)");

    switch (weekday) {
        case "火曜":
            switch (period) {
                case "1":
                case "2":
                    alert("情報基礎2");
                    break;
                default:
                    alert("原則リモート勤務");
            }
            break;
        case "月曜":
        case "水曜":
        case "木曜":
        case "金曜":
        default:
            alert("出社日");
    }
}

