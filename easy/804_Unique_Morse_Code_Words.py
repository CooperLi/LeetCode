"""
International Morse Code defines a standard encoding where each letter is mapped to a series of dots and dashes, as follows: "a" maps to ".-", "b" maps to "-...", "c" maps to "-.-.", and so on.

For convenience, the full table for the 26 letters of the English alphabet is given below:

[".-","-...","-.-.","-..",".","..-.","--.","....","..",".---","-.-",".-..","--","-.","---",".--.","--.-",".-.","...","-","..-","...-",".--","-..-","-.--","--.."]

Now, given a list of words, each word can be written as a concatenation of the Morse code of each letter. For example, "cab" can be written as "-.-.-....-", (which is the concatenation "-.-." + "-..." + ".-"). We'll call such a concatenation, the transformation of a word.

Return the number of different transformations among all words we have.

Example:
Input: words = ["gin", "zen", "gig", "msg"]
Output: 2
Explanation: 
The transformation of each word is:
"gin" -> "--...-."
"zen" -> "--...-."
"gig" -> "--...--."
"msg" -> "--...--."

There are 2 different transformations, "--...-." and "--...--.".

Note:

The length of words will be at most 100.
Each words[i] will have length in range [1, 12].
words[i] will only consist of lowercase letters.
"""

class Solution:
    def uniqueMorseRepresentations(self, words):
        """
        :type words: List[str]
        :rtype: int
        """
        morse = [".-","-...","-.-.","-..",".","..-.","--.","....","..",".---","-.-",".-..","--","-.","---",".--.","--.-",".-.","...","-","..-","...-",".--","-..-","-.--","--.."]
        trans = lambda x: morse[ord(x) - ord('a')]
        map_to_word = lambda word: ''.join([trans(x) for x in word])
        result = map(map_to_word, words)
        return len(set(result))

def main():
    words = ['gin', 'zen', 'gig', 'msg']
    print(Solution().uniqueMorseRepresentations(words))

if __name__ == '__main__':
    main()



"""
目标: 找出 List 中的 word 一共有几种 morse 表达方式
方法: 
    1. 先把字母 a,b,c,d,...,z 和摩尔斯配对, 这里用的是 trans 方法, 将字母和摩尔斯一一配对; ord 在这里的作用是返回字母的 ascii 值, 比如 ord('a') = 97
    2. 字母配对后, 把每个词和摩尔斯配对(注意, 这里说的是每个词, 之前 trans 方法是将每个字母配对). 这里的 map_to_word 中, 用了 join 这个 string 中自带的方法, 把每个词中已经转换成摩尔斯的字母连接起来
    3. 之后使用 map 方法, 将 map_to_word 应用在 words 中的每个元素上
    4. 最后用一个 set 来存储结果, 因为 set 自带去重, 所以返回值唯一
    5. 返回 set 的大小就是题目的结果
思考: 熟悉使用数据结构, lambda表达式, 高级编程思维
"""