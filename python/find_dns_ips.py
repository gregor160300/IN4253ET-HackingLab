import csv
import socket

# define input and output file paths
input_file = 'input.csv'
output_file = 'output.csv'

# define set to store unique DNS servers
dns_servers = set()

# open input file and read URLs from the URL column
with open(input_file, 'r', encoding='utf-8-sig') as csvfile:
    reader = csv.DictReader(csvfile, delimiter=',')
    for row in reader:
        url = row['URL']

        # extract the hostname from the URL
        hostname = url.split('//')[1].split('/')[0]

        # get the list of IP addresses for the hostname
        try:
            ip_list = socket.getaddrinfo(hostname, 80)
        except:
            print("Website: " + hostname + " doesn't exist")

        # extract the DNS server IP address from the first IP address in the list
        dns_server = ip_list[0][4][0]

        # add the DNS server to the set of unique DNS servers
        dns_servers.add(dns_server)

# write the unique DNS servers to the output file
with open(output_file, 'w', newline='') as csvfile:
    writer = csv.writer(csvfile)
    writer.writerow(['DNS Server'])
    for dns_server in dns_servers:
        writer.writerow([dns_server])
