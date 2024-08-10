import sys

from generated.server_pb2 import Credentials
from src.client import AuthClient


def main():
    if len(sys.argv) == 0:
        raise Exception("port must be provided")

    port = sys.argv[1]
    client = AuthClient('0.0.0.0:{}'.format(port))

    call = client.login(Credentials(login='kuzin57', password='kuzinr2003'))

    print(call)


if __name__ == "__main__":
    main()
