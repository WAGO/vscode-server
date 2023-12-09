# OPC UA

## Prerequisite

* Go installed on workstation
* Any Wago controller with OPC UA server
* WBM:
  * OPC UA Server enabled
  * Anonymous Access on
  * Port Authentication off

## About

This example could be runned directly on the device to achieve easy access to CoDeSys variables. It will read one variable cyclic from the CoDeSys program and output to the terminal.&#x20;

## Codesys

The program provide one variable to the 'symbol Configuration' with the OPC UA feature turned on. Add this to PLC\_PRG:

```
PROGRAM PLC_PRG
VAR
    iCnt : int;
END_VAR
```

Configure endless counter as OPC UA tag:

```
iCnt := iCnt + 1
```

## GO Client

We need to adapt the [code](opc-ua-client.go) to fit the actual device. IP address \<IP\_ADDRESS> and connection string \<DEVICE\_STRING> must be changed to actual device. Connection string could be found in WBM -> Fieldbus -> OPC UA tab.

```go
endpoint = flag.String("endpoint", "opc.tcp://<IP_ADDRESS>:4840", "OPC UA Endpoint URL")
nodeID   = flag.String("node", "ns=4;s=|var|<DEVICE_STRING>.Application.PLC_PRG.iCnt", "NodeID to read")
```

