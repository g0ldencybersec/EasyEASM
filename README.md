# EasyEASM
EasyEASM repository for Black Hat Arsenal 2023

# Description
Easy EASM is just that... the easiest to set-up tool to give your organization visibility into its external facing assets.

The industry is dominated by $30k vendors selling "Attack Surface Management," but OG bug bounty hunters and red teamers know the truth. External ASM was born out of the bug bounty scene. Most of these $30k vendors use this open-source tooling on the backend.

With ten lines of setup or less, using open source tools, and one button deployment, Easy EASM will give your organization a complete view of your online assets. Easy EASM scans you daily and alerts you via Slack or Discord on newly found assets! Easy EASM also spits out an Excel skeleton for a Risk Register or Asset Database! This isn't rocket science.. but it's USEFUL. Don't get scammed. Grab Easy EASM and feel confident you know what's facing attackers on the internet.

# Installation
```sh
go install github.com/g0ldencybersec/EasyEASM/easyeasm@latest
```

# Example Config file
The tool will expect a configuration file named "config.yml" to be in the directory you are running from. An example of this yml file is below:
```yaml
# EasyEASM configurations
runConfig:
  domains:  # List root domains here.
    - example.com
    - mydomain.com
  slack: https://hooks.slack.com/services/DUMMYDATA/DUMMYDATA/RANDOM # Slack webhook url for slack notificaitions.
  discord: https://discord.com/api/webhooks/DUMMYURL/Dasdfsdf # Discord webhook for discord notifications.
  runType: fast   # Set to either fast (Passive enum) or complete (Active enumeration).
  activeWordList: subdomainWordlist.txt
  activeThreads: 100
```

# Running the tool
To run the tool, fill out the config file then simply run the easyeasm module:
```sh
$ ./easyeasm
```
After the run is complete you should see the output CSV (EasyEASM.csv) in the run directory. This can be added to your asset database and risk register!
