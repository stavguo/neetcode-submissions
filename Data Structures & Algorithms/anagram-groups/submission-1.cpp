class Solution {
public:
    vector<vector<string>> groupAnagrams(vector<string>& strs) {
        // initialize a map of character count to strings
        unordered_map<string,vector<string>> mp;
        // loop for every word
        for (const auto& s : strs) {
            // initialize char count array
            vector<int> count(26, 0);
            // loop for every character in string
            for (char c : s) {
                // increment character
                count[c - 'a']++;
            }
            string key = to_string(count[0]);
            // convert char array to comma separated string (key)
            for (int i = 1; i < 26; i++) {
                key += "," + to_string(count[i]);
            }
            // push_back current word (val)
            mp[key].push_back(s);
        }
        // initialize return vector of vector string
        vector<vector<string>> res;
        // loop for every key,val pair
        for (auto pair : mp) {
            // push_back val
            res.push_back(pair.second);
        }
        return res;
    }
};
