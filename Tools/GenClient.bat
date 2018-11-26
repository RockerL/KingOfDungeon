@echo off
for %%i in (*.proto) do (
   echo gen %%~nxi...
   bin\protoc.exe --csharp_out=../ClientLogic/Protos  %%~nxi)

echo finish... 
pause