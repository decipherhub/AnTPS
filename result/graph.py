import matplotlib.pyplot as plt
import numpy as np
import glob
import os
import sys
import platform
import psutil
import re
from cpuinfo import get_cpu_info
from datetime import datetime
import matplotlib.transforms as transforms

html_template = """
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>AnTPS Report</title>
    <link rel="stylesheet" type="text/css" href="../styles/style.css"/>
</head>
<body>
    <h1>AnTPS(Anti-TPS): {title}</h1>
    <div class="description">
        <span class="description-title">Description:</span>
        <span class="description-content">benchmark the Blockchain TPS From Antps.
        <a href="https://github.com/rrhlrmrr/AnTPS/">Sourcecode</a> can be found in the Github repository.
        </span>
    </div>
        <div class="container">
            <h2>Test Environment & Network Architecture</h2>
                <div class="content">
                    <table class="table">
                        <tr>
                            <th>Parameter</th>
                            <th>Value</th>
                        </tr>
                        <tr>
                            <td>Network</td>
                            <td>{network}</td>
                        </tr>
                        <tr>
                            <td>OS</td>
                            <td>{os_info}</td>
                        </tr>
                        <tr>
                            <td>CPU</td>
                            <td>{cpu_info}</td>
                        </tr>
                        <tr>
                            <td>Memory</td>
                            <td>{memory_info}</td>
                        </tr>
                        <tr>
                            <td>TotalNode</td>
                            <td>2</td>
                        </tr>
                        <tr>
                            <td>Total Transaction</td>
                            <td>{total_transaction_count}</td>
                        </tr>
                        <tr>
                            <td>SendRate</td>
                            <td>{sendRate}</td>
                        </tr>
                    </table>
                <img src="../img/arch.png" alt="Network Architecture" class="image">
            </div>
        </div>
    <div class="container">
        <h2>Benchmark Results</h2>
        <div class="content">
            <table class="table">
                <tr>
                    <th>Parameter</th>
                    <th>Value</th>
                </tr>
                <tr>
                    <td>Theoretical TPS</td>
                    <td>{theoretical_tps}</td>
                <tr>
                    <td>Max Latency</td>
                    <td>{max_latency}</td>
                </tr>
                <tr>
                    <td>Max TPS</td>
                    <td>{max_tps}</td>
                </tr>
                <tr>
                    <td>Max Transaction Size</td>
                    <td>{max_transaction_size}</td>
                </tr>
            </table>
            <img src="{chain_type}_graph.png" alt="Benchmark Result" class="image">
        </div>
    </div>
</body>
</html>
"""

def extract_numbers(filename):
    match = re.search(r'\.(\d+)\.(\d+)\.', filename)
    if match:
        return int(match.group(1)), int(match.group(2))
    else:
        return None, None

def get_title(file_name):
    if "erc20" in file_name:
        contract_type = "erc20"
    elif "erc721" in file_name:
        contract_type = "erc721"
    elif "erc1155" in file_name:
        contract_type = "erc1155"
    elif "native" in file_name:
        contract_type = "Native"
    else:
        contract_type = "multi"


    if "mint" in file_name:
        title = "Mint"
    else:
        title = "Transfer"

    return {
        'contract_type': contract_type,
        'title': contract_type + title
    }


def get_system_info():
    os_info = f"{platform.system()} {platform.release()}"
    cpu_info = get_cpu_info()['brand_raw']
    memory = psutil.virtual_memory()
    memory_gb = round(memory.total / (1024**3), 2)

    return {
        'OS': os_info,
        'CPU': cpu_info,
        'Memory': f"{memory_gb}GB"
    }

def check_chain_type(chain):
    valid_chains = ['ava', 'eth', 'klay']

    if len(sys.argv) != 2 or sys.argv[1] not in valid_chains:
        print("Invalid Input Detected  Supported Blockchain Networks: ava/eth/klay")
        sys.exit(1)

    if chain == 'ava':
        network = 'avalanchego/1.11.8'
        theoretical_tps = 4500
    elif chain == 'eth':
        network = 'geth/1.14.6'
        theoretical_tps = 50
    elif chain == 'klay':
        network = 'klaytn/0.9.2'
        theoretical_tps = 4000

    return {
        'theoretical_tps': theoretical_tps,
        'network': network
    }

