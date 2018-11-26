@echo off
for %%i in (*.proto) do (
   echo gen %%~nxi...
   protoc.exe --go_out=../Server/src/server/msg  %%~nxi
   protoc.exe --go_out=../Server/src/center/msg  %%~nxi)

echo finish... 
pause