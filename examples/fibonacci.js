function fibonacci(n) {
  return aux(n, 0, 1);
}

function aux(n, x, y) {
  if (n == 0) {
    return x;
  }
  return aux(n - 1, y, x + y);
}

var i = 1;
var t = 40;
wg.add(t);

console.log('Calculating fibonaccis 40 fibonacci numbers');
while (i <= t) {
  go function(x) {
    console.log('fib(', x, ') =', fibonacci(x));
    wg.done()
  }(i)
  i = i + 1;
}

wg.wait();
