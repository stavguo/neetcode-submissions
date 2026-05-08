class Solution {
public:
    int longestConsecutive(vector<int>& nums) {
        set<int> st(nums.begin(), nums.end());
        int longest = 0;
        for (int num : st) {
            int current = num;
            int counter = 1;
            while (st.find(current - 1) != st.end()) {
                counter++;
                current--;
            }
            longest = max(longest, counter);
        }
        return longest;
    }
};
