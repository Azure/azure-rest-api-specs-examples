from azure.identity import DefaultAzureCredential

from azure.mgmt.cosmosdb import CosmosDBManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-cosmosdb
# USAGE
    python cosmos_db_private_endpoint_connection_update.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = CosmosDBManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="00000000-1111-2222-3333-444444444444",
    )

    response = client.private_endpoint_connections.begin_create_or_update(
        resource_group_name="rg1",
        account_name="ddb1",
        private_endpoint_connection_name="privateEndpointConnectionName",
        parameters={
            "properties": {
                "privateLinkServiceConnectionState": {
                    "description": "Approved by johndoe@contoso.com",
                    "status": "Approved",
                }
            }
        },
    ).result()
    print(response)


# x-ms-original-file: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/preview/2024-12-01-preview/examples/CosmosDBPrivateEndpointConnectionUpdate.json
if __name__ == "__main__":
    main()
