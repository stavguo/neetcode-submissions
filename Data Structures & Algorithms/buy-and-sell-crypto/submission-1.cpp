class Solution {
public:
    int maxProfit(vector<int>& prices) {
        int maxP = 0;
        int left = 0;
        for (int right = 1; right < prices.size(); right++) {
            if (prices[right] < prices[left]) {
                left = right;
                continue;
            }
            maxP = max(maxP, prices[right] - prices[left]);
        }
        return maxP;
    }
};
