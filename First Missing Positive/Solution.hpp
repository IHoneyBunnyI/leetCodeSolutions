#include <iostream>
#include <vector>
#include <set>
#include <utility>
#include <algorithm>
#include <map>

using namespace std;

class Solution {
public:
	int firstMissingPositive(vector<int>& nums) {
		if (nums.empty())
			return 1;
		std::map<int, int> m_nums;
		for (int i = 0; i < nums.size(); i++)
			m_nums[nums[i]] = 1;
		for (int i = 1; i <= nums.size(); i++)
		{
			if (m_nums[i] != 1)
				return i;
		}
		return nums.size() + 1;
	}
};
