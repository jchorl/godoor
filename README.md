# godoor
A simple webserver that opens our front door

# Running as a systemd service
## Installing
1) Copy the service file to the systemd service dir
`sudo cp ~/godoor/godoor.service /etc/systemd/system/godoor.service`
2) Enable the service
`sudo systemctl enable /etc/systemd/system/godoor.service`
3) Start the service
`sudo systemctl start godoor.service`

## Updating
If the systemd service file is updated:
1) Copy the updated file
`sudo cp ~/godoor/godoor.service /etc/systemd/system/godoor.service`
2) Reload the daemon
`sudo systemctl daemon-reload`
