-- https://leetcode.com/problems/capital-gainloss/description/ 
SELECT stock_name , SUM(
    case WHEN operation = "Buy" then -price 
    else price
    end
) as capital_gain_loss from Stocks group by stock_name;
