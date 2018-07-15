"""
Given a binary matrix A, we want to flip the image horizontally, then invert it, and return the resulting image.

To flip an image horizontally means that each row of the image is reversed.  For example, flipping [1, 1, 0] horizontally results in [0, 1, 1].

To invert an image means that each 0 is replaced by 1, and each 1 is replaced by 0. For example, inverting [0, 1, 1] results in [1, 0, 0].

Example 1:

Input: [[1,1,0],[1,0,1],[0,0,0]]
Output: [[1,0,0],[0,1,0],[1,1,1]]
Explanation: First reverse each row: [[0,1,1],[1,0,1],[0,0,0]].
Then, invert the image: [[1,0,0],[0,1,0],[1,1,1]]
Example 2:

Input: [[1,1,0,0],[1,0,0,1],[0,1,1,1],[1,0,1,0]]
Output: [[1,1,0,0],[0,1,1,0],[0,0,0,1],[1,0,1,0]]
Explanation: First reverse each row: [[0,0,1,1],[1,0,0,1],[1,1,1,0],[0,1,0,1]].
Then invert the image: [[1,1,0,0],[0,1,1,0],[0,0,0,1],[1,0,1,0]]
Notes:

1 <= A.length = A[0].length <= 20
0 <= A[i][j] <= 1
"""

class Solution:
    def flipAndInvertImage(self, A):
        """
        :type A: List[List[int]]
        :rtype: List[List[int]]
        """
        flip = list(map(lambda x: x[::-1], A))
        for i in range(len(flip)):
            for j in range(len(flip[i])):
                flip[i][j] = 1 - flip[i][j]
        return flip


def main():
    A = [[1,1,0,0],[1,0,0,1],[0,1,1,1],[1,0,1,0]]
    print(Solution().flipAndInvertImage(A))
    
if __name__ == "__main__":
    main()

 

"""
目标:   先把list中的每个元素从后往前 reverse (颠倒)一下, 然后把所有的元素 inverse (取反)

方法:
    1.  先reverse. 用一个lambda, 用array中的[::-1]代表倒序输出;
        然后用map这个方法应用到A中的所有元素
    2.  处理好的map结果是一个map object, 要用list() 转换成list
    3.  对于每位取反, 用了一个取巧的方法: 1 -> 0, 0 -> 1
        不就是用1 减每位数吗? 如果这个数是1, 那么1-1=0, 如果是0, 那么1-0=1
        这样就得到了取反后的结果

思考:   还是不够熟悉数组操作, 看大牛各种一句话代码, 泪奔...
        而且取反操作没有普适性

附上一句话代码:

class Solution:
    def flipAndInvertImage(self, A):

       return [list(map(lambda x: x ^ 1, item[::-1])) for item in A]

"""