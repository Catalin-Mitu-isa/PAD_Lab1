import socket, threading

connections = []


def handle_connection(connection: socket.socket, addr: str) -> None:
    """
    Get connection in order to implement multiple socket connections
    :param connection: incoming socket connection
    :param addr: incoming socket connection address
    :return: None
    """
    while True:
        try:
            msg = connection.recv(1024)

            if msg:
                print(f"{addr[0]}:{addr[1]} - {msg.decode()}")

                msg_to_send = f"From ${addr[0]}:${addr[1]} - {msg.decode()}"
        except Exception as e:
            print(f"Error to handle connection: {e}")
            remove_connection(connection)
            break


def remove_connection(conn: socket.socket) -> None:
    """
    Remove specified connection from connections list
    :param conn: incoming socket connection address
    :return: None
    """
    if conn in connections:
        conn.close()
        connections.remove(conn)


def server() -> None:
    """
    Main process that receive client's connections and start a new thread
    to handle their messages
    :return: None
    """

    LISTENING_PORT = 12000

    try:
        socket_instance = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
        socket_instance.bind(('', LISTENING_PORT))
        socket_instance.listen()

        print("server running!")

        while True:
            socket_conn, addr = socket_instance.accept()
            connections.append(socket_conn)

            threading.Thread(target=handle_connection, args=[socket_conn, addr]).start()

    except Exception as e:
        print(f"An error has occurred when instancing socket: {e}")
    finally:
        if len(connections) > 0:
            for conn in connections:
                remove_connection(conn)

        socket_instance.close()


def start():
    server()
