# HackingLab

# Goals
## Goal of the project
- Find the DNS servers of the Dutch government
- Check if ANY request is available on DNS servers
- Check for TXT records
- Find out if DNSSEC can be used for amplification
- Check for DNSSEC support
- Record request and response sizes
- Compute amplification potential
- Calculate cost of doing a DDOS of size x
- Look into the effect of rate limiting


## Optional goals
- Also include Belgian and Serbian government
- Try actual attack with KPN network
- Improve DNS compression estimates
- Set up a proof of concept on own infrastructure (TU Delft/VPS/home)

## Weekly goals
### Week 3
- List the failing DNS lookups
  - Figure out why this happens
  - Or take them out of the research
- Look into security of Azure / DNS ports
- Use wireshark/tshark as ground truth
- Do count headers in our amplification calculation 
  - Compare this with only calculating the payload like the original paper
- Come up with a strategy to nicely do a check for ratelimiting
- Start proof of concept of spoofing
  - Docker or real IPs
- Make a short midterm presentation

# Progress
## Week 2 (02-03-2023 - 09-03-2023)
- Bash script that does ns lookups
- Bash script that checks ANY size
- Analyzed the quirks of these scripts

## Week 1 (23-02-2023 - 02-03-2023)
- Formalized the goals of the project
- Experimented with some scripts
- Read some on the topic
