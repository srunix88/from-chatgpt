##
# How can I write a network server that is similar to netcat where it reads all data 
# from the port and just prints it out. 

import socket

# Create a TCP/IP socket
server_socket = socket.socket(socket.AF_INET, socket.SOCK_STREAM)

# Bind the socket to a specific address and port
server_address = ('localhost', 12345)
server_socket.bind(server_address)

# Listen for incoming connections
server_socket.listen(1)
print(f"Listening on {server_address}")

while True:
    # Wait for a connection
    connection, client_address = server_socket.accept()
    print(f"Accepted connection from {client_address}")

    try:
        # Receive data from the connection
        while True:
            data = connection.recv(1024)
            if data:
                # Print the received data to the console
                print(data.decode('utf-8'), end='')
            else:
                # No more data from the client
                break
    finally:
        # Clean up the connection
        connection.close()

