function solveEquation(terms) {
  const coefficients = terms.coefficients;
  const exponents = terms.exponents;

  for (let i = 0; i < exponents.length - 1; i++) {
    if (exponents[i] === exponents[i + 1]) {
      // Terms with the same exponent
      return "The equation has an infinite number of solutions.";
    }
  }

  // For non-constant terms, we need to check each coefficient separately.
  const constantCoefficients = coefficients.filter(coefficient => coefficient === 0).map(coefficient => `x^${coefficient}`);
  for (let i = 0; i < exponents.length - 1; i++) {
    if (exponents[i] !== exponents[i + 1]) {
      return `${constantCoefficients[i]}x${i + 2}+x${i + 3}`;
    }
  }

  // For a constant term, we need to check the remaining coefficients.
  if (terms.coefficients[0] === 0) {
    return "The equation has one solution: x = "+`${exponents[exponents.length - 1]}`+`.";
  } else {
    return `x${exponents.length-1}`;
  }
}
