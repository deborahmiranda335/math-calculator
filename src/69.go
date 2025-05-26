function calculate(a, b, operation) {
  switch (operation) {
    case 'add':
      return a + b;
    case 'subtract':
      return a - b;
    case 'multiply':
      return a * b;
    case 'divide':
      if (b === 0) {
        throw new Error('Cannot divide by zero');
      }
      return a / b;
    default:
      throw new Error(`Unknown operation: ${operation}`);
  }
}

// Example usage
console.log(calculate(5, 3, 'add')); // Output: 8
console.log(calculate(5, 3, 'subtract')); // Output: 2
console.log(calculate(5, 3, 'multiply')); // Output: 15
console.log(calculate(5, 3, 'divide')); // Output: 1.6666666666666667
