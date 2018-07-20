"""
We have two special characters. The first character can be represented by one bit 0. The second character can be represented by two bits (10 or 11).

Now given a string represented by several bits. Return whether the last character must be a one-bit character or not. The given string will always end with a zero.

Example 1:
Input: 
bits = [1, 0, 0]
Output: True
Explanation: 
The only way to decode it is two-bit character and one-bit character. So the last character is one-bit character.
Example 2:
Input: 
bits = [1, 1, 1, 0]
Output: False
Explanation: 
The only way to decode it is two-bit character and two-bit character. So the last character is NOT one-bit character.
Note:

1 <= len(bits) <= 1000.
bits[i] is always 0 or 1.

"""


class Solution:
    def isOneBitCharacter(self, bits):
        """
        :type bits: List[int]
        :rtype: bool
        """
        import re
        string = "".join(map(str, bits))
        if '1' in string:
            last = re.search(r'1*0+$',string).group(0)
            print(last)
            print(last.count('1'))
            if last.count('1') % 2 == 0:
                return True
            else:
                print([last[last.count('1')+1:]])
                return '0' in last[last.count('1')+1:]
        return True
        

def main():
    bits = [1, 0, 1,1,1,0,0 ]
    # bits = [1, 0, 0]
    print(bits)
    print(Solution().isOneBitCharacter(bits))

if __name__ == '__main__':
    main()


"""
目标:   有两种字符, 一种是1bit的: 0; 另一种是2bit的, 10 或者 11
        判断一个list中最后一位数是不是0
        比如说, 例子是[1,1,1,0], 那么就只能有一种情况: 11, 10, 输出False
        再比如说, [1,0,0], 那么也只能是10, 0, 输出True, 因为没有单独的1

方法:
    1.  使用正则表达式, import re
    2.  还是老方法, 把list转化成string
    3.  大体判断, 如果没有1在string里面, 那就直接输出True, 因为全是0
    4.  如果有1在string里面, 那么, 挑出符合的:
        正则表达式: 1*0+& 任意多个1后面接至少一个0, 后面没东西了(用$)
        这里最后的group(0), 表示返回满足正则表达式的整个字符串(用last代表)
    5.  如果last中1的个数是偶数, 那么就返回True, 注意, 如果原字符串是[1,0,1,0]
        那么last就是10, 因为有一个$, 这里仔细考虑. 所以如果是偶数, 就肯定是True
        还是这个例子, 如果最后有俩1, 那么最后一个肯定是0, 因为题目规定的最后一个必须是0
    6.  接上面, 如果1不是偶数, 那么返回0是不是在last[last.count('1')+1:]里面
        如果是, 返回True,其实就是问这个0是不是单独的一个0而不是和1组成10的那个0

思考:   熟练运用reg可以快速找到我们要的东西
        网上还有一个更简单的方法

        def isOneBitCharacter(self, bits):
            import re
            return bool(re.match(r'(0|1.)*0$', ''.join(map(str, bits))))

"""