var className = "table table-striped";

function changeTableStyle() {
    var tbl = document.getElementById("tbl");
    switch (className) {
        case "table table-striped":
            className = "table table-bordered";
            break;
        case "table table-bordered":
            className = "table table-striped table-sm";
            break;
        case "table table-striped table-sm":
            className = "table table-striped";
            break;
    }
    tbl.className = className;
}