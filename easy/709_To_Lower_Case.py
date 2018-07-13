"""
Implement function ToLowerCase() that has a string parameter str, and returns the same string in lowercase.
"""

class Solution:
    def toLowerCase(self, str):
        """
        :type str: str
        :rtype: str
        """
        return str.lower()

def main():
    j = 'HELLO'
    print(Solution().toLowerCase(j))

if __name__ == '__main__':
    main()

"""
这里用了python 的str自带函数 lower()
没啥说的
"""
