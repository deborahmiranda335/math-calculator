func main() {
    // Set up our math library
    math := NewMathLibrary()

    // Define our function
    func calculate(numbers []int) int {
        // Calculate the sum of the numbers
        var sum = 0
        for _, number := range numbers {
            sum += number
        }

        return sum
    }

    // Test our function
    t.Run("Calculates the sum of two numbers", func(t *testing.T) {
        result := calculate([]int{1, 2})
        if result != 3 {
            t.Errorf("Expected result to be 3 but was %d", result)
        }
    })
}
