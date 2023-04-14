import sys

import matplotlib.pyplot as plt
import pandas as pd


def main(csv_file, domain_output, ns_output):
    # Read the CSV file
    df = pd.read_csv(csv_file)

    # Drop rows with missing values in 'request' or 'response' columns
    df = df.dropna(subset=['request', 'response'])

    # Calculate amplification
    df['amplification'] = df['response'] / df['request']

    # Separate truncated and non-truncated data
    df_truncated = df[df['tc'] == 'tc']
    df_not_truncated = df[df['tc'] != 'tc']

    # Calculate amplification per domain and nameserver
    domain_amplification_truncated = df_truncated.groupby('domain')[
        'amplification'].mean()
    domain_amplification_not_truncated = df_not_truncated.groupby('domain')[
        'amplification'].mean()

    ns_amplification_truncated = df_truncated.groupby(
        'nameserver')['amplification'].mean()
    ns_amplification_not_truncated = df_not_truncated.groupby('nameserver')[
        'amplification'].mean()

    # Plot top and bottom 10 domains
    domain_amplification_sorted = domain_amplification_not_truncated.add(
        domain_amplification_truncated, fill_value=0).sort_values(ascending=False)
    top_domains = domain_amplification_sorted.head(10).index
    bottom_domains = domain_amplification_sorted.tail(10).index

    top_bottom_domains_truncated = domain_amplification_truncated.reindex(top_domains).fillna(
        0).append(domain_amplification_truncated.reindex(bottom_domains).fillna(0))
    top_bottom_domains_not_truncated = domain_amplification_not_truncated.reindex(top_domains).fillna(
        0).append(domain_amplification_not_truncated.reindex(bottom_domains).fillna(0))

    plt.figure(figsize=(10, 6))
    top_bottom_domains_truncated.plot(
        kind='bar', color='#377eb8', label='Truncated')
    top_bottom_domains_not_truncated.plot(
        kind='bar', bottom=top_bottom_domains_truncated, color='#ff7f00', label='Not Truncated')
    plt.xlabel('Domain')
    plt.ylabel('Amplification')
    plt.legend()
    plt.xticks(rotation=45, ha='right')
    plt.subplots_adjust(bottom=0.5)
    plt.savefig(domain_output)
    plt.clf()

    # Plot top and bottom 10 nameservers
    ns_amplification_sorted = ns_amplification_not_truncated.add(
        ns_amplification_truncated, fill_value=0).sort_values(ascending=False)
    top_ns = ns_amplification_sorted.head(10).index
    bottom_ns = ns_amplification_sorted.tail(10).index

    top_bottom_ns_truncated = ns_amplification_truncated.reindex(top_ns).fillna(
        0).append(ns_amplification_truncated.reindex(bottom_ns).fillna(0))
    top_bottom_ns_not_truncated = ns_amplification_not_truncated.reindex(top_ns).fillna(
        0).append(ns_amplification_not_truncated.reindex(bottom_ns).fillna(0))

    plt.figure(figsize=(10, 6))
    top_bottom_ns_truncated.plot(kind='bar', color='#377eb8', label='Truncated')
    top_bottom_ns_not_truncated.plot(
        kind='bar', bottom=top_bottom_ns_truncated, color='#ff7f00', label='Not Truncated')
    plt.xlabel('Nameserver')
    plt.ylabel('Amplification')
    plt.legend()
    plt.xticks(rotation=45, ha='right')
    plt.subplots_adjust(bottom=0.5)
    plt.savefig(ns_output)


if __name__ == '__main__':
    if len(sys.argv) != 4:
        print('Usage: python script.py <csv_file> <domain_output> <ns_output>')
        sys.exit(1)

    main(sys.argv[1], sys.argv[2], sys.argv[3])
