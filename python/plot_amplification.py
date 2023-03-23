import sys
import pandas as pd
import matplotlib.pyplot as plt
import numpy as np


def create_barcharts(csv_path, output_top10, output_bottom10):
    df = pd.read_csv(csv_path, sep=",")

    # Calculate the amplification factor
    df['amplification'] = df['response'] / df['request']

    # Group by domain and name, and compute the average amplification factor
    domain_avg = df.groupby(['domain', 'truncated']).mean().reset_index()
    name_avg = df.groupby(['name', 'truncated']).mean().reset_index()

    # Create top 10 and bottom 10 domain and name DataFrames with truncated information
    domain_agg = domain_avg.groupby(
        'domain')['amplification'].mean().reset_index()
    domain_agg = domain_avg.merge(
        domain_agg, on='domain', suffixes=('', '_mean'))

    name_agg = name_avg.groupby('name')['amplification'].mean().reset_index()
    name_agg = name_avg.merge(name_agg, on='name', suffixes=('', '_mean'))

    domain_top10 = domain_agg.nlargest(10, 'amplification_mean')
    domain_bottom10 = domain_agg.nsmallest(10, 'amplification_mean')
    name_top10 = name_agg.nlargest(10, 'amplification_mean')
    name_bottom10 = name_agg.nsmallest(10, 'amplification_mean')

    def plot_barchart(df, ylabel, title, output_path):
        fig, ax = plt.subplots()

        # Set width of bars
        bar_width = 0.25

        # Set positions of bars on the X-axis
        ind = np.arange(len(df[ylabel].unique()))
        bar_positions = [ind, ind + bar_width, ind + 2 * bar_width]

        # Create bars for different truncated categories
        for i, trunc in enumerate(['yes', 'no', None]):
            if trunc is None:
                trunc_df = df
            else:
                trunc_df = df[df['truncated'] == trunc]
            mean_values = trunc_df.groupby(ylabel)['amplification'].mean()
            label = 'Truncated Ignore' if trunc is None else f'Truncated {trunc.capitalize()}'

            # Align the data to have the same index as the X-axis ticks
            aligned_values = [mean_values.get(
                entry, 0) for entry in df[ylabel].unique()]

            ax.bar(bar_positions[i], aligned_values,
                   width=bar_width, label=label)

        # Set the labels, title, and legend
        ax.set_xticks(ind + bar_width)
        ax.set_xticklabels(df[ylabel].unique())
        ax.set_ylabel('Average Amplification')
        ax.set_title(title)
        ax.legend()

        # Save the plot
        plt.xticks(rotation=90)
        plt.savefig(output_path)
        plt.close()

    def plot_combined_barchart(df_top10, df_bottom10, ylabel, title, output_path):
        # Combine top 10 and bottom 10 DataFrames
        df_combined = pd.concat([df_top10, df_bottom10]).reset_index(drop=True)

        # Call the plot_barchart function
        plot_barchart(df_combined, ylabel, title, output_path)

    # Create the combined bar charts
    plot_combined_barchart(domain_top10, domain_bottom10, 'domain',
                           'Top and Bottom 10 Domains', output_domains_plot)
    plot_combined_barchart(name_top10, name_bottom10,
                           'name', 'Top and Bottom 10 Names', output_dns_plot)


if len(sys.argv) < 4:
    print("Usage: python plot_amplification.py <csv_file> <output_domains_plot> <output_dns_plot>")
    sys.exit()

# Read the CSV file
csv_file = sys.argv[1]
output_domains_plot = sys.argv[2]
output_dns_plot = sys.argv[3]

create_barcharts(csv_file, output_domains_plot, output_dns_plot)
