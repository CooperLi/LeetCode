"""
In MATLAB, there is a very useful function called 'reshape', which can reshape a matrix into a new one with different size but keep its original data.

You're given a matrix represented by a two-dimensional array, and two positive integers r and c representing the row number and column number of the wanted reshaped matrix, respectively.

The reshaped matrix need to be filled with all the elements of the original matrix in the same row-traversing order as they were.

If the 'reshape' operation with given parameters is possible and legal, output the new reshaped matrix; Otherwise, output the original matrix.

Example 1:
Input: 
nums = 
[[1,2],
 [3,4]]
r = 1, c = 4
Output: 
[[1,2,3,4]]
Explanation:
The row-traversing of nums is [1,2,3,4]. The new reshaped matrix is a 1 * 4 matrix, fill it row by row by using the previous list.
Example 2:
Input: 
nums = 
[[1,2],
 [3,4]]
r = 2, c = 4
Output: 
[[1,2],
 [3,4]]
Explanation:
There is no way to reshape a 2 * 2 matrix to a 2 * 4 matrix. So output the original matrix.
Note:
The height and width of the given matrix is in range [1, 100].
The given r and c are all positive.

"""


class Solution:
    def matrixReshape(self, nums, r, c):
        """
        :type nums: List[List[int]]
        :type r: int
        :type c: int
        :rtype: List[List[int]]
        """
        print('Original List: ', nums)
        print('Wanted row: ', r)
        print('Wanted col: ', c)
        if len(nums) * len(nums[0]) != r * c:
            return nums

        tmp = []
        for i in range(len(nums)):
            for j in range(len(nums[0])):
                tmp.append(nums[i][j])
        
        result = []
        for i in range(r):
            result.append([])
            print('i result: ', result)
            for j in range(c):
                result[i].append(tmp[i*c+j])
                print('j result: ', result)
        return result

def main():
    nums = [[1,2,3], [4,5,6], [7,8,9], [10,11,12]]
    print(Solution().matrixReshape(nums,6,2))

if __name__ =='__main__':
    main()


"""
目标:   将一个矩阵随意变换成用户希望的形状
        如: 用户输入2, 4(即两行四列)
        那么原矩阵如果满足转换条件, 则变换
        如果不满足, 则返回原矩阵
        具体例子见上方

方法:
    1.  先判断用户的输入是否满足可变换的条件
        这里的可变换条件就是, r * c =? len(list) * len(list[0])
        这里隐形规定原矩阵中每个子 list 中的元素个数是相同的
        要不然不会输出原矩阵, 只会输出删掉多于元素的矩阵
    2.  先将矩阵拉成一条
        具体的意思就是, 把原矩阵中所有的元素都抽出来放进一个临时矩阵tmp里面
    3.  再根据用户的r,c 来进入循环(说明有两层循环, 一层 r, 一层 c)
        外层循环中, 只是插入新的空的子 list, 一共插入多少个, 看 r 是几.
        r 如果是6, 就说明新的 result 这个矩阵里面要有6个子 list
        内层循环中, 插入临时矩阵 tmp[i*c+j]那个元素到最终的result里面.
    4.  用中间插入 print 的方法可以更好地理解程序运行的逻辑

思考:   想问题不要先陷入局部, 宏观的感觉有时候更重要
    
"""