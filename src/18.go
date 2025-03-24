function sum(a, b) {
  return a + b;
}

function difference(a, b) {
  return a - b;
}

function multiply(a, b) {
  return a * b;
}

function divide(a, b) {
  if (b !== 0) {
    return a / b;
  }
  throw "Division by zero is not allowed.";
}

function squareRoot(x) {
  return Math.sqrt(x);
}
