# Websockets

## Prerequisite

* Go installed on workstation
* Wago controller&#x20;
* IoT library from CODESYS store

## About

A CODESYS websocket client communicating with a websocket server. Server must be started before CODESYS program.

## Codesys

```
PROGRAM PLC_PRG
VAR	
   FbWebSocket : WEB_SOCKET.WebSocketClient;
   sRx, sTx : STRING(1024);
   xReceived, xSend: BOOL;
   udiCountSend, uiCountRec: UDINT;  
END_VAR
```

Program:

```
FbWebSocket(xEnable:= TRUE, sUri:= 'ws://localhost:8080);

FbWebSocket.Read(pData:= ADR(sRx), udiSize:= SIZEOF(sRx), udiCount=> , xReceived=> xReceived, eFrameType=> , xIsFinalFragment=> );
IF xReceived THEN
   uiCountRec := uiCountRec + 1;
END_IF

IF xSend THEN
   sTx := 'hello from CODESYS';
   FbWebSocket.Write(pData:= ADR(sTx), udiSize:= SIZEOF(sTx), eFrameType:= , udiCount=> udiCountSend);
   IF udiCountSend > 0 THEN
	xSend := FALSE;
   END_IF
END_IF
```

## Websocket server&#x20;

Example listens for initial connection on localhost and http port 8080:

```go
http.ListenAndServe(":8080", nil)
```

Output from server when CODESYS client is connected:

```go
WEBSOCKET server: Client connected!
```

Output from server when recieving CODESYS message:&#x20;

```
WEBSOCKET server:  hello from CODESYS
```

Server will respond "WEBSOCKET server: Greetings!" back to CODESYS.
