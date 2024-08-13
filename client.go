package main

import (
    "bufio"
    "fmt"
    "net"
    "os"
    "encoding/json"
    "log"
)

type Header struct {
    Type string `json:"type"`
    Tid uint32  `json:"tid"`
    Timestamp string `json:"timestamp"`
}


type Message struct {
    Header Header   `json:"header"`
    Body    interface{} `json:"body"`
}

type SampleMessage struct {
    Field1  string  `json:"field1"`
    Field2  int     `json:"field2"`
}

func main() {
    /*
    jsonData1 := `{
        "header"    : {
            "type"      : "request",
            "tid"       : 12345678,
            "message"   : "track_advisory"
        },
        "body"      : {
        }
    }`
    */
    var message Message
    msg := SampleMessage{"Hello", 123}
    jsonBytes, err := json.Marshal(msg)

    //err := json.Unmarshal([]byte(jsonData1), &message)
    if err != nil {
        log.Fatalf("Error : %v", err)
    }

    fmt.Printf("Message : %+v", message)

    if len(os.Args) == 1 {
        fmt.Println("Please provide host:port to connect to")
        os.Exit(1)
    }

    // Resolve the string address to a UDP address
    udpAddr, err := net.ResolveUDPAddr("udp", os.Args[1])

    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }

    // Dial to the address with UDP
    conn, err := net.DialUDP("udp", nil, udpAddr)

    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    for i:=0; i<4; i++ {
        // Send a message to the server
        // _, err = conn.Write([]byte(jsonData1))
        _, err = conn.Write(jsonBytes)
        fmt.Println("send...")
        if err != nil {
            fmt.Println(err)
            os.Exit(1)
        }

        // Read from the connection untill a new line is send
        data, err := bufio.NewReader(conn).ReadString('\n')
        if err != nil {
            fmt.Println(err)
            return
        }

        // Print the data read from the connection to the terminal
        fmt.Print("> ", string(data))
    }

}
