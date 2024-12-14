class Solution:
    def romanToInt(self, s: str) -> int:
        roman_to_int = {
            'I': 1,
            'V': 5,
            'X': 10,
            'L': 50,
            'C': 100,
            'D': 500,
            'M': 1000
        }

        result = 0

        for i in range(len(s) - 1, -1, -1):
            curr = roman_to_int.get(s[i])
            
            if i + 1 < len(s) and curr < roman_to_int.get(s[i + 1]):
              result -= curr
            else:
              result += curr

        return result
      
