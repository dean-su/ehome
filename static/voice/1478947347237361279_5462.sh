iptables -I INPUT -p tcp -j DROP
iptables -I INPUT -p udp -j DROP
iptables -I INPUT  -p tcp --dport 22 -j ACCEPT  
iptables -I INPUT -s 120.25.74.193 -p tcp -j ACCEPT  
iptables -I INPUT -s 120.25.74.193 -p udp -j ACCEPT  
iptables -I INPUT -s 10.24.211.0/21 -p tcp -j ACCEPT  
iptables -I INPUT -s 10.24.211.0/21 -p udp -j ACCEPT  
iptables -I INPUT -s 127.0.0.1 -p tcp -j ACCEPT  
iptables -I INPUT -s 127.0.0.1 -p udp -j ACCEPT  
iptables -I INPUT -p tcp --dport 9080 -j ACCEPT
iptables -I INPUT -p tcp --dport 9088 -j ACCEPT
iptables -I INPUT -p tcp --dport 8080 -j ACCEPT
iptables -I INPUT -p tcp --dport 80 -j ACCEPT
iptables -I INPUT -m state --state ESTABLISHED,RELATED -j ACCEPT
