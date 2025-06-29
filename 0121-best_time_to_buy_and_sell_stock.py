from typing import List


class Solution:
    def maxProfit(self, prices: List[int]) -> int:
        max_profit = 0
        sell_day = 1
        buy_day = 0

        while(sell_day < len(prices)):
          if (prices[sell_day] > prices[buy_day]):
            profit = prices[sell_day] - prices[buy_day]

            max_profit = max(max_profit, profit)
          else:
            buy_day = sell_day

          sell_day += 1

        return max_profit

            