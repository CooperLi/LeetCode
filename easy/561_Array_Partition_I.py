"""

Given an array of 2n integers, your task is to group these integers into n pairs of integer, say (a1, b1), (a2, b2), ..., (an, bn) which makes sum of min(ai, bi) for all i from 1 to n as large as possible.

Example 1:
Input: [1,4,3,2]

Output: 4
Explanation: n is 2, and the maximum sum of pairs is 4 = min(1, 2) + min(3, 4).
Note:
n is a positive integer, which is in the range of [1, 10000].
All the integers in the array will be in the range of [-10000, 10000].

"""


class Solution:
    def arrayPairSum(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """
        return sum(sorted(nums)[::2])

def main():
    nums = [1, 4, 3, 2]
    print(nums)
    print(Solution().arrayPairSum(nums))

if __name__ == '__main__':
    main()
        

"""

目标:   输出一组大小为2n的数组中所有两两配对的最小值之和的最大值. 这么说有点乱...
        举个例子: 当n=2时, 数组A的大小为2*2 = 4
        A = [1, 4, 3, 2]
        输出A中元素两两配对的最小值之和, 这个和要最大
        b = [1, 2], c = [3, 4]
        那么 sum(min(b) , min(c)) = 4, 这个是所有组合最大的

方法:   1.  先给A排序(从大到小). 用sorted 这个自带函数
            默认是从小到大排序, 如果从大到小, 用sorted(A, reverse=True)
        2.  然后按照间距为2输出这个数列的和, 间距为2表示为[::2]
            即sum(sorted(A)[::2])

思考:   
        1.  为什么要排序?
            因为这里要求和最大, 而且是两两配对的最小值之和
            那么如果我这个list是有序的, 那么两两配对中, 左边的数肯定是我要的, 因为用了min()这个方法
        2.  为什么用[::2]? 
            因为既然数组已经排好(从小到大) 那么所有组合中, 我真正需要的就是每个组合中左边的数字
            而且数组的大小为2, 那就不管右边是几, 无脑输出左边的数字, 所以就步长为2的跳跃
        3.  最后用sum把这些数字加起来, 就肯定是所有组合中最大的了
        4.  步骤拆分, 还是 A = [1, 4, 3, 2]
            sorted(A) = [1, 2, 3, 4]
            sorted(A)[::2] = [1, 3]
            sum(sorted(A)[::2]) = 4
            反推, 所有的组合方式: 
            [1, 4][3, 2] -> 1 + 2 = 3
            [1, 3][2, 4] -> 1 + 2 = 3
            [1, 2][3, 4] -> 1 + 3 = 4 -> biggest

"""