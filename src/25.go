function calculate(a, b) {
  if (typeof a === 'number' && typeof b === 'number') {
    return Math.pow(a, b);
  } else if (Array.isArray(a)) {
    return a.reduce((acc, val) => acc + evaluate(val), 0);
  }
}

function evaluate(expression) {
  let result = expression;
  
  for (let i = 0; i < result.length; i++) {
    if (result[i] === ' ') continue;
    
    const currentOperator = '+-*/';
    if (result[i + 1] === null || currentOperator.includes(result[i + 1])) {
      break;
    } else {
      let nextNum = parseFloat(result.slice(i, i + 2));
      
      result = currentOperator.replace(/\+/g, '').replace(currentOperator.replace(/\-/g, '-').replace(/ /g, ''));
      continue;
    }
  }
  
  return Number(result);
}
