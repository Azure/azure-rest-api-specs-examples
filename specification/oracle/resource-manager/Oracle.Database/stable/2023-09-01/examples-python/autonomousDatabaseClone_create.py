from typing import Any, IO, Union

from azure.identity import DefaultAzureCredential

from azure.mgmt.oracledatabase import OracleDatabaseMgmtClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-oracledatabase
# USAGE
    python autonomous_database_clone_create.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = OracleDatabaseMgmtClient(
        credential=DefaultAzureCredential(),
        subscription_id="00000000-0000-0000-0000-000000000000",
    )

    response = client.autonomous_databases.begin_create_or_update(
        resource_group_name="rg000",
        autonomousdatabasename="databasedb1",
        resource={
            "location": "eastus",
            "properties": {
                "adminPassword": "********",
                "characterSet": "AL32UTF8",
                "cloneType": "Full",
                "computeCount": 2,
                "computeModel": "ECPU",
                "dataBaseType": "Clone",
                "dataStorageSizeInTbs": 1,
                "displayName": "example_autonomous_databasedb1_clone",
                "ncharacterSet": "AL16UTF16",
                "sourceId": "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg000/providers/Oracle.Database/autonomousDatabases/databasedb1",
                "subnetId": "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg000/providers/Microsoft.Network/virtualNetworks/vnet1/subnets/subnet1",
                "vnetId": "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg000/providers/Microsoft.Network/virtualNetworks/vnet1",
            },
            "tags": {"tagK1": "tagV1"},
        },
    ).result()
    print(response)


# x-ms-original-file: specification/oracle/resource-manager/Oracle.Database/stable/2023-09-01/examples/autonomousDatabaseClone_create.json
if __name__ == "__main__":
    main()
