from azure.identity import DefaultAzureCredential

from azure.mgmt.servicelinker import ServiceLinkerManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-servicelinker
# USAGE
    python put_connector.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = ServiceLinkerManagementClient(
        credential=DefaultAzureCredential(),
    )

    response = client.connector.begin_create_or_update(
        subscription_id="00000000-0000-0000-0000-000000000000",
        resource_group_name="test-rg",
        location="westus",
        connector_name="connectorName",
        parameters={
            "properties": {
                "authInfo": {"authType": "secret"},
                "secretStore": {
                    "keyVaultId": "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/test-rg/providers/Microsoft.KeyVault/vaults/test-kv"
                },
                "targetService": {
                    "id": "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/test-rg/providers/Microsoft.DocumentDb/databaseAccounts/test-acc/mongodbDatabases/test-db",
                    "type": "AzureResource",
                },
            }
        },
    ).result()
    print(response)


# x-ms-original-file: specification/servicelinker/resource-manager/Microsoft.ServiceLinker/preview/2024-07-01-preview/examples/PutConnector.json
if __name__ == "__main__":
    main()
