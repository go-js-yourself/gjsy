var foo = function(x) {
  var i = 10;
  while(i > 0) {
    console.log("Thread", x, "iter", i);
    i = i - 1;
  }
}

function dumbest_sleep() {
  console.log("Sleeping on main thread")
  t = 1000000
  while (t > 0) {
    t = t - 1
  }
  console.log("Done")
}

var t = 5;
while (t > 0) {
  go foo(t);
  t = t - 1;
}

dumbest_sleep()
