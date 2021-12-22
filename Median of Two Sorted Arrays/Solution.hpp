#include <iostream>
#include <vector>
#include <set>
#include <utility>
using namespace std;

class Solution {
public:
	double findMedianSortedArrays(vector<int>& nums1, vector<int>& nums2) {
		std::vector<int> merge;
		int i = 0;
		int j = 0;
		while (i < nums1.size() && j < nums2.size())
		{
			if (nums1[i] < nums2[j])
				merge.push_back(nums1[i++]);
			else
				merge.push_back(nums2[j++]);
		}
		while (i < nums1.size())
			merge.push_back(nums1[i++]);
		while (j < nums2.size())
			merge.push_back(nums2[j++]);
		int index = -1;
		if (merge.size() % 2 == 0)
		{
			index = merge.size() / 2;
			double result = (double)(merge[index] + merge[index - 1]) / 2.0;
			return result;
		}
		else
			return merge[merge.size() / 2];
	}
};
