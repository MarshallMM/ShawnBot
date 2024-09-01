cd C:\Users\marsh\ShawnBot\bot
setlocal EnableDelayedExpansion
set "command=env GOOS=linux GOARCH=arm GOARM=5 go build main.go"
"C:\Program Files\Git\git-bash.exe" -c "!command!"
scp main pi@192.168.1.77:/home/pi/ShawnBot/bot
rem scp dojo.txt pi@192.168.1.77:/home/pi/shawnBot
rem scp shawn.txt pi@192.168.1.77:/home/pi/shawnBot
pause
del main