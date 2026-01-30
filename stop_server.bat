@echo off
echo ========================================
echo   Stopping Kasir API Server
echo ========================================
echo.

echo Checking for processes using port 8080...
netstat -ano | findstr :8080 > temp_port.txt

if %errorlevel% equ 0 (
    echo Found process using port 8080
    echo.
    
    for /f "tokens=5" %%a in (temp_port.txt) do (
        echo Stopping process with PID: %%a
        taskkill /F /PID %%a 2>nul
    )
    
    echo.
    echo Server stopped successfully!
) else (
    echo No process found using port 8080
)

del temp_port.txt 2>nul

echo.
echo ========================================
echo Done!

