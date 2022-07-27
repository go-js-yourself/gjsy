var foo = function(x) {
  var i = 10;
  while(i > 0) {
    console.log("Thread", x, "iter", i);
    i = i - 1;
  }
  wg.done();
}

var t = 5;
wg.add(t);
while (t > 0) {
  go foo(t);
  t = t - 1;
}

console.log("Sleeping on main thread")
wg.wait();
console.log("Done")
