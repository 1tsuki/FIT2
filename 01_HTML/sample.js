document.addEventListener("DOMContentLoaded", function(){
  [].forEach.call(document.getElementsByClassName('btn'), function(v,i,a) {
    v.onclick = function(){ alert("You clicked a button!")};
  })
});
