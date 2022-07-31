var foo = function(x) {
  console.log("Foo", x);
}

function bar(y) {
  console.log("Bar", y);
}

function baz(z) {
  return z * 2;
}

foo(1);
bar(10);

foo(baz(2));
bar(baz(baz(baz(10))));
