@echo off
setlocal

REM Directory where gpm.exe should be placed
set "INSTALL_DIR=C:\Tools"

REM Move gpm.exe to the installation directory
if not exist "%INSTALL_DIR%" mkdir "%INSTALL_DIR%"
copy /y "%~dp0gpm.exe" "%INSTALL_DIR%"

REM Update PATH environment variable
echo %PATH% | findstr /i /c:"%INSTALL_DIR%" >nul
if %ERRORLEVEL% NEQ 0 (
    setx PATH "%PATH%;%INSTALL_DIR%"
    echo Directory added to PATH: %INSTALL_DIR%
) else (
    echo Directory is already in PATH: %INSTALL_DIR%
)

endlocal
