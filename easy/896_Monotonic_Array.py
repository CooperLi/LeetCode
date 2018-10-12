"""
An array is monotonic if it is either monotone increasing or monotone decreasing.

An array A is monotone increasing if for all i <= j, A[i] <= A[j].  An array A is monotone decreasing if for all i <= j, A[i] >= A[j].

Return true if and only if the given array A is monotonic.

 

 Example 1:

 Input: [1,2,2,3]
 Output: true
 Example 2:

 Input: [6,5,4,4]
 Output: true
 Example 3:

 Input: [1,3,2]
 Output: false
 Example 4:

 Input: [1,2,4,5]
 Output: true
 Example 5:

 Input: [1,1,1]
 Output: true
  

  Note:

  1 <= A.length <= 50000
  -100000 <= A[i] <= 100000
 """
 class Solution:
    def isMonotonic(self, A):
        """
        :type A: List[int]
        :rtype: bool
        """
        return A == sorted(A) or A == sorted(A,reverse=True)
         


"""
目标: 检测一个array是不是连续上升或者连续下降的序列(包括相等)

方法: 把序列排序, 如果原序列和排序过后的序列相等或者等于已排序序列的倒序, 那么就返回True

思考: 考虑连续的上升或者下降就排序, 然后比较两个序列是不是相等
"""
