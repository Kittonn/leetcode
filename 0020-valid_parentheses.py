class Solution:
    def isValid(self, s: str) -> bool:
        stack = []

        paren_map = {
          '}': '{',
          ')': '(',
          ']': '['
        }
        
        for i in s:
          if i in paren_map:
            if  stack and stack[len(stack) - 1] == paren_map.get(i) :
              stack.pop()
            else:
              return False
          else:
            stack.append(i)
        
        return not stack