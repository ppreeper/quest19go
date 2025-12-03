#! /usr/bin/env python3

import json
import requests

class JSON2:
    hostname: str
    port: int
    schema: str
    api_key: str

    def __init__(
        self,
        hostname: str = "localhost",
        port: int = 8069,
        schema: str = "http",
        api_key: str = "",
    ):
        self.hostname = hostname
        self.port = port
        self.schema = schema
        self.api_key = api_key

    def url(self, model: str, method: str) -> str:
        return f"{self.schema}://{self.hostname}:{self.port}/json/2/{model}/{method}"

    def Call(self, model: str, method: str, args: dict = {}):
        response = requests.post(
            self.url(model, method),
            headers={
                "Authorization": f"Bearer {self.api_key}",
                # "X-Odoo-Database": "...",
            },
            json=args,
        )
        # response.raise_for_status()
        data = response.json()
        return data

    def SearchRead(
        self,
        model: str,
        domain: list = [],
        fields: list = [],
        limit: int = 0,
        offset: int = 0,
        order: str = "",
    ):
        return self.Call(model,
            "search_read",
            {
                "domain": domain,
                "fields": fields,
                "limit": limit,
                "offset": offset,
                "order": order,
            }
        )

    def Count(
        self,
        model: str,
        domain: list = [],
        limit: int = 0,
    ):
        return self.Call(model,
            "search_count",
            {
            "domain": domain,
            "limit": limit,
            }
        )

    def GetID(
        self,
        model: str,
        domain: list = [],
        limit: int = 0,
        offset: int = 0,
        order: str = "",
    ):
        ids = self.Call(model,
            "search",
            {
                "domain": domain,
                "limit": limit,
                "offset": offset,
                "order": order,
            }
        )
        if ids and len(ids) > 0:
            return ids[0]
        return -1


    def Search(
        self,
        model: str,
        domain: list = [],
        limit: int = 0,
        offset: int = 0,
        order: str = "",
    ):
        return self.Call(model,
            "search",
            {
                "domain": domain,
                "limit": limit,
                "offset": offset,
                "order": order,
            }
        )


    def Read(
        self,
        model: str,
        fields: list = [],
        load: str = "_classic_read",
    ):
        return self.Call(model,
            "read",
            {
                "fields": fields,
                "load": load,
            }
        )


    def Create(
        self,
        model: str,
        vals_list: list = [],
    ):
        return self.Call(model,
            "create",
            {
                "vals_list": vals_list,
            }
        )

    def Write(
        self,
        model: str,
        ids: list = [],
        vals: dict = {},
    ):
        return self.Call(model,
            "write",
            {
                "ids": ids,
                "vals": vals,
            }
        )

    def Unlink(
        self,
        model: str,
        ids: list = [],
    ):
        return self.Call(model,
            "unlink",
            {
                "ids": ids,
            }
        )

    def FieldsGet(
        self,
        model: str,
        context: dict = {},
        allFields: str = None,
        attributes: str = None,
    ):
        return self.Call(model,
            "fields_get",
            {
                "context": context,
                "allFields": allFields,
                "attributes": attributes,
            }
        )


def prettyprint(object):
    return json.dumps(object, indent=4)
