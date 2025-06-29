class Solution:
    def isAnagram(self, s: str, t: str) -> bool:
        hashmap = {}
        
        if len(s) != len(t):
          return False
        
        for i in s:
          hashmap[i] = hashmap.get(i, 0) + 1
          
        for i in t:
          hashmap[i] = hashmap.get(i, 0) - 1
          
          if hashmap.get(i) < 0:
            return False
          
        return True
          
        
s = Solution()
print(s.isAnagram('ab', 'a'))