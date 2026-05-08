class Solution {
public:
    int maxArea(vector<int>& heights) {
        int i = 0;
        int j = heights.size() - 1;
        int max_area = (j - i) * min(heights[i], heights[j]);
        while (i != j) {
            int curr_area = (j - i) * min(heights[i], heights[j]);
            // check if best
            if (curr_area > max_area) {
                max_area = curr_area;
            }
            // update
            if (heights[i] < heights[j]) {
                i++;
            } else {
                j--;
            }
        }
        return max_area;
    }
};
