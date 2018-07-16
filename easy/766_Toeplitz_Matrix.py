"""
A matrix is Toeplitz if every diagonal from top-left to bottom-right has the same element.

Now given an M x N matrix, return True if and only if the matrix is Toeplitz.
 

Example 1:

Input:
matrix = [
  [1,2,3,4],
  [5,1,2,3],
  [9,5,1,2]
]
Output: True
Explanation:
In the above grid, the diagonals are:
"[9]", "[5, 5]", "[1, 1, 1]", "[2, 2, 2]", "[3, 3]", "[4]".
In each diagonal all elements are the same, so the answer is True.
Example 2:

Input:
matrix = [
  [1,2],
  [2,2]
]
Output: False
Explanation:
The diagonal "[1, 2]" has different elements.

Note:

matrix will be a 2D array of integers.
matrix will have a number of rows and columns in range [1, 20].
matrix[i][j] will be integers in range [0, 99].

Follow up:

What if the matrix is stored on disk, and the memory is limited such that you can only load at most one row of the matrix into the memory at once?
What if the matrix is so large that you can only load up a partial row into the memory at once?

"""

class Solution:
    def isToeplitzMatrix(self, matrix):
        """
        :type matrix: List[List[int]]
        :rtype: bool
        """

        for i in range(0,len(matrix)-1):
            if matrix[i][:-1] != matrix[i+1][1:]:
                return False
        return True

def main():
    matrix = [[1,2,3,4], [5,1,2,3], [9,5,1,2]]
    print(matrix)
    print(Solution().isToeplitzMatrix(matrix))

if __name__ == '__main__':
    main()


"""

目标:   判断一个矩阵是不是Toeplitz Matrix(托普利兹矩阵)
        这个矩阵的特点是, 所有对角线(左上到右下)上的元素是一样的

 方法:
    1.  仔细观察这个matrix
            matrix = [
            [1,2,3,4],
            [5,1,2,3],
            [9,5,1,2]
            ]
        会发现, 这是一个托普利兹矩阵(废话_)
    2.  仔细看, 会发现, 这个矩阵的每一行和它的下一行有关系, 有一些东西是一样的
    3.  比如, 第一行的1, 2, 3和第二行的1,2,3一样
        第二行的5,1,2和第三行的5,1,2一样
    4.  得出结论, 每相邻两行的重叠部分是重点
    5.  那么, 只需要判断相邻两行的部分就好了
    6.  如果martix[这一行][从头到倒数第一个] 不等于 martix[下一行][从第二个到末尾]
        那么返回False

思考:   这道题做了好久, 总在纠结每个元素的位置关系, 反而忽略了整体上的联系
        看完大牛的代码醍醐灌顶
        菜逼玩个蛋的游戏
        去写代码吧
        PS: 游戏都删了, 找到工作以前谁玩谁是狗

"""