@echo off
echo ========================================
echo   Starting Kasir API Server
echo ========================================
echo.

REM Check if port 8080 is in use
echo Checking port 8080...
netstat -ano | findstr :8080 > temp_port.txt

if %errorlevel% equ 0 (
    echo Port 8080 is already in use. Stopping existing process...
    
    for /f "tokens=5" %%a in (temp_port.txt) do (
        echo Stopping process with PID: %%a
        taskkill /F /PID %%a 2>nul
    )
    
    echo Waiting for port to be released...
    timeout /t 2 /nobreak >nul
    echo.
)

del temp_port.txt 2>nul

echo Starting Kasir API Server...
echo.
echo ========================================
echo   Server is starting...
echo   Press Ctrl+C to stop the server
echo ========================================
echo.

go run main.go

