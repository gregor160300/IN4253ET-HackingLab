import csv
import dns.resolver
import dns.message
import dns.rdatatype
import socket

# Open the input and output CSV files
with open('names.csv', newline='') as infile, open('sizes.csv', 'w', newline='') as outfile:
    reader = csv.DictReader(infile)
    writer = csv.DictWriter(outfile, fieldnames=['URL', 'Domain', 'DNS', 'Request Size', 'Response Size', 'TXT Records',
                            'TXT Request Size', 'TXT Response Size', 'DNSSEC', 'DNSSEC Request Size', 'DNSSEC Response Size'])
    writer.writeheader()

    # Iterate over each row in the input CSV file
    for row in reader:
        # Get the domain name from the row
        domain = row['Domain']

        # Query the DNS for the domain with an ANY query
        resolver = dns.resolver.Resolver(configure=False)
        # Resolve the hostname to an IP address
        resolver.nameservers = [socket.gethostbyname(row['DNS'])]
        resolver.use_edns(False)  # Disable EDNS to avoid NoMetaqueries error
        # try:
        #     response = resolver.query(domain, 'ANY')
        # except (dns.resolver.NoAnswer, dns.resolver.LifetimeTimeout):
        #     response = None

        # # Get the request and response sizes
        # request_size = response.request.__sizeof__() if response else 0
        # response_size = response.__sizeof__() if response else 0

        # Check for DNSSEC
        try:
            response_dnssec = resolver.resolve(
                domain, 'DNSKEY', raise_on_no_answer=False)
            dnssec = True
        except (dns.resolver.NoAnswer, dns.resolver.LifetimeTimeout):
            response_dnssec = None
            dnssec = False

        # Get the request and response sizes for DNSSEC query
        dnssec_request_size = response_dnssec.request.__sizeof__() if response_dnssec else 0
        dnssec_response_size = response_dnssec.__sizeof__() if response_dnssec else 0

        # Check for TXT records
        txt_records = []
        try:
            response_txt = resolver.resolve(
                domain, 'TXT', raise_on_no_answer=False)
            for rdata in response_txt:
                txt_records.append(str(rdata))
        except (dns.resolver.NoAnswer, dns.resolver.LifetimeTimeout):
            response_txt = None

        # Get the request and response sizes for TXT query
        txt_request_size = response_txt.request.__sizeof__() if response_txt else 0
        txt_response_size = response_txt.__sizeof__() if response_txt else 0

        # Write the results to the output CSV file
        writer.writerow({
            'URL': row['URL'],
            'Domain': domain,
            # 'DNS': str(response.ns[0]) if response else '',
            # 'Request Size': request_size,
            # 'Response Size': response_size,
            'DNS': '',
            'Request Size': '',
            'Response Size': '',
            'TXT Records': '; '.join(txt_records),
            'TXT Request Size': txt_request_size,
            'TXT Response Size': txt_response_size,
            'DNSSEC': dnssec,
            'DNSSEC Request Size': dnssec_request_size,
            'DNSSEC Response Size': dnssec_response_size
        })
