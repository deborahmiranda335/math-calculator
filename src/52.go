function add(a, b) {
  return a + b;
}

function subtract(a, b) {
  return a - b;
}

function multiply(a, b) {
  return a * b;
}

function divide(a, b) {
  if (b !== 0) {
    return a / b;
  } else {
    throw new Error("Error: Division by zero");
  }
}

function power(base, exponent) {
  if (exponent === 0) {
    return 1;
  }

  return base ** exponent;
}
