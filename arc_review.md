```
2024-07-24 13:25:46 | Info | Runner | Received done signal: context canceled
2024-07-24 13:25:46 | Info | Runner | Stopping all recurring tasks
2024-07-24 13:25:46 | Debug | Generators | Task stopped
2024-07-24 13:25:46 | Debug | USB | Task stopped
2024-07-24 13:25:47 | Debug | Sniffer | (Ethernet0) Sniffed 128 packets, 0.02 MB
2024-07-24 13:25:47 | Debug | Sniffer | Cannot send packets: Post "https://54.217.33.210/api/sync/send_packets": context canceled
upstream connection error, message buffered
2024-07-24 13:25:47 | Debug | Sniffer | Dropped 0.02 MB
2024-07-24 13:25:47 | Debug | Sniffer | (Ethernet0) Terminating routine
2024-07-24 13:25:47 | Debug | Sniffer | (Ethernet2) Sniffed 8 packets, 0.00 MB
2024-07-24 13:25:47 | Debug | Sniffer | Cannot send packets: message buffered, waiting for connection to recover
2024-07-24 13:25:47 | Debug | Sniffer | Dropped 0.00 MB
2024-07-24 13:25:47 | Debug | Sniffer | (Ethernet2) Terminating routine
2024-07-24 13:25:47 | Debug | Sniffer | (Ethernet1) Sniffed 566 packets, 0.11 MB
2024-07-24 13:25:47 | Debug | Sniffer | Cannot send packets: message buffered, waiting for connection to recover
2024-07-24 13:25:47 | Debug | Sniffer | Dropped 0.11 MB
2024-07-24 13:25:47 | Debug | Sniffer | (Ethernet1) Terminating routine
2024-07-24 13:25:47 | Debug | LiveCaptureContinuous | Done sniffing
2024-07-24 13:25:47 | Debug | LiveCaptureContinuous | Shutdown
2024-07-24 13:25:47 | Debug | Sniffer | Terminating routine
2024-07-24 13:25:47 | Debug | LiveCaptureContinuous | Terminating routine
2024-07-24 13:25:47 | Debug | Network sniffing | Task stopped
2024-07-24 13:25:48 | Debug | Slow Node Points | fetching network interfaces
2024-07-24 13:25:48 | Debug | Slow Node Points | fetching antivirus info
2024-07-24 13:25:48 | Debug | Slow Node Points | fetching CPU Hardware info
2024-07-24 13:25:48 | Debug | Slow Node Points | Slow Nodepoint production terminated
2024-07-24 13:25:48 | Debug | Slow Node Points | Task stopped
2024-07-24 13:25:50 | Debug | Discovery | Node 192.168.45.12 discovered through S7
2024-07-24 13:25:50 | Debug | Discovery | Node 192.168.45.232 discovered through S7
2024-07-24 13:25:50 | Debug | Discovery | Node 192.168.45.130 discovered through S7
2024-07-24 13:25:50 | Debug | Discovery | Node 192.168.45.227 discovered through S7
2024-07-24 13:25:50 | Debug | Discovery | Node 192.168.45.224 discovered through S7
2024-07-24 13:25:50 | Debug | Discovery | Node 192.168.45.215 discovered through S7
2024-07-24 13:25:50 | Info | Discovery | Got 6 targets from discovery
2024-07-24 13:25:50 | Info | Strategist | 6 responses retrieved in current Smart Polling cycle
2024-07-24 13:25:50 | Error | Strategist | error while sending additional rules: 'message buffered, waiting for connection to recover'
2024-07-24 13:25:50 | Error | Strategist | error while sending node executions: 'message buffered, waiting for connection to recover'
2024-07-24 13:25:50 | Debug | Discovery | Task stopped
2024-07-24 13:25:53 | Info | Smart Polling | Smart Polling cycle started
2024-07-24 13:25:53 | Error | Strategist | Could not fetch Smart Polling targets from upstream
2024-07-24 13:25:53 | Debug | Strategist | GetSmartPollingTargets failed with 'Get "https://54.217.33.210/api/sync/get_smartpolling_targets?sp_version=1": context canceled'
2024-07-24 13:25:53 | Info | Smart Polling | Got 0 targets from upstream
2024-07-24 13:25:53 | Debug | Smart Polling | Polling node 192.168.45.215 (ID: 192.168.45.215, properties {}) using strategy s7
2024-07-24 13:25:53 | Debug | Smart Polling | Polling node 192.168.45.130 (ID: 192.168.45.130, properties {}) using strategy s7
2024-07-24 13:25:53 | Debug | Smart Polling | Polling node 192.168.45.232 (ID: 192.168.45.232, properties {}) using strategy s7
2024-07-24 13:25:53 | Debug | Smart Polling | Polling node 192.168.45.12 (ID: 192.168.45.12, properties {}) using strategy s7
2024-07-24 13:25:53 | Debug | Smart Polling | Polling node 192.168.45.227 (ID: 192.168.45.227, properties {}) using strategy s7
2024-07-24 13:25:53 | Debug | Smart Polling | Polling node 192.168.45.224 (ID: 192.168.45.224, properties {}) using strategy s7
2024-07-24 13:25:58 | Debug | Smart Polling | s7 strategy error: could not connect to host 192.168.45.215:102: dial tcp 192.168.45.215:102: i/o timeout
2024-07-24 13:25:58 | Debug | Smart Polling | s7 strategy error: no information gathered
2024-07-24 13:26:03 | Debug | Smart Polling | s7 strategy error: could not connect to host 192.168.45.130:102: dial tcp 192.168.45.130:102: i/o timeout
2024-07-24 13:26:03 | Debug | Smart Polling | s7 strategy error: no information gathered
2024-07-24 13:26:08 | Debug | Smart Polling | s7 strategy error: could not connect to host 192.168.45.232:102: dial tcp 192.168.45.232:102: i/o timeout
2024-07-24 13:26:08 | Debug | Smart Polling | s7 strategy error: no information gathered
2024-07-24 13:26:13 | Debug | Smart Polling | s7 strategy error: could not connect to host 192.168.45.12:102: dial tcp 192.168.45.12:102: i/o timeout
2024-07-24 13:26:13 | Debug | Smart Polling | s7 strategy error: no information gathered
2024-07-24 13:26:18 | Debug | Smart Polling | s7 strategy error: could not connect to host 192.168.45.227:102: dial tcp 192.168.45.227:102: i/o timeout
2024-07-24 13:26:18 | Debug | Smart Polling | s7 strategy error: no information gathered
2024-07-24 13:26:23 | Debug | Smart Polling | s7 strategy error: could not connect to host 192.168.45.224:102: dial tcp 192.168.45.224:102: i/o timeout
2024-07-24 13:26:23 | Debug | Smart Polling | s7 strategy error: no information gathered
2024-07-24 13:26:23 | Info | Strategist | 6 responses retrieved in current Smart Polling cycle
2024-07-24 13:26:23 | Error | Strategist | error while sending node executions: 'Post "https://54.217.33.210/api/sync/send_sp_node_executions": context canceled
upstream connection error, message buffered'
2024-07-24 13:26:23 | Debug | Smart Polling | Task stopped
2024-07-24 13:26:23 | Info | Runner | Stopping resource monitor
2024-07-24 13:26:23 | Info | Runner | Waiting for all routines to terminate
2024-07-24 13:26:23 | Info | Runner | Save state
2024-07-24 13:26:23 | Debug | Runner | Terminating routine
2024-07-24 13:26:23 | Error | Strategist | error while sending additional rules: 'message buffered, waiting for connection to recover' (5 occurrences in the last 33.3997555s)
2024-07-24 13:26:23 | Error | Strategist | error while sending node executions: 'message buffered, waiting for connection to recover' (0 occurrences in the last 2562047h47m16.854775807s)
2024-07-24 13:26:23 | Error | Strategist | Could not fetch Smart Polling targets from upstream (0 occurrences in the last 2562047h47m16.854775807s)
2024-07-24 13:26:23 | Error | Strategist | error while sending node executions: 'Post "https://54.217.33.210/api/sync/send_sp_node_executions": context canceled
upstream connection error, message buffered' (0 occurrences in the last 2562047h47m16.854775807s)
2024-07-24 13:26:23 | Error | USB | USBPcap driver is not ready to be used, reboot required (0 occurrences in the last 2562047h47m16.854775807s)
```
