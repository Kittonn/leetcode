class Stack:
  def __init__(self):
    self.stack = []

  def push(self, value):
    return self.stack.append(value)

  def pop(self):
    return self.stack.pop()

  def peek(self):
    return self.stack[-1]

  def is_empty(self):
    return self.stack == []

  def size(self):
    return len(self.stack)


class Solution:
  def isValid(self, s):
    stack = Stack()

    opening = ['(', '[', '{']
    closing = [')', ']', '}']

    for i in s:
      if i in opening:
        stack.push(i)
      elif i in closing:
        if not stack.is_empty() and stack.peek() == opening[closing.index(i)]:
          stack.pop()
        else:
          stack.push(i)

    if stack.is_empty():
      return True
    else:
      return False


Solution().isValid("()[]{}")
