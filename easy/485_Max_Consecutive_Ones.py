"""
Given a binary array, find the maximum number of consecutive 1s in this array.

Example 1:
Input: [1,1,0,1,1,1]
Output: 3
Explanation: The first two digits or the last three digits are consecutive 1s.
    The maximum number of consecutive 1s is 3.
Note:

The input array will only contain 0 and 1.
The length of input array is a positive integer and will not exceed 10,000
"""


class Solution:
    def findMaxConsecutiveOnes(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """
        return max(len(x) for x in ["".join([str(s) for s in nums]).split('0')])
        


def main():
    nums = [1,1,0,1,1,1]
    print(nums)
    print(Solution().findMaxConsecutiveOnes(nums))

if __name__=='__main__':
    main()


"""
目标:   输出一个 list 中最多连续出现1的个数

方法:
    1.  把 list 中的所有元素遍历一遍, 并存成一个 string
    2.  用 split('0') 的方法把 list 用0分隔开, 形成若干子 list
    3.  对每个子 list 统计其 len()
    4.  用 max() 找出最大连续1的次数

思考:
    1.  怎么统计连续1的个数?
        重点就是之间夹杂着0, 只要以0为界, 就方便统计
    2.  但是 list 中没有 split 的方法啊?
        所以就先把 list 中的元素全部存成一个 string, 用"".join()
    3.  再用 string 自带的 split(z) 方法, 把 string 分成若干部分(元素均为1)
        最后统计

    可以分步来计算
    [str(s) for s in nums] -> 把 list 中元素变成 string
    "".join([str(s) for s in nums]) -> 把 list 中元素相连变成一个 string
    "".join([str(s) for s in nums]).split("0") -> 以0为界, 分成若干小 string
    ["".join(str(s) for s in nums].split("0"))] -> 结果存成一个 list
    len(x) for x in ["".join(str(s) for s in nums).split("0")] -> 结果成了每个 list 的大小
    max(len(x) for x in ["".join(str(s) for s in nums).split("0")]) -> 返回上面结果的最大值

"""