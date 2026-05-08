class Solution {
public:
    int maxProfit(vector<int>& prices) {
        int maxP = 0, left = 0;
        for (int right = 1; right < prices.size(); right++) {
            int profit = prices[right] - prices[left];
            if (profit > 0) {
                maxP = max(maxP, profit);
            } else {
                left = right;
            }
        }
        return maxP;
    }
};
