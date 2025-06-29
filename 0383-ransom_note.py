from collections import Counter

class Solution:
    def canConstruct(self, ransomNote: str, magazine: str) -> bool:
        ransom_note_hashmap = Counter(ransomNote)
        magazine_hashmap = Counter(magazine)

        for item, counter in ransom_note_hashmap.items():
          if magazine_hashmap[item] < counter:
            return False
          
        return True
      
  