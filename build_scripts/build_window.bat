@echo off

set "projectPath=%~dp0..\"
set "binPath=%projectPath%\bin"
set "projectExeName=zimmer.exe"

if not exist "%binPath%" mkdir "%binPath%"

go build -o "%binPath%\%projectExeName%" "%projectPath%\cmd"

@echo build complete!

pause