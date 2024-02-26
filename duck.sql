SELECT column0, MIN(column1), AVG(column1), MAX(column1)
FROM read_csv("data.csv")
GROUP BY column0
