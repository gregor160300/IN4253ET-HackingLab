import csv

# Open the input CSV file
with open('input.csv', 'r') as infile:
    reader = csv.reader(infile)

    # Create a list of modified rows
    rows = []
    for row in reader:
        modified_row = [line[1:-1] for line in row]  # remove first and last character of every line
        rows.append(modified_row)

# Write the modified rows to a new CSV file
with open('output.csv', 'w', newline='') as outfile:
    writer = csv.writer(outfile)
    writer.writerows(rows)
