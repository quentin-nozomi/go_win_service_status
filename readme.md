```
GOOS=windows go build -o testServ.exe
```

```
C:\Users\Administrator\Desktop\tmp>testServ.exe
current Arc status
Running
requesting stop
returned status
StopPending

C:\Users\Administrator\Desktop\tmp>sc query Arc

SERVICE_NAME: Arc
TYPE               : 10  WIN32_OWN_PROCESS
STATE              : 1  STOPPED
WIN32_EXIT_CODE    : 0  (0x0)
SERVICE_EXIT_CODE  : 0  (0x0)
CHECKPOINT         : 0x0
WAIT_HINT          : 0x0
```