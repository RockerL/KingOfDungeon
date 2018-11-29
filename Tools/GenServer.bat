@echo off
for %%i in (*.proto) do (
   echo gen %%~nxi...
   protoc.exe --go_out=../Server/src/proto  %%~nxi)

echo finish... 
pause