class Solution {
public:
    int lengthOfLongestSubstring(string s) {
        unordered_set<char> charSet;
        int l = 0;
        int res = 0;
        for (int r = 0; r < s.size(); r++) {
            // Delete existing
            while (charSet.find(s[r]) != charSet.end()) {
                charSet.erase(s[l]);
                l++;
            }
            // Add new one
            charSet.insert(s[r]);
            // Update res
            res = max(res, r - l + 1);
        }
        return res;
    }
};
