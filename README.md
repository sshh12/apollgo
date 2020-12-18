# apollgo

> An android app for hosting a variety of proxy servers, proxy chains, and port forwarding (via [glider](https://github.com/nadoo/glider) and [hermes](https://github.com/sshh12/hermes)).

## Install

### Android

> Note: You may be able to download, extract, and install the APK all from the android device without needing a PC to set things up. This is untested.

1. Download and extract the [latest release](https://github.com/sshh12/apollgo/releases) on a PC.
2. Install [adb](https://developer.android.com/studio/command-line/adb)
3. Run `$ adb install apollgo.apk` with an android device connected
4. (Optional) While the app is running use `$ adb forward tcp:8888 tcp:8888` and you'll be able to access the app's settings by visiting [http://localhost:8888](http://localhost:8888)

## Use Cases

### Simple (SOCKS5, HTTP, HTTPS) Proxy Server

1. Launch the app and go to [http://localhost:8888](http://localhost:8888) in your device's browser
2. Go to the `glider` tab
3. Add a listener on `mixed://:1080`, then **Apply**, and restart the app.
4. Forward the proxy server to a PC using `$ adb forward tcp:1080 tcp:1080` (ensure device is plugged in via USB)
5. Update your PC's proxy settings to use `127.0.0.1:1080` for all traffic.

### Android Hotspot (Without Using Hotspot Mode 🔮)

1. Disable WiFi on your device as well as any auto-enable-WiFi settings
2. Launch the app and go to [http://localhost:8888](http://localhost:8888) in your device's browser
3. Go to the `glider` tab
4. Add `mixed://:1080`, then **Apply**, and restart the app.
5. Start adb as [`$ adb -a nodaemon server start`](https://stackoverflow.com/questions/56130335/adb-port-forwarding-to-listen-on-all-interfaces) (ensure device is plugged in via USB)
6. Forward the proxy server to a PC using `$ adb forward tcp:1080 tcp:1080`
7. Update your PC's proxy settings to use `127.0.0.1:1080` for all traffic.
8. (Optional) Connect your PC to a router (via WiFi or preferably ethernet) and have other devices use `<local PC ip>:1080` (eg `192.168.1.12:1080`) as a proxy server.

### Multi-Android Hotspot

Use several cellular-enabled devices as a single hotspot by load balancing proxy traffic between them.

1. Follow steps **1**-**5** from above on every device (every device should be connected via USB to the same PC).
2. Run `$ adb forward tcp:1080 tcp:1080`, `$ adb forward tcp:2080 tcp:1080`, ... for each device.
3. Download the latest [glider release](https://github.com/nadoo/glider/releases) (this will have to be done while the PC has an internet connection)
4. `$ glider -listen mixed://:1180 -forward socks5://127.0.0.1:1080 -forward socks5://127.0.0.1:2080 -checkwebsite www.google.com -checkinterval 300 -strategy rr -verbose` (include a `-forward socks5://127.0.0.1:<port>` for every connected device)
5. Update your PC's proxy settings to use `127.0.0.1:1180` for all traffic.
6. (Optional) Connect your PC to a router (via WiFi or preferably ethernet) and have other devices use `<local PC ip>:1180` (eg `192.168.1.12:1180`) as a proxy server.

### Port Forwarding An IP Camera App

1. Setup a [hermes server](https://github.com/sshh12/hermes) on eg DigitalOcean.
2. Install and start an [IP Camera](https://play.google.com/store/apps/details?id=com.pas.webcam) app, ensure it's running an HTTP webserver by visiting [http://localhost:8080](http://localhost:8080) (or whatever port its on)
3. Launch the apollgo and go to [http://localhost:8888](http://localhost:8888) in your device's browser
4. Go to the `glider` tab
5. Enable hermes and fill in fields to match your hermes server setup. Set forwards to `8080/8080` (forwarding local port 8080 to `<hermes server ip>:8080`). **Apply** and restart app.
6. Visit `<hermes server ip>:8080` anywhere to access the IP Camera stream

## Building

1. `$ cd react && yarn build:assets`
2. `$ gomobile build -target=android -o apollgo.apk ./cmd/android`

## Alternatives

If apollgo doesn't work for you, [Every Proxy](https://play.google.com/store/apps/details?id=com.gorillasoftware.everyproxy) is a closed-source alternative that works pretty well for SOCKS/HTTPS proxying.
