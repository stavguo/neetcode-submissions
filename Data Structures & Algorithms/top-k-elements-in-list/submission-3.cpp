class Solution {
public:
    vector<int> topKFrequent(vector<int>& nums, int k) {
        unordered_map<int, int> mp;
        for (const auto& num : nums) {
            mp[num]++;
        }
        vector<vector<int>> freq(nums.size() + 1);
        for (auto pair : mp) {
            freq[pair.second].push_back(pair.first);
        }
        vector<int> res;
        for (int i = nums.size(); i >= 0 && res.size() < k; i--) {
            for (int num : freq[i]) {
                if (res.size() < k) {
                    res.push_back(num);
                }
            }
        }
        return res;
    }
};
