"""
Given an array nums, write a function to move all 0's to the end of it while maintaining the relative order of the non-zero elements.

Example:

Input: [0,1,0,3,12]
Output: [1,3,12,0,0]
Note:

You must do this in-place without making a copy of the array.
Minimize the total number of operations.
"""

class Solution:
    def moveZeroes(self, nums):
        """
        :type nums: List[int]
        :rtype: void Do not return anything, modify nums in-place instead.
        """
        for i in range(len(nums))[::-1]:
            if nums[i] == 0:
                nums.append(nums.pop(i))



"""
目标:   将 list 中的0全部移到最后, 而且保持非0元素间的顺序

方法:   
    1.  用了 pop 的方法, 找到是0的位置元素i, 然后 pop(i), 接着在 append()到 list 中
    2.  注意! 这里在进入 for 循环的时候, 顺序不能是从0-len()
        因为那样, 边 pop 边插入, 后面的顺序都会被打乱, 就会造成跳过某些元素的问题
    3.  使用 range(x)[::-1], 就成了从后往前数数了, 就避免了跳过某些元素的问题

思考:
    1.  数据操作, 用到 pop 这个知识
    2.  逆序, 用到了 range(x)[::-1]. [::-1] 用在range 后面还是第一次
    3.  上面的代码效率不高, 因为每次pop 和 append 之后数据都要移动, 造成浪费, 下面是双指针方法:
        值得注意的是, 在 if nums[i] 这句话中, if 后面直接跟数字: 只有0才是 False, 非0就是 True
        如果 if 后面跟的字符串, 那么就是空字符串为 False, 非空为 True
"""

class Solution2:
    def moveZeroes(self, nums):
        """
        :type nums: List[int]
        :rtype: void Do not return anything, modify nums in-place instead.
        """
        pos = 0
        for i in range(len(nums)):
            if nums[i]:
                print(nums[i])
                nums[pos] = nums[i]
                pos += 1

        for i in range(pos, len(nums)):
            nums[i] = 0


def main():
    nums = [0,1,0,3,12]
    print(nums)
    s1, s2 = Solution(), Solution2()
    # s1.moveZeroes(nums)
    s2.moveZeroes(nums)
    print(nums)

if __name__ == '__main__':
    main()