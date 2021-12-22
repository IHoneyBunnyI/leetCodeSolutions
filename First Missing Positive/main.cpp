#include <iostream>
#include "Solution.hpp"

int main()
{
	std::vector<int> nums1 {1, 2, 0};
	std::vector<int> nums2 {3,4,-1,1};
	std::vector<int> nums3 {7,8,9,11,12};
	std::vector<int> nums4 {0};
	std::vector<int> nums5 {5, 3 , 2, 1};
	Solution s;

	std::cout << s.firstMissingPositive(nums1) << std::endl;
	std::cout << s.firstMissingPositive(nums2) << std::endl;
	std::cout << s.firstMissingPositive(nums3) << std::endl;
	std::cout << s.firstMissingPositive(nums5) << std::endl;
}
