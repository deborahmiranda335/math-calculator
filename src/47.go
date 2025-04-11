class Solution {
  public int calculate(int a, int b, char operator) {
    switch(operator){
      case '+':
        return a + b;
      case '-':
        return a - b;
      case '*':
        return a * b;
      case '/':
        if(b == 0)
          throw new IllegalArgumentException("Division by zero is not allowed");
        else
          return a / b;
    }
  }
}
