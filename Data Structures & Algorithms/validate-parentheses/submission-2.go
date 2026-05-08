type stack []rune

func (s stack) Push(v rune) stack {
    return append(s, v)
}

func (s stack) Pop() (stack, rune) {
    l := len(s)
    return s[:l - 1], s[l - 1]
}

func (s stack) Empty() bool {
    return len(s) == 0
}

func isValid(s string) bool {
    closeToOpen := map[rune]rune{')': '(', ']': '[', '}': '{'}
    st, top := make(stack, 0), rune(0)
    for _, c := range s {
        if open, exists := closeToOpen[c]; exists {
            if !st.Empty() {
                st, top = st.Pop()
                if top != open {
                    return false
                }
            } else {
                return false
            }
        } else {
            st = st.Push(c)
        }
    }
    return st.Empty()
}
