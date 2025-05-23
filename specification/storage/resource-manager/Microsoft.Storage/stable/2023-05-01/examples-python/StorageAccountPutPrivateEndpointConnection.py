from azure.identity import DefaultAzureCredential

from azure.mgmt.storage import StorageManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-storage
# USAGE
    python storage_account_put_private_endpoint_connection.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = StorageManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="{subscription-id}",
    )

    response = client.private_endpoint_connections.put(
        resource_group_name="res7687",
        account_name="sto9699",
        private_endpoint_connection_name="{privateEndpointConnectionName}",
        properties={
            "properties": {"privateLinkServiceConnectionState": {"description": "Auto-Approved", "status": "Approved"}}
        },
    )
    print(response)


# x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2023-05-01/examples/StorageAccountPutPrivateEndpointConnection.json
if __name__ == "__main__":
    main()
