import sys
import pandas as pd
import matplotlib.pyplot as plt
import seaborn as sns

if len(sys.argv) < 4:
    print("Usage: python plot_amplification.py <csv_file> <output_domains_plot> <output_dns_plot>")
    sys.exit()

# Read the CSV file
csv_file = sys.argv[1]
output_domains_plot = sys.argv[2]
output_dns_plot = sys.argv[3]
df = pd.read_csv(csv_file, sep=',')

# Calculate amplification
df['amplification'] = df['response'] / df['request']

# Calculate average amplification by domain
avg_amplification_by_domain = df.groupby(
    'domain')['amplification'].mean().reset_index()

# Calculate average amplification by dns
avg_amplification_by_dns = df.groupby(
    'dns')['amplification'].mean().reset_index()

# Sort and get top 10 and bottom 10 amplifying domains
top_10_domains = avg_amplification_by_domain.nlargest(10, 'amplification')
bottom_10_domains = avg_amplification_by_domain.nsmallest(10, 'amplification')

# Sort and get top 10 and bottom 10 amplifying DNS
top_10_dns = avg_amplification_by_dns.nlargest(10, 'amplification')
bottom_10_dns = avg_amplification_by_dns.nsmallest(10, 'amplification')

# Merge top and bottom domains
merged_domains = pd.concat(
    [top_10_domains, bottom_10_domains]).reset_index(drop=True)
merged_domains['type'] = ['Top 10'] * 10 + ['Bottom 10'] * 10

# Merge top and bottom DNS
merged_dns = pd.concat([top_10_dns, bottom_10_dns]).reset_index(drop=True)
merged_dns['type'] = ['Top 10'] * 10 + ['Bottom 10'] * 10

# Plot top 10 and bottom 10 amplifying domains
plt.figure(figsize=(12, 6))
plt.title('Top 10 and Bottom 10 Amplifying Domains')
sns.barplot(x='domain', y='amplification', hue='type',
            data=merged_domains, palette='viridis')
plt.xticks(rotation=90)
plt.ylabel('Average Amplification')
plt.savefig(output_domains_plot, bbox_inches='tight')
plt.close()

# Plot top 10 and bottom 10 amplifying DNS
plt.figure(figsize=(12, 6))
plt.title('Top 10 and Bottom 10 Amplifying DNS')
sns.barplot(x='dns', y='amplification', hue='type',
            data=merged_dns, palette='viridis')
plt.xticks(rotation=90)
plt.ylabel('Average Amplification')
plt.savefig(output_dns_plot, bbox_inches='tight')
plt.close()
