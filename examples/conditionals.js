function foo(a, b) {
  return a + b
}

function bar(a, b) {
  return foo(a, b) + 1
}

if (bar(1, 2) > foo(1, 2)) {
  console.log('bar foo')
}

if (foo(1, 2) > bar(1, 2)) {
  console.log('foo bar')
} else {
  console.log('else bar foo')
}
