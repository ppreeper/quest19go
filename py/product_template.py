#!/usr/bin/env python3
from tqdm import tqdm

from jsonrpc import JSONRPC
from json2 import JSON2, prettyprint


# Products
def storable_products(source: JSONRPC, dest: JSON2):
    products, error = source.SearchRead(
        "product.product",
        fields=[
            "name",
            "barcode",
            "default_code",
            "list_price",
            "standard_price",
            "description",
            "description_picking",
            "description_pickingin",
            "description_pickingout",
            "description_sale",
            "description_purchase",
            "hs_code",
            "image_1920",
            "weight",
            "volume",
        ],
        domain=[["detailed_type", "=", "product"]],
    )
    if error:
        print("Error fetching products", error)
        return

    with tqdm(total=len(products)) as pbar:
        for product in products:
            # print(prettyprint(product))
            pid = dest.GetID(
                "product.template",
                domain=[
                    ["name", "=", product["name"]],
                    ["default_code", "=", product.get("default_code")],
                    ["barcode", "=", product.get("barcode")],
                ],
            )
            ur = {
                    "name": product["name"],
                    "default_code": product.get("default_code"),
                    "barcode": product.get("barcode"),
                    "list_price": product.get("list_price"),
                    "standard_price": product.get("standard_price"),
                    "description": product.get("description"),
                    "description_picking": product.get("description_picking"),
                    "description_pickingin": product.get("description_pickingin"),
                    "description_pickingout": product.get("description_pickingout"),
                    "description_sale": product.get("description_sale"),
                    "description_purchase": product.get("description_purchase"),
                    "hs_code": product.get("hs_code"),
                    "image_1920": product.get("image_1920"),
                    "weight": product.get("weight"),
                    "volume": product.get("volume"),
                    "type": "consu",
                    "is_storable": True,
                }
            if pid == -1:
                # print(prettyprint(product))
                # print(prettyprint(ur))
                ids = dest.Create("product.template", [ur])
                if ids == []:
                    print("Create error")
            # else:
            #     res = dest.Write("product.template", pid, product)
            #     if res is False:
            #         print("Write error")
            pbar.update(1)


def main():
    q15 = JSONRPC(
        hostname="quest15data.odoopro.ca",
        port=443,
        schema="https",
        database="quest15_data",
        username="aventador",
        password="57c850aeec0310c8c67bc167701c1145ecfff111",
    )
    q15.Login()
    q19 = JSON2(
        hostname="quest19data.odoopro.ca",
        port=443,
        schema="https",
        api_key="5ef6cf19b1c270ad4489095e9839b80aec61d32e",
    )

    storable_products(q15, q19)


if __name__ == "__main__":
    main()
