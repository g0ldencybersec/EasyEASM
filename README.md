<div align="center">

![banner](banner.gif)

# EasyEASM

Zero-dollar attack surface management tool

featured at [Black Hat Arsenal 2023](https://www.blackhat.com/us-23/arsenal/schedule/index.html#easy-easm---the-zero-dollar-attack-surface-management-tool-33645) and [Recon Village @ DEF CON 2023](https://reconvillage.org/recon-village-talks-2023-defcon-31/).

![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)
<a href="https://github.com/g0ldencybersec/EasyEASM/blob/main/LICENSE.md">![MIT license](https://img.shields.io/badge/License-MIT-violet.svg?style=for-the-badge)</a>
![Discord](https://img.shields.io/badge/Discord-%235865F2.svg?style=for-the-badge&logo=discord&logoColor=white)
![Slack](https://img.shields.io/badge/Slack-4A154B?style=for-the-badge&logo=slack&logoColor=white)

</div>

## Description

Easy EASM is just that... the easiest to set-up tool to give your organization visibility into its external facing assets.

The industry is dominated by $30k vendors selling "Attack Surface Management," but OG bug bounty hunters and red teamers know the truth. External ASM was born out of the bug bounty scene. Most of these $30k vendors use this open-source tooling on the backend.

With ten lines of setup or less, using open-source tools, and one button deployment, Easy EASM will give your organization a complete view of your online assets. Easy EASM scans you daily and alerts you via Slack or Discord on newly found assets! Easy EASM also spits out an Excel skeleton for a Risk Register or Asset Database! This isn't rocket science, but it's USEFUL. Don't get scammed. Grab Easy EASM and feel confident you know what's facing attackers on the internet.

## Installation

```sh
go install github.com/g0ldencybersec/EasyEASM/easyeasm@latest
```

## Example config file

The tool expects a configuration file named `config.yml` to be in the directory you are running from.

Here is example of this yaml file:

```yaml
# EasyEASM configurations
runConfig:
  domains:  # List root domains here.
    - example.com
    - mydomain.com
  slack: https://hooks.slack.com/services/DUMMYDATA/DUMMYDATA/RANDOM # Slack webhook url for Slack notifications.
  discord: https://discord.com/api/webhooks/DUMMYURL/Dasdfsdf # Discord webhook for Discord notifications.
  runType: fast   # Set to either fast (passive enum) or complete (active enumeration).
  activeWordList: subdomainWordlist.txt
  activeThreads: 100
```

## Usage

To run the tool, fill out the config file: `config.yml`. Then, run the `easyeasm` module:

```sh
./easyeasm
```

After the run is complete, you should see the output CSV (`EasyEASM.csv`) in the run directory. This CSV can be added to your asset database and risk register!

## Warranty

The creator(s) of this tool provides no warranty or assurance regarding its performance, dependability, or suitability for any specific purpose.

The tool is furnished on an "as is" basis without any form of warranty, whether express or implied, encompassing, but not limited to, implied warranties of merchantability, fitness for a particular purpose, or non-infringement.

The user assumes full responsibility for employing this tool and does so at their own peril. The creator holds no accountability for any loss, damage, or expenses sustained by the user or any third party due to the utilization of this tool, whether in a direct or indirect manner.

Moreover, the creator(s) explicitly renounces any liability or responsibility for the accuracy, substance, or availability of information acquired through the use of this tool, as well as for any harm inflicted by viruses, malware, or other malicious components that may infiltrate the user's system as a result of employing this tool.

By utilizing this tool, the user acknowledges that they have perused and understood this warranty declaration and agree to undertake all risks linked to its utilization.

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) for details.

## Contact

For assistance, use the Issues tab. If we do not respond within 7 days, please reach out to us here.

- [Gunnar Andrews](https://twitter.com/G0LDEN_infosec)
- [Olivia Gallucci](https://oliviagallucci.com)
- [Jason Haddix](https://twitter.com/Jhaddix)
