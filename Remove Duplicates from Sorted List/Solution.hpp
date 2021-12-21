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
		ListNode *result = new ListNode(head->val);
		ListNode *tmp_res = result;
		while (head->next)
		{
			ListNode *next = head->next;
			if (next->val != head->val)
			{
				ListNode *node = new ListNode(next->val);
				tmp_res->next = node;
				tmp_res = tmp_res->next;
			}
			head = head->next;
		}
		return result;
    }
};
