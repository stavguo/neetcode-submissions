class Solution {
public:
    string minWindow(string s, string t) {
        if (t.empty()) {
            return "";
        }

        unordered_map<char, int> t_count, window;
        for(char &c : t) {
            t_count[c]++;
        }
        pair<int, int> res = {-1, -1};
        int resLen = INT_MAX;
        int have = 0, need = t_count.size();
        int l = 0;
        // sliding window
        for (int r = 0; r < s.length(); r++) {
            window[s[r]]++;
            if (t_count.count(s[r]) && t_count[s[r]] == window[s[r]]) {
                have++;
            }
            while (have == need) {
                // update res
                if ((r - l + 1) < resLen) {
                    resLen = r - l + 1;
                    res = {l, r};
                }
                // force remove l from window
                window[s[l]]--;
                // update have
                if (t_count.count(s[l]) && window[s[l]] < t_count[s[l]]) {
                    have--;
                }
                // lastly increment l
                l++;
            }
        }
        return resLen == INT_MAX ? "": s.substr(res.first, resLen);
    }
};
