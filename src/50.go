function solveEquation(a, b, c) {
    const determinant = b * b - 4 * a * c;
    
    if (determinant >= 0) {
        let x1 = (-b + Math.sqrt(determinant)) / (2 * a);
        let x2 = (-b - Math.sqrt(determinant)) / (2 * a);
        
        return `x1 = ${x1}, x2 = ${x2}`;
    } else {
        return "This equation has no real roots.";
    }
}

solveEquation(1, 2, 3)
