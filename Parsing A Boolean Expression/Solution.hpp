#include <iostream>
#include <vector>
#include <set>
#include <utility>

using namespace std;
#define NO 1
#define OR 2
#define AND 3

class Solution {
	private:
	bool check_str(std::string str)
	{
		std::string alphabet = "()&|!ft, ";
		for (int i = 0; i < str.size(); i++)
			if (alphabet.find(str[i]) == string::npos)
			{
				std::cout << "ERR" << std::endl;
				return false;
			}
		int open = 0;
		int close = 0;
		for (int i = 0; i < str.size(); i++)
		{
			if (str[i] == '(')
				open++;
			if (str[i] == ')')
				close++;
		}
		if (open != close)
		{
			std::cout << "ERR" << std::endl;
			return false;
		}
		return true;
	}

	bool two_brackets(int start, int end, std::string &str)
	{
		int open_br = 0;
		int close_br = 0;
		for (int i = start + 1; i < end; i++)
		{
			if (str[i] == '(')
				open_br++;
			if (str[i] == ')')
				close_br++;
		}
		if (open_br == 1 || close_br == 1)
			return true;
		return false;
	}

	int find_operation(std::string &str)
	{
		std::string opers = "!|&";
		int res = -1;
		for (int i = 0; i < 3; i++)
		{
			res = str.find(opers[i]);
			if (res != string::npos)
				break;
		}
		if (str[res] == '|')
			return OR;
		if (str[res] == '&')
			return AND;
		if (str[res] == '!')
			return NO;
		return 0;
	}

	bool calculate_operation(std::string &str)
	{
		std::vector<int> numbers;
		int operation = find_operation(str);
		if (operation == 0)
			return 0;
		for (int i = str.find('('); str[i] != ')'; i++)
		{
			//std::cout << str[i];
			if (str[i] == 't')
				numbers.push_back(1);
			else if (str[i] == 'f')
				numbers.push_back(0);
		}
		//for (int i = 0; i < numbers.size(); i++)
			//std::cout << numbers[i] << std::endl;
		//std::cout << operation;

		//
		bool res;
		int i_res;
		if (operation == OR)
		{
			i_res = numbers[0];
			for (int i = 0; i < numbers.size(); i++)
				if (numbers[i] == 1)
					i_res = 1;
		}
		if (operation == AND)
		{
			i_res = numbers[0];
			for (int i = 0; i < numbers.size(); i++)
				if (numbers[i] == 0)
					i_res = 0;
		}
		if (operation == NO)
		{
			i_res = !numbers[0];
		}
		if (i_res == 0)
			res = false;
		if (i_res == 1)
			res = true;
		//
		return res;
	}

	char next_br(std::string s)
	{
		for (int i = 0; i < s.size(); i++)
		{
			if (s[i] == '(')
				return 0;
			if (s[i] == ')')
				return ')';
		}
		return 0;
	}

	int close_bracket(std::string &str)
	{
		int i = 0;
		for (i = 0; str[i] != ')'; i++) ;
		return i;
	}

	int open_bracket(std::string& str, int close)
	{
		int i = 0;
		for (i = close; str[i] != '('; i--) ;
		return i;
	}

	int choice_operation(char c)
	{
		if (c == '!')
			return NO;
		if (c == '&')
			return AND;
		else
			return OR;
	}

	std::string calculate_and_erase(std::string str, int start, int end)
	{
		std::vector<int> numbers;
		int operation = choice_operation(str[start - 1]);
		if (operation == 0)
			return 0;
		for (int i = start; i < end; i++)
		{
			if (str[i] == 't')
				numbers.push_back(1);
			else if (str[i] == 'f')
				numbers.push_back(0);
		}

		bool res;
		int i_res;
		if (operation == OR)
		{
			i_res = numbers[0];
			for (int i = 0; i < numbers.size(); i++)
				if (numbers[i] == 1)
					i_res = 1;
		}
		if (operation == AND)
		{
			i_res = numbers[0];
			for (int i = 0; i < numbers.size(); i++)
				if (numbers[i] == 0)
					i_res = 0;
		}
		if (operation == NO)
		{
			i_res = !numbers[0];
		}
		if (i_res == 0)
			res = false;
		if (i_res == 1)
			res = true;

		std::string new_str(str, 0, start - 1);
		if (res == true)
			new_str.push_back('t');
		if (res == false)
			new_str.push_back('f');
		new_str.insert(new_str.end(), str.begin() + end + 1, str.end());
		//std::cout << new_str << std::endl;
		return new_str;
	}

public:
	/*
	 * 1) - !
	 * 2) - &
	 * 3) - |
	 */

	bool parseBoolExpr(string str) {
		if (!check_str(str))
			return false;
		while (!two_brackets(0, str.size(), str))
		{
			int close_br = close_bracket(str);
			int open_br = open_bracket(str, close_br);
			str = calculate_and_erase(str, open_br, close_br);
		}
		return calculate_operation(str);
	}
};
