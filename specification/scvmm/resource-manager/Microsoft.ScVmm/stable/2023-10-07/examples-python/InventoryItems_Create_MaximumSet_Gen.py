from typing import Any, IO, Union

from azure.identity import DefaultAzureCredential

from azure.mgmt.scvmm import ScVmmMgmtClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-scvmm
# USAGE
    python inventory_items_create_maximum_set_gen.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = ScVmmMgmtClient(
        credential=DefaultAzureCredential(),
        subscription_id="79332E5A-630B-480F-A266-A941C015AB19",
    )

    response = client.inventory_items.create(
        resource_group_name="rgscvmm",
        vmm_server_name="O",
        inventory_item_resource_name="1BdDc2Ab-bDd9-Ebd6-bfdb-C0dbbdB5DEDf",
        resource={"kind": "M\\d_,V.", "properties": {"inventoryType": "InventoryItemProperties"}},
    )
    print(response)


# x-ms-original-file: specification/scvmm/resource-manager/Microsoft.ScVmm/stable/2023-10-07/examples/InventoryItems_Create_MaximumSet_Gen.json
if __name__ == "__main__":
    main()
