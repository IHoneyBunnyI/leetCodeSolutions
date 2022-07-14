#include <iostream>
#include <vector>
#include <set>
#include <utility>
using namespace std;

struct ListNode {
	 int val;
	 ListNode *next;
	 ListNode() : val(0), next(nullptr) {}
	 ListNode(int x) : val(x), next(nullptr) {}
	 ListNode(int x, ListNode *next) : val(x), next(next) {}
};

class Solution {
public:
    ListNode* deleteDuplicates(ListNode* head) {
		if (!head)
			return 0;
		if (!head->next)
			return head;
		ListNode *result;
		if (head->val != head->next->val)
			result = new ListNode(head->val);
		else
			result = new ListNode(500);
		ListNode *tmp_res = result;
		while (head->next && head->next->next)
		{
			ListNode *next = head->next;
			ListNode *next_next = head->next->next;
			if (head->val != next->val && next->val != next_next->val)
			{
				ListNode *node = new ListNode(next->val);
				tmp_res->next = node;
				tmp_res = tmp_res->next;
			}
			head = head->next;
		}
		if (head->val != head->next->val)
		{
				ListNode *node = new ListNode(head->next->val);
				tmp_res->next = node;
				tmp_res = tmp_res->next;
		}
		if (result->val == 500)
			return result->next;
		else 
			return result;
    }
};
