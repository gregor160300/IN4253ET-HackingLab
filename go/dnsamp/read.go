package dnsamp

import(
    "os"
    "encoding/csv"
)

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

