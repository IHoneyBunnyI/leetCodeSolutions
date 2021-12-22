#include <iostream>
#include <vector>
#include <set>
#include <utility>
#include <algorithm>
#include <map>

using namespace std;

class Solution {
	private:
		bool gorisontal_lines(vector<vector<char>>& board)
		{
			for (int i = 0; i < 9; i++)
			{
				std::set<char> t_set;
				for (int j = 0; j < 9; j++)
				{
					std::pair<std::set<char>::iterator, bool> p;
					p.second = true;
					if (board[i][j] != '.')
						p = t_set.insert(board[i][j]);
					if (p.second != true)
						return false;
				}
			}
			return true;
		}

		bool vertical_lines(vector<vector<char>>& board)
		{
			for (int i = 0; i < 9; i++)
			{
				std::set<char> t_set;
				for (int j = 0; j < 9; j++)
				{
					std::pair<std::set<char>::iterator, bool> p;
					p.second = true;
					if (board[j][i] != '.')
						p = t_set.insert(board[j][i]);
					if (p.second != true)
						return false;
				}
			}
			return true;
		}

		bool rectangle(vector<vector<char>>& board, int x1, int y1, int x2, int y2)
		{
			std::set<char> t_set;
			for (int i = y1; i < y2; i++)
			{
				for (int j = x1; j < x2; j++)
				{
					std::pair<std::set<char>::iterator, bool> p;
					p.second = true;
					if (board[i][j] != '.')
						p = t_set.insert(board[i][j]);
					if (p.second != true)
						return false;
				}
			}
			return true;
		}

public:
	bool isValidSudoku(vector<vector<char>>& board) {
		bool res;
		if (!(res = gorisontal_lines(board)))
			return false;
		if (!(res = vertical_lines(board)))
			return false;
		if (!(res = rectangle(board, 0, 0, 3, 3)))
			return false;
		if (!(res = rectangle(board, 3, 0, 6, 3)))
			return false;
		if (!(res = rectangle(board, 6, 0, 9, 3)))
			return false;

		if (!(res = rectangle(board, 0, 3, 3, 6)))
			return false;
		if (!(res = rectangle(board, 3, 3, 6, 6)))
			return false;
		if (!(res = rectangle(board, 6, 3, 9, 6)))
			return false;

		if (!(res = rectangle(board, 0, 6, 3, 9)))
			return false;
		if (!(res = rectangle(board, 3, 6, 6, 9)))
			return false;
		if (!(res = rectangle(board, 6, 6, 9, 9)))
			return false;
		return true;
	}
};
