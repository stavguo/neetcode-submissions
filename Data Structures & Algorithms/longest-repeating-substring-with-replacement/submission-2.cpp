class Solution {
public:
    int characterReplacement(string s, int k) {
        int l = 0;
        int res = 0;
        int max_freq = 0;
        unordered_map<char, int> count;
        for (int r = 0; r < s.size(); r++) {
            count[s[r]]++;
            max_freq = max(max_freq, count[s[r]]);
            while (r - l + 1 - max_freq > k) {
                count[s[l]]--;
                l++;
            }
            res = max(res, r - l + 1);
        }
        return res;
    }
};
