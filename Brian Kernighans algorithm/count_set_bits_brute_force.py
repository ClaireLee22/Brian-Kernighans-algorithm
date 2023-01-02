class Solution(object):
    def hammingWeight(self, n):
        """
        :type n: int
        :rtype: int
        """
        count = 0
        while n:
            count += n & 1 # check last bit
            n = n >> 1
            
        return count

if __name__ == "__main__":
    solution = Solution()
    number = 22
    print("{0:d} in binary is {1:b}".format(number, number))
    print("the number of set bits in {0:d} is {1:d}".format(number, solution.hammingWeight(number)))
    
"""
22 in binary is 10110
the number of set bits in 22 is 3
"""