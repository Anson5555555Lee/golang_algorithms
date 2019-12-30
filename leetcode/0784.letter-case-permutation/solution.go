package problem0784

// class Solution {
// 	public:
// 		vector<string> letterCasePermutation(string S) {
// 			vector<string> result;
// 			traceback(result, S, 0);
// 			return result;
// 		}
// 	private:
// 		void traceback(vector<string> &result, string s, int len) {
// 			if (len == s.size()) {
// 				result.push_back(s);
// 				return;
// 			}
// 			traceback(result, s, len + 1);
// 			if (s[len] > '9') {
// 				s[len] ^= (1 << 5);
// 				traceback(result, s, len + 1);
// 			}
// 		}
// };
