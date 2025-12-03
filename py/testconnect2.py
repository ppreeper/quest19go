#!/usr/bin/env python3

import json
from jsonrpc import JSONRPC
from json2 import JSON2, prettyprint


def main():
    # q15 = JSONRPC(
    #     hostname="quest15data.odoopro.ca",
    #     port=443,
    #     schema="https",
    #     database="quest15_data",
    #     username="aventador",
    #     password="57c850aeec0310c8c67bc167701c1145ecfff111",
    # )
    # q15.Login()

    q19 = JSON2(
        hostname="quest19data.odoopro.ca",
        port=443,
        schema="https",
        api_key="5ef6cf19b1c270ad4489095e9839b80aec61d32e",
    )

    # Create
    # Load
    # Count
    # count, error = q15.Count(
    #     "product.template", [["name", "=", "F:DOUBLEENDSTUD.10874-04647"]]
    # )
    # if error:
    #     print("Count error", error)
    # else:
    #     print("Count", count)

    count = q19.Count("product.template", [["name", "like", "F:DOUBLEEND"]])
    print(prettyprint(count))

    # FieldsGet
    fields = q19.FieldsGet(
        "product.template",
    )
    # if error:
    #     print("FieldsGet error", error)
    # else:
    #     print("FieldsGet", len(fields), json.dumps(fields, indent=4))

    # GetID
    # pid = q15.GetID("product.template", [["name", "=", "F:DOUBLEENDSTUD.10874-04647"]])
    # if pid == -1:
    #     print("GetID not found")
    # else:
    #     print("GetID", pid)

    # Search
    # productIDs, error = q15.Search("product.template")
    # if error:
    #     print("Search error", error)
    # else:
    #     print("Search", type(productIDs), len(productIDs))
    #     print(json.dumps(productIDs, indent=4))

    # productIDs = q19.Search("product.template")
    # print(dumps(productIDs, indent=4))

    # Read
    # pid = q15.GetID("product.template", [["name", "=", "F:DOUBLEENDSTUD.10874-04647"]])
    # product, error = q15.Read(
    #     "product.template", [pid], ["name", "list_price", "standard_price"]
    # )
    # if error:
    #     print("Read error", error)
    # else:
    #     print("Read", product)
    #     print("Read", product.get("list_price"))

    # SearchRead
    # products, error = q15.SearchRead(
    #     "product.template",
    #     fields=["name", "list_price", "standard_price"],
    # )
    # if error:
    #     print("SearchRead error", error)
    # else:
    #     print("SearchRead", len(products))
    #     print(prettyprint(products))

    # Unlink
    # pid = q19.GetID("product.template", [["name", "=", "PREEPER TEST PRODUCT X"]])
    # if pid == -1:
    #     print("Unlink not found")
    # else:
    #     print("Unlink", pid)
    #     u, e = q19.Unlink("product.template", [pid])
    #     if e:
    #         print("Unlink error", e)
    #     else:
    #         print("Unlink success", u)

    # Create
    # new_product = {
    #     "name": "PREEPER TEST PRODUCT X",
    #     "type": "consu",
    #     "list_price": 19.99,
    #     "standard_price": 9.99,
    #     "sale_ok": True,
    #     "purchase_ok": True,
    # }

    # c, e = q19.Create("product.template", new_product)
    # if e:
    #     print("Create error", e)
    # else:
    #     print("Create success", c)

    # Write
    # pid = q19.GetID("product.template", [["name", "=", "PREEPER TEST PRODUCT X"]])
    # if pid == -1:
    #     print("Write not found")
    # else:
    #     print("Write", pid)
    #     vals = {
    #         "list_price": 29.99,
    #         "standard_price": 14.99,
    #         "can_be_expensed": True,
    #     }
    #     w, e = q19.Write("product.template", pid, vals)
    #     if e:
    #         print("Write error", e)
    #     else:
    #         print("Write success", w)


if __name__ == "__main__":
    main()
