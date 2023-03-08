# Read a CSV file that contains IPs of DNS servers
# Query each of the DNS servers with an ANY request
# Record the size of the outgoing DNS request and the response in bytes
# Save the results in a new CSV file, include the IP, the request size and the response size
import csv
import socket

# Function to send an ANY request to a DNS server and return the response size in bytes
def query_dns(ip):
    # Set up a UDP socket
    s = socket.socket(socket.AF_INET, socket.SOCK_DGRAM)
    s.settimeout(10)
    
    # Set up the DNS request packet
    message = bytearray()
    message += b'\x00\x01'    # ID
    message += b'\x00\x00'    # Flags
    message += b'\x00\x01'    # Question count
    message += b'\x00\x00'    # Answer count
    message += b'\x00\x00'    # Authority count
    message += b'\x00\x00'    # Additional count
    message += b'\x03\x77\x77\x77'    # Query name (www)
    message += b'\x06\x67\x6f\x6f\x67\x6c\x65'    # Query name (google)
    message += b'\x03\x63\x6f\x6d'    # Query name (com)
    message += b'\x00'    # End of query name
    message += b'\x00\x01'    # Query type (ANY)
    message += b'\x00\x01'    # Query class (IN)
    
    # Send the DNS request to the server
    request_size = s.sendto(message, (ip, 53))
    
    # Receive the DNS response from the server
    response, _ = s.recvfrom(65535)
    response_size = len(response)
    
    # Close the socket
    s.close()
    
    return request_size, response_size

# Open the input CSV file and read the IPs
with open('ips.csv', 'r') as infile:
    reader = csv.reader(infile)
    ips = [row[0] for row in reader]

# Query each DNS server and record the request and response sizes
results = []
for ip in ips:
    try:
        request_size, response_size = query_dns(ip)
        results.append([ip, request_size, response_size])
        print(f'Queried {ip}, request size = {request_size}, response size = {response_size}')
    except:
        print(f'Error querying {ip}')

# Open the output CSV file and write the results
with open('dns_results.csv', 'w', newline='') as outfile:
    writer = csv.writer(outfile)
    writer.writerow(['IP', 'Request Size', 'Response Size'])
    writer.writerows(results)
