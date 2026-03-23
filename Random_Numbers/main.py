import csv
import random

output_file = "random_numbers.csv"

with open(output_file, "w", newline="") as f:
    writer = csv.writer(f)
    writer.writerow(["number"])
    for _ in range(200):
        writer.writerow([random.randint(1, 6)])

print(f"Generated 200 random numbers in '{output_file}'")
