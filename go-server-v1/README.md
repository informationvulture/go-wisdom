# When the client vists '/'
-Server index.html with basic info

# When the client vists '/hello'
-Print the date

# When the client vists '/form'
-Get the weather


# Had to use:
```bash
$ sudo iptables -A PREROUTING -t nat -i eth0 -p tcp --dport 80 -j REDIRECT --to-port 8080
```

Also had to create a systemctl service to get the website running.
