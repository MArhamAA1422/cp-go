// URL: https://leetcode.com/problems/valid-anagram/

func sorted(s string) string {
    n := len(s)
    a := make([]int, n)
    for i := 0; i < n; i++ {
        a[i] = int(s[i] - 'a')
    }
    sort.Ints(a)
    ss := make([]rune, n)
    for i := 0; i < n; i++ {
        ss[i] = rune(a[i] + 'a')
    }
    return string(ss)
}

func isAnagram(s string, t string) bool {
    s = sorted(s)
    t = sorted(t)
    fmt.Println(s, t)
    return s == t
}