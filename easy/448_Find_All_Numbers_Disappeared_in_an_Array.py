"""
Given an array of integers where 1 ≤ a[i] ≤ n (n = size of array), some elements appear twice and others appear once.

Find all the elements of [1, n] inclusive that do not appear in this array.

Could you do it without extra space and in O(n) runtime? You may assume the returned list does not count as extra space.

Example:

Input:
[4,3,2,7,8,2,3,1]

Output:
[5,6]
"""

class Solution:
    def findDisappearedNumbers(self, nums):
        """
        :type nums: List[int]
    :rtype: List[int]
        """
        return list(set(range(1,len(nums)+1)) - set(nums))



def main():
    nums = [4,3,2,7,8,2,3,1]
    print(nums)
    print(Solution().findDisappearedNumbers(nums))


if __name__ == '__main__':
    main()


"""
目标:   找到在原 list 中没出现的元素, 以list(1-len(原list)) 和 原 list 比较

方法:   看到 list 之间的比较, 还是不管重复的元素, 那就直接上 set 了
        再用 set(A) - set(b) = 出现在 A 却不出现在 B 的元素
        得到答案

思考:   第一次思路这么清晰地想出来了, 用时不到30秒, 喜欢上做题的感觉了
        还有, 建一个 set, 不用 set([s for s in range(1, len(nums) + 1)])
        上面的这个就是旧思维, 先想着建一个 list, 再转化成 set
        直接 set(range(1, len(nums) + 1)) 就完事儿了
"""