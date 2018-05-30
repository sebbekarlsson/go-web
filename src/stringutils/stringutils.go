package stringutils


func RmFirstChar(s string) string {
    for i := range s {
        if i > 0 {
            return s[i:]
        }
    }

    return s[:0]
}
