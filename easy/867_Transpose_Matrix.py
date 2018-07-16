"""

Given a matrix A, return the transpose of A.

The transpose of a matrix is the matrix flipped over it's main diagonal, switching the row and column indices of the matrix.

 

Example 1:

Input: [[1,2,3],[4,5,6],[7,8,9]]
Output: [[1,4,7],[2,5,8],[3,6,9]]
Example 2:

Input: [[1,2,3],[4,5,6]]
Output: [[1,4],[2,5],[3,6]]
 

Note:

1 <= A.length <= 1000
1 <= A[0].length <= 1000

"""


class Solution:
    def transpose(self, A):
        """
        :type A: List[List[int]]
        :rtype: List[List[int]]
        """
        result = [[None for i in range(len(A))] for j in range(len(A[0]))]
        for i in range(len(A)):
            for j in range(len(A[0])):
                result[j][i] = A[i][j]
        return result


def main():
    A = [[1,2,3],[4,5,6],[7,8,9]]
    print(A)
    print(Solution().transpose(A))

if __name__ == '__main__':
    main()


"""

目标:   二维矩阵变换, 详细的例子在上方
方法:
    1.  创建一个新的list, 大小是目标大小.
        比如原矩阵是2, 3 (有两个子list, 每个list中有3个元素)
        那么新的list就是3, 2(有三个子list, 每个list中有2个元素)
    2.  然后就是两个for循环, 第一个for是按照原list的子list个数
        第二个for是按照原list中每个子list中的元素个数
    3.  把原list中的[i][j] 赋值给新的[j][i]
思考:   
    1.  这道题想明白了意思, 但是再赋值之前, 没有创建好result这个新的list
        最初设置的新list就是result = [], 没有考虑到这是一个矩阵
        然后看了网上的代码, 发现了创建方法
    2.  其实网上的一句话代码可以在创建新list的基础上做改动:
        return [[A[i][j] for i in range(len(A[0]))] for j in range(len(A))]
    3.  除此之外, 甚至有更简单的, 用zip这个自带函数. 因为没有学过, 所以根本不知道怎么应用
        return list(zip(*A))
        但这样返回的虽然是一个list, 但是list中的每一个子元素是一个tuple而不仍是list. 
    4.  更好的方法是:
        return list(map(list, zip(*A)))
        这样返回的就都是list了, 因为用了map这个映射
    5.  help(zip)

        class zip(object)
        |  zip(iter1 [,iter2 [...]]) --> zip object
        |
        |  Return a zip object whose .__next__() method returns a tuple where
        |  the i-th element comes from the i-th iterable argument.  The .__next__()
        |  method continues until the shortest iterable in the argument sequence
        |  is exhausted and then it raises StopIteration.

        list.zip() 函数用于将可迭代的对象作为参数,将对象中对应的元素打包成一个个元组,然后返回由这些元组组成的列表.
        如果各个迭代器的元素个数不一致,则返回列表长度与最短的对象相同
        利用 * 号操作符,可以将元组解压为列表.        
        在 Python 3.x 中为了减少内存, zip() 返回的是一个对象. 如需展示列表, 需手动 list() 转换.

        >>> x = [1, 2, 3]
        >>> y = [4, 5, 6]
        >>> z = [7, 8, 9]
        >>> xyz = zip(x, y, z)
        >>> print(list(xyz))
        [(1, 4, 7), (2, 5, 8), (3, 6, 9)]

"""