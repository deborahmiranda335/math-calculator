<code>
import { useState } from 'react';

const Calculator = () => {
  const [currentOperand, setCurrentOperand] = useState('');
  const [lastOperation, setLastOperation] = useState('');

  const handleCalculate = (operation) => {
    if (isNaN(parseInt(currentOperand))) {
      return;
    }

    const operands = currentOperand.split(' ');
    const firstOperand = parseInt(operands[0]);
    const secondOperand = parseInt(operands[2]);

    switch (operation) {
      case 'add':
        setLastOperation(firstOperand + secondOperand);
        break;
      case 'subtract':
        setLastOperation(firstOperand - secondOperand);
        break;
      case 'multiply':
        setLastOperation(firstOperand * secondOperand);
        break;
      case 'divide':
        if (secondOperand === 0) {
          setLastOperation('Error: Division by zero');
        } else {
          setLastOperation(firstOperand / secondOperand);
        }
        break;
    }

    setCurrentOperand(lastOperation);
  };

  return (
    <div>
      <h2>数学计算器</h2>
      <input
        type="text"
        value={currentOperand}
        onChange={(e) => setCurrentOperand(e.target.value)}
      />
      {lastOperation && <p>{lastOperation}</p>}
      <button onClick={() => handleCalculate('add')}>+</button>
      <button onClick={() => handleCalculate('subtract')}>-</button>
      <button onClick={() => handleCalculate('multiply')}>x</button>
      <button onClick={() => handleCalculate('divide')}>/</button>
    </div>
  );
};

export default Calculator;
