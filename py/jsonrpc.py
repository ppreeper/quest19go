#! /usr/bin/env python3
import json
import random

import requests


class JSONRPC:
    hostname: str
    port: int
    schema: str
    database: str
    username: str
    password: str
    url: str
    uid: int

    def __init__(
        self,
        hostname: str,
        port: int,
        schema: str,
        database: str,
        username: str,
        password: str,
    ):
        self.hostname = hostname
        self.port = port
        self.schema = schema
        self.database = database
        self.username = username
        self.password = password
        self.url = f"{self.schema}://{self.hostname}:{self.port}/jsonrpc"
        self.uid = 0

    def encodeClientRequest(self, service: str, method: str, args):
        return {
            "jsonrpc": "2.0",
            "method": "call",
            "id": random.randint(1, 1_000_000_000),
            "params": {
                "service": service,
                "method": method,
                "args": args,
            },
        }

    def Call(self, service: str, method: str, *args):
        req = self.encodeClientRequest(service, method, args)
        resp = requests.post(self.url, json=req).json()
        # print("resp", resp)
        return resp.get("result"), resp.get("error")

    def Login(self):
        self.uid, error = self.Call(
            "common", "login", self.database, self.username, self.password
        )
        if not self.uid:
            return -1
        return self.uid, error

    def Create(self, model: str, vals: dict):
        return self.Call(
            "object",
            "execute",
            self.database,
            self.uid,
            self.password,
            model,
            "create",
            [vals],
        )

    def Load(self, model: str, header: list, values: list):
        return self.Call(
            "object",
            "execute",
            self.database,
            self.uid,
            self.password,
            model,
            "load",
            [header, values],
        )

    def Count(self, model: str, domain: list = []):
        return self.Call(
            "object",
            "execute",
            self.database,
            self.uid,
            self.password,
            model,
            "search_count",
            domain,
        )

    def FieldsGet(self, model: str, fields: list = [], field_attributes: list = []):
        return self.Call(
            "object",
            "execute",
            self.database,
            self.uid,
            self.password,
            model,
            "fields_get",
            fields,
            field_attributes,
        )

    def GetID(self, model: str, domains: list = []):
        ids, _ = self.Call(
            "object",
            "execute",
            self.database,
            self.uid,
            self.password,
            model,
            "search",
            domains,
        )
        if ids and len(ids) > 0:
            return ids[0]
        return -1

    def Search(self, model: str, domain: list = []):
        return self.Call(
            "object",
            "execute",
            self.database,
            self.uid,
            self.password,
            model,
            "search",
            domain,
        )

    def Read(self, model: str, ids: list = [], fields: list = []):
        res, error = self.Call(
            "object",
            "execute",
            self.database,
            self.uid,
            self.password,
            model,
            "read",
            ids,
            fields,
        )
        if res and len(res) > 0:
            return res[0], error
        return {}, error

    def SearchRead(
        self,
        model: str,
        offset: int = 0,
        limit: int = 0,
        fields: list = [],
        domain: list = [],
    ):
        return self.Call(
            "object",
            "execute",
            self.database,
            self.uid,
            self.password,
            model,
            "search_read",
            domain,
            fields,
            offset,
            limit,
        )

    def Write(self, model: str, id: int, vals: dict = {}):
        return self.Call(
            "object",
            "execute",
            self.database,
            self.uid,
            self.password,
            model,
            "write",
            id,
            vals,
        )

    def Unlink(self, model: str, ids: list = []):
        return self.Call(
            "object",
            "execute",
            self.database,
            self.uid,
            self.password,
            model,
            "unlink",
            ids,
        )

    def Execute(self, model: str, method: str, *args):
        return self.Call(
            "object",
            "execute",
            self.database,
            self.uid,
            self.password,
            model,
            method,
            args,
        )

    def ExecuteKw(self, model: str, method: str, *args, **kwargs):
        return self.Call(
            "object",
            "execute_kw",
            self.database,
            self.uid,
            self.password,
            model,
            method,
            args,
            kwargs,
        )


def prettyprint(object):
    return json.dumps(object, indent=4)
