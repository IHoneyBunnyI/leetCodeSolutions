#include <iostream>
#include "Solution.hpp"

int main()
{
	std::vector<int> nums1 {1, 2, 3};
	std::vector<int> nums2 {5, 6, 7};
	Solution s;
	std::cout << s.findMedianSortedArrays(nums1, nums2) << std::endl;
}
