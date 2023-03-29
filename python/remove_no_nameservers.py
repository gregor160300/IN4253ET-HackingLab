import argparse
import csv

# Parse command line arguments
parser = argparse.ArgumentParser(
    description='Remove rows with no nameservers from a CSV file')
parser.add_argument('input_file', help='Input CSV file name')
parser.add_argument('output_file', help='Output CSV file name')
args = parser.parse_args()

# Open input and output files
with open(args.input_file, 'r') as csv_in_file, open(args.output_file, 'w') as csv_out_file:
    # Create CSV reader and writer objects
    csv_reader = csv.DictReader(csv_in_file)
    csv_writer = csv.DictWriter(csv_out_file, fieldnames=csv_reader.fieldnames)

    # Write header row to output file
    csv_writer.writeheader()

    # Iterate over input rows and write to output file if there are nameservers
    for row in csv_reader:
        if not row['nameservers']:
            csv_writer.writerow(row)
