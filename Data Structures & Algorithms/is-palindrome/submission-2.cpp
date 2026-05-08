class Solution {
public:
    bool isPalindrome(string s) {
        string spaceless = "";
        for (char c : s) {
            if (isalnum(c)) {
                spaceless += tolower(c);
            }
        }
        int left = 0;
        int right = spaceless.length() - 1;
        while (left < right) {
            if (spaceless[left] != spaceless[right]) {
                return false;
            }
            left++;
            right--;
        }
        return true;
    }
};
