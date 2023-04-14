# IN4253ET "Hacking Lab"-Applied Security Analysis (2022/23 Q3)

## [Read our paper "When More Isn’t Merrier: Exploring DNS Amplification Hazards in the Dutch Digital Landscape"](results/paper.pdf)

## Goals
- Find the DNS servers of the Dutch government
- Check if ANY request is available on DNS servers
- Check for TXT records
- Find out if DNSSEC can be used for amplification
- Check for DNSSEC support
- Record request and response sizes
- Compute amplification potential
- Look into the effect of rate limiting
- Set up a proof of concept on own infrastructure (TU Delft/VPS/home)

## Folders
### Data
In the `input` folder you'll find a CSV file with the domain names owned by the Dutch government, this is an export from the [original DNS file](https://www.communicatierijk.nl/vakkennis/rijkswebsites/verplichte-richtlijnen/websiteregister-rijksoverheid) and the start of our entire research.

### Bash
This folder contains the scripts used to obtain the CSV files in the `results` folder. These are the scripts we relied upon for the results in the paper.

### Python
This folder contains initial scripts we used to analyze the data and experiment. We later realized that the results from our Python scripts were unreliable so we switched to using `dig` inside bash scripts. The only scripts we used for the paper are the plotting scripts.

### Go
This folder contains the code for the proof of concept attack that was implemented as part of our experiment.

### Results
This folder contains both CSV files as well as plots used in the paper. It might be interesting to look at the raw data and compare it a few years from the release of the paper.

