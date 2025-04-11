// Define functions to calculate different types of operations
function add(x, y) {
    return x + y;
}

function subtract(x, y) {
    return x - y;
}

function multiply(x, y) {
    return x * y;
}

function divide(x, y) {
    if (y !== 0) {
        return x / y;
    }
    throw new Error("Division by zero is not allowed.");
}

// Main function to handle user input and perform calculations
function main() {
    console.log("Select operation:");
    console.log("1. Add");
    console.log("2. Subtract");
    console.log("3. Multiply");
    console.log("4. Divide");

    const choice = parseInt(prompt("Enter your choice (1/2/3/4):"));

    switch (choice) {
        case 1:
            console.log(`Result: ${add(5, 3)}`);
            break;
        case 2:
            console.log(`Result: ${subtract(10, 4)}`);
            break;
        case 3:
            console.log(`Result: ${multiply(4, 5)}`);
            break;
        case 4:
            try {
                const result = divide(10, 0);
                console.log(`Result: ${result}`);
            } catch (error) {
                console.error(error.message);
            }
            break;
    }
}
