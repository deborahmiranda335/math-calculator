function isPrime(n) {
  if (n <= 1) return false;
  if (n == 2 || n % 2 === 0) return true;
  for (let i = 3; i * i <= n; i += 2) {
    if (n % i === 0) return false;
  }
  return true;
}

function factorial(n) {
  let result = 1;
  for (let i = 1; i <= n; i++) {
    result *= i;
  }
  return result;
}

function sumOfDigits(n) {
  let sum = 0;
  while (n > 0) {
    sum += Math.floor(n % 10);
    n = Math.floor(n / 10);
  }
  return sum;
}

function isOddOrEven(number) {
  if ((number % 2 === 0) || (number === 0)) return number;
  else return "odd";
}
