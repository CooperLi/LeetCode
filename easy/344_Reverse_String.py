"""
Write a function that takes a string as input and returns the string reversed.

Example:
Given s = "hello", return "olleh".
"""

class Solution:
    def reverseString(self, s):
        """
        :type s: str
        :rtype: str
        """
        return s[::-1]

def main():
    s = 'hello'
    print(s)
    print(Solution().reverseString(s))

if __name__ == '__main__':
    main()

"""
实在没得说, 用了[::-1] 这个反转的方法
"""