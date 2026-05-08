class Solution {
public:
    bool isValid(string s) {
        if (s.length() % 2 != 0) return false;
        string stack = "";
        for (int i = 0; i < s.length(); i++) {
            char c = s[i];
            if (c == '(' || c == '{' || c == '[') {
                stack.push_back(c);
            } else {
                if (stack.empty()) return false;
                char last = stack.back();
                if ((c == ')' && last != '(') || 
                    (c == '}' && last != '{') || 
                    (c == ']' && last != '[')) {
                    return false;
                }
                stack.pop_back();
            }
        }
        return stack.empty();
    }
};