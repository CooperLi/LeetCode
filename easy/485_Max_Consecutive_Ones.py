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
目标

"""