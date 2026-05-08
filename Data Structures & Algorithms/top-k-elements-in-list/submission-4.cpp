class Solution {
public:
    vector<int> topKFrequent(vector<int>& nums, int k) {
        unordered_map<int, int> mp;
        vector<vector<int>> freq(nums.size() + 1);
        for (int num : nums) {
            mp[num]++;
        }
        for (const auto& pair : mp) {
            freq[pair.second].push_back(pair.first);
        }
        vector<int> res;
        for (int i = nums.size(); i > 0 && res.size() < k; i--) {
            for (int num : freq[i]) {
                res.push_back(num);
                if (res.size() == k) {
                    return res;
                }
            }
        }
        return res;
    }
};
