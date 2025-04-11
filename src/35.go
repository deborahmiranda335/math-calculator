function add(x, y) {
  return Math.floor(x + y);
}

function subtract(x, y) {
  return x - y;
}

function multiply(x, y) {
  return x * y;
}

function divide(x, y) {
  if (y === 0) throw "Error: Division by zero";
  return x / y;
}