def generate_output(system_info):
    file_list = sorted(
        glob.glob(f'{chain_type}*.txt'),
        key=os.path.getmtime,
        reverse=False
    )


    if not file_list:
        raise FileNotFoundError('No files starting with "ava" found.')

    file_name = file_list[0]
    print(file_name)
    data = np.loadtxt(file_name, dtype=int)

    if len(data.shape) == 1:
        data = data.reshape(-1, 5)

    if data.size == 0:
        raise ValueError(f"No data found in {file_name}")

    block_height = data[:, 0]
    latency = data[:, 1]
    pending_txs = data[:, 2]
    confirmed_txs = data[:, 3]
    tps = data[:, 4]

    total_transaction = int(np.sum(confirmed_txs))

    min_block_height = min(block_height)
    max_tps_value = max(max(tps), max(confirmed_txs))
    fig, ax1 = plt.subplots()

    tps_color = 'dimgrey'
    ax1.set_xlabel('Block Number')
    ax1.set_ylabel('TPS', color=tps_color)
    bars1 = ax1.bar(block_height - 0.2, tps, width=0.4, color=tps_color, label='TPS')
    ax1.tick_params(axis='y', labelcolor=tps_color)
    ax1.set_ylim(0, max_tps_value + 50)

    theoretical_tps = check_chain_type(chain_type)['theoretical_tps']

    if theoretical_tps > max_tps_value + 50:
        y_max = ax1.get_ylim()[1]
        theoretical_tps_line = 0.98 * y_max
        break_pos = 0.94 * y_max 
        
        ax1.axhline(y=theoretical_tps_line, color='red', linestyle='--', label=f'Theoretical TPS')
        ax1.text(min_block_height - 0.7, theoretical_tps_line, theoretical_tps, color='red', verticalalignment='center', horizontalalignment='right')

        trans = transforms.blended_transform_factory(ax1.transAxes, ax1.transData)
        ax1.text(0, break_pos, '≈', transform=trans, ha='center', va='center', fontsize=18, color='black', clip_on=False)
    else:
        ax1.axhline(y=theoretical_tps, color='red', linestyle='--', label=f'Theoretical TPS')
        ax1.text(min_block_height - 0.7, theoretical_tps, theoretical_tps, color='red', verticalalignment='center', horizontalalignment='right')

    ax2 = ax1.twinx()
    confirmed_txs_color = 'silver'
    ax2.set_ylabel('Confirmed Txs per Block', color=confirmed_txs_color)
    bars2 = ax2.bar(block_height + 0.2, confirmed_txs, width=0.4, color=confirmed_txs_color, label='Confirmed Txs')
    ax2.tick_params(axis='y', labelcolor=confirmed_txs_color)
    ax2.set_ylim(0, max_tps_value + 50)

    ax1.legend(loc='upper left', bbox_to_anchor=(0, 1.2))
    ax2.legend(loc='upper right', bbox_to_anchor=(1, 1.143))

    max_latency = max(latency)
    max_tps = max(tps)
    max_transaction_size = max(confirmed_txs)

    result_dir = f"{chain_type}_{current_time}"

    os.makedirs(result_dir, exist_ok=True)

    output_file_name = os.path.join(result_dir, f'{chain_type}_graph.png')
    plt.savefig(output_file_name, bbox_inches='tight')
    plt.close()

    html_content = html_template.format(
        title=get_title(file_name)['title'].upper(),
        contract_type=get_title(file_name)['contract_type'].upper(),
        chain_type=chain_type,

        os_info=system_info['OS'],
        cpu_info=system_info['CPU'],
        memory_info=system_info['Memory'],
        network=check_chain_type(chain_type)['network'],
        total_transaction_count=extract_numbers(file_name)[0],
        sendRate=extract_numbers(file_name)[1],
        theoretical_tps=check_chain_type(chain_type)['theoretical_tps'],
        max_latency=max_latency,
        max_tps=max_tps,
        max_transaction_size=max_transaction_size,
    )

    html_file_name = os.path.join(result_dir, f'report_{chain_type}.html')
    with open(html_file_name, 'w') as html_file:
        html_file.write(html_content)
    print(f"Report saved as {html_file_name}")

if __name__ == "__main__":
    current_time = datetime.now().strftime("%Y%m%d_%H%M%S")
    chain_type = sys.argv[1]
    check_chain_type(chain_type)
    system_info = get_system_info()
    generate_output(system_info)