## networkTrafficCLI

This Go program is designed to capture and analyze network traffic on a specified network interface. It uses the [gopacket](https://github.com/google/gopacket) library to handle the capturing and decoding of network packets.

### Usage

To run the program, you'll need to have Go and the gopacket library installed on your system.

1. Clone the repository to your local machine
2. Change the target network interface in the code. (default is "en0")
3.  Make sure to run this program as administrator/root.
4. Run the program
The program will begin capturing packets and logging the source and destination ports, and protocol. To stop capturing, use `ctrl-c` in the terminal.

### Dependencies

- "github.com/google/gopacket"
- "github.com/google/gopacket/layers"
- "github.com/google/gopacket/pcap"

### Additional Information
The program opens a live connection to the network interface specified (in this case "en0", but this should be changed to the desired target network), and sets a BPF filter to capture only TCP packets with destination port 80 and source port range of 3000-3010. 
It then uses the PacketSource function to read the packets and pass them to the range packetSource.Packets().
For each packet, it extracts the TCP and IP layer, and logs the source and destination ports and protocol.

### Conclusion

This program is a simple and efficient tool for capturing and analyzing network traffic on a specified network interface using Go programming language. Feel free to customize it and use it for your own needs.
