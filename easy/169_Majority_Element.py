"""
Given an array of size n, find the majority element. The majority element is the element that appears more than ⌊ n/2 ⌋ times.

You may assume that the array is non-empty and the majority element always exist in the array.

Example 1:

Input: [3,2,3]
Output: 3
Example 2:

Input: [2,2,1,1,1,2,2]
Output: 2
"""

class Solution:
    def majorityElement(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """
        dict = {}
        for i in set(nums):
            dict[i] = nums.count(i)
        # return max(dict, key=lambda x: dict[x])
        return max(dict, key=dict.get)
        
def main():
    # nums = [2,2,1,1,1,2,2]
    nums = [3,3,4]
    print(nums)
    print(Solution().majorityElement(nums))

if __name__ == '__main__':
    main()


"""
目标:   返回最多出现的元素

方法:   
    1.  创建一个dictionary, 按照{'元素': 元素出现次数} 依次填入数据
    2.  循环中用的set, 避免重复元素出现
    3.  用nums.count(i) 来统计每个元素出现次数
    4.  最后输出字典最大value对应的key
        这里有两种方法: max(dict, key=dict.get) 和 max(dict, key=lambda x: dict[x])

思考:   又一道一遍过的题, 信心有了啊, 只不过明早3点要起来去海鲜市场买海鲜, 今天估计睡不了了.
        看了大牛的提交, 觉得自己写的还是太罗嗦:

        return [x for x in set(nums) if nums.count(x) > len(nums) // 2][0]
        主要的思想差不多, 但是这个答案没有用dict, 直接用list存那个具体的值是多少, 比我的强
        (我为什么要多此一举用上dict这种吃力不讨好的东西呢???)

一遍过
"""