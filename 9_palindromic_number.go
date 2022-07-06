//9. Palindrome Number
//Runtime: 88 ms, faster than 5.30% of Go online submissions for Palindrome Number.
//Memory Usage: 6.3 MB, less than 13.54% of Go online submissions for Palindrome Number.
//not int casting to string

func isPalindrome(x int) bool {
    if x < 0 {
        return false
    }
    var digits []int
    for {
        digits = append(digits, x % 10)
        x = x / 10
        if x == 0 {
            break
        }
    }
    l := len(digits)
    for i := 0; i < l / 2; i++ {
        if digits[i] != digits[l-i-1] {
            return false
        }
    }
    return true
}