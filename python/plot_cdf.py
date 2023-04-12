import sys
import pandas as pd
import matplotlib.pyplot as plt
import numpy as np

def plot_cdf(csv_path):
    df = pd.read_csv(csv_path, sep=",")
    
    fig, ax = plt.subplots()
    ax.hist(df['response']-df['request'], bins=100, density=True, cumulative=-1, histtype='step', alpha=0.8, label="response - request sizes")
    ax.hist(df['response'], bins=100, density=True, cumulative=-1, histtype='step', alpha=0.8, label="response sizes")
    
    ax.annotate('holland.com', xy=(1000, 0.45), xytext=(1000, 0.80), arrowprops=dict(arrowstyle="->"), ha='center')
    ax.annotate('toetsingonline.nl', xy=(1500, 0.4), xytext=(1500, 0.75), arrowprops=dict(arrowstyle="->"), ha='center')
    ax.annotate('emissieautoriteit.nl', xy=(2000, 0.05), xytext=(2000, 0.40), arrowprops=dict(arrowstyle="->"), ha='center')
    ax.annotate('lijstvangevallenen.nl', xy=(3000, 0.01), xytext=(3000, 0.36), arrowprops=dict(arrowstyle="->"), ha='center')
    ax.annotate('politie.nl', xy=(4000, 0), xytext=(4000, 0.35), arrowprops=dict(arrowstyle="->"), ha='center')

    ax.grid(True)
    ax.legend(loc='right')
    ax.set_xlabel('Size')
    ax.set_ylabel('% of responses < size')

    plt.show()


if __name__ == "__main__":
    csv_path = sys.argv[1]
    plot_cdf(csv_path)

