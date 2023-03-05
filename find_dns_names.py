import csv
import socket

input_file = 'input.csv'
output_file = 'names.csv'
delimiter = ','
encoding = 'utf-8-sig'

dns_servers = set()

with open(input_file, 'r', encoding=encoding) as infile, \
        open(output_file, 'w') as outfile:
    reader = csv.DictReader(infile, delimiter=delimiter)
    writer = csv.writer(outfile)
    writer.writerow(['URL', 'Domain', 'DNS'])

    for row in reader:
        url = row['URL']
        domain = url.split('//')[-1].split('/')[0]
        try:
            ip = socket.gethostbyname(domain)
            dns = socket.getfqdn(ip)
            dns_servers.add(dns)
            writer.writerow([url, domain, dns])
        except socket.gaierror:
            writer.writerow([url, domain, ''])
            continue

print(f'Found {len(dns_servers)} unique DNS servers.')
