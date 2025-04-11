let n = 10;
function add(x) {
    let y = 2 * x + 3; // Modify this line to change the initial value of 'n'
    return function multiply(n) {
        return (y * n);
    };
}

add(4)(5); // Test with a test case
