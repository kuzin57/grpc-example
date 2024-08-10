import asyncio
import grpc
import logging

import generated.server_pb2 as server_py_pb2
import generated.server_pb2_grpc as server_py_pb2_grpc

logger = logging.getLogger(__name__)


class AuthClient:
    def __init__(self, server_addr: str):
        self.server_addr = server_addr

    
    def login(self, request: server_py_pb2.Credentials) -> server_py_pb2.Token:
        return asyncio.get_event_loop().run_until_complete(self._login(request))


    async def _login(self, request: server_py_pb2.Credentials) -> server_py_pb2.Token:
        response: server_py_pb2.Token

        async with grpc.aio.insecure_channel(self.server_addr) as channel:
            stub = server_py_pb2_grpc.AuthServiceStub(channel)
            response = await stub.Login(request)

        return response
