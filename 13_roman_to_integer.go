//13. Roman to Integer
//Runtime: 11 ms, faster than 66.50% of Go online submissions for Roman to Integer.
//Memory Usage: 2.8 MB, less than 81.26% of Go online submissions for Roman to Integer.

func romanToInt(s string) int {
    var num int
    var prev rune
    var m = map[rune]int{
        'I': 1,
        'V': 5,
        'X': 10,
        'L': 50,
        'C': 100,
        'D': 500,
        'M': 1000,
    }
    for _, ch := range s {
        if m[ch] > m[prev] {
            num += m[ch] - 2 * m[prev] 
        } else {
            num += m[ch]   
        }
        prev = ch
    }
    return num
}