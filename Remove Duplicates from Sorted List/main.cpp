#include <iostream>
#include "Solution.hpp"

ListNode *generate_list(std::vector<int> &nums)
{
	if (nums.size() == 0)
		return 0;
	ListNode *first = new ListNode(nums[nums.size() - 1]);
	if (nums.size() > 1)
	{
		for (int i = nums.size() - 2; i >= 0; i--)
		{
			ListNode *node = new ListNode(nums[i], first);
			first = node;
		}
	}
	return first;
}

void print_list(ListNode *list)
{
	ListNode *tmp = list;
	while (tmp)
	{
		std::cout << tmp->val << " ";
		tmp = tmp->next;
	}
	std::cout << std::endl;
}

int main()
{
	std::vector<int> nums {1, 1, 2};
	ListNode *list = generate_list(nums);
	print_list(list);
	Solution s;
}
