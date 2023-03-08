
<h2 align="center">portfwd</h2>
<p align="center">A simple and efficient TCP/UDP port forwarder</p>
<p align="center">
    <a href="#about">About</a> •
    <a href="#features">Features</a> •
    <a href="#build">Build</a> •
    <a href="#configuration">Configuration</a> •
    <a href="#run">Run</a> •
    <a href="#license">License</a>
</p>

## About

portfwd is a simple tool designed to forward TCP or UDP traffic to different ports.

## Features

- ✅ A **single** statically compiled **binary** for each OS/architecture
- ✅ Support **TCP traffic**
- ✅ Support **UDP traffic**
- ✅ Can forward **multiple ports** to **multiple locations** simultaneously
- ✅ An **easily configurable** tool
- ✅ Can operate effortlessly as a **daemon**

## Build

First check that you have **Golang** installed on your machine.  
Then, **run**:  
```bash
make 
```
Quite simply!

## Configuration

Here is an example of a configuration (YAML): 
```yaml
forward:
  # Forward TCP traffic from one port to a single other
  - protocol: tcp
    from: 127.0.0.1:5100
    to: [10.2.0.1:5101]
  # Forward of TCP traffic from one port to several others
  - protocol: tcp
    from: 127.0.0.1:5200
    to: [127.0.0.1:5201, 10.2.0.1:5202]
  # Forward UDP traffic from one port to a single other
  - protocol: udp
    from: 127.0.0.1:5300
    to: [10.2.0.1:5301]
  # Forward of UDP traffic from one port to several others
  - protocol: udp
    from: 127.0.0.1:5400
    to: [127.0.0.1:5401, 10.2.0.1:5402]
```

## Run
### Direct use

To be able to use portfwd directly, you must set the environment variable `PORTFWD_CONFIG_FILE_PATH` as the path to the configuration file.  
For example :
```
export PORTFWD_CONFIG_FILE_PATH=./config.yaml
```
Then you can run the program:
```
./portfwd
```

### As a systemd service

It is possible to easily use portfwd as a daemon with the provided systemd service.  
The systemd service provided can be adapted to your needs.

Steps to install portfwd as a systemd service:

- Create a group:

```
groupadd portfwd
```

 - Create an user:

```
useradd -r -s /sbin/nologin -g portfwd portfwd
```

- Copy the `portfwd` binary to `/usr/bin/`

```
cp portfwd /usr/bin/
```

- Create a `portfwd` folder in `/etc/` for the configuration file

```
mkdir /etc/portfwd
```

- Copy the `config.yaml` configuration file to `/etc/portfwd/`

```
cp config.yaml /etc/portfwd/
```

- Copy the `portfwd.service` systemd service file to `/etc/systemd/system/`

```
cp portfwd.service /etc/systemd/system/
```

- Start the systemd service

```
systemctl start portfwd
```

## License

This program is free software: you can redistribute it and/or modify it under the terms of the GNU General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.

This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU General Public License for more details.

You should have received a copy of the GNU General Public License along with this program. If not, see http://www.gnu.org/licenses/.
