#include <iostream>
#include "Solution.hpp"

int main()
{
	Solution s;
	std::cout << s.parseBoolExpr("!(t)") << std::endl;
	std::cout << s.parseBoolExpr("|(t,f)") << std::endl;
}
