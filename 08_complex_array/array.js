// 以下４パターンは全て同じ配列を作成している

// パターン1
var array = ['hoge', 'fuga'];

// パターン2
var array = new Array('hoge', 'fuga');

// パターン3
var array = Array('hoge', 'fuga');

// パターン4
var array = [];
array[0] = 'hoge';
array[1] = 'fuga';


var array1 = ["佐藤", "鈴木", "田中"];
var array2 = ["渡辺", "伊藤", "山本"];

// 配列を結合する
var array = array1.concat(array2);
// array = ["佐藤", "鈴木", "田中", "渡辺", "伊藤", "山本"];

// 先頭の要素を取り出して削除する
var first = array.shift();
// first = "佐藤"
// array = ["鈴木", "田中", "渡辺", "伊藤", "山本"];


// 以下５パターンは全て同じオブジェクトを作成している

// パターン1
var obj = { sato: 'taro', suzuki: 'jiro'};

// パターン2
var obj = { 'sato': 'taro', 'suzuki': 'jiro' };

// パターン3
var obj = {};
obj.sato = 'taro';
obj.suzuki = 'jiro';

// パターン4
var obj = {};
obj['sato'] = 'taro';
obj['suzuki'] = 'jiro';

// パターン5
var obj = new Object();
obj.sato = 'taro';
obj.suzuki = 'jiro';




// 例えば、複数の商品の値段を連想配列にする
var prices = {"納豆": 108, "豆腐": 81, "醤油": 149}
alert(prices["納豆"] + "円"); // 108円

// 例えば、複数の地点の緯度経度を連想配列で扱う
var places = {
    "丹沢山": {"lat": 35.474546, "lng": 139.1543662},
    "焼岳": {"lat": 36.2267693, "lng": 137.5788741},
};
alert(places["丹沢山"]["lat"] + " / " + places["丹沢山"]["lng"]); // 35.474546 / 139.1543662


