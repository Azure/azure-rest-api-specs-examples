from azure.identity import DefaultAzureCredential

from azure.mgmt.containerregistry import ContainerRegistryManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-containerregistry
# USAGE
    python connected_registry_update.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = ContainerRegistryManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="00000000-0000-0000-0000-000000000000",
    )

    response = client.connected_registries.begin_update(
        resource_group_name="myResourceGroup",
        registry_name="myRegistry",
        connected_registry_name="myScopeMap",
        connected_registry_update_parameters={
            "properties": {
                "clientTokenIds": [
                    "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.ContainerRegistry/registries/myRegistry/tokens/client1Token",
                    "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.ContainerRegistry/registries/myRegistry/tokens/client2Token",
                ],
                "garbageCollection": {"enabled": True, "schedule": "0 5 * * *"},
                "logging": {"auditLogStatus": "Enabled", "logLevel": "Debug"},
                "notificationsList": ["hello-world:*:*", "sample/repo/*:1.0:*"],
                "syncProperties": {"messageTtl": "P30D", "schedule": "0 0 */10 * *", "syncWindow": "P2D"},
            }
        },
    ).result()
    print(response)


# x-ms-original-file: specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/stable/2025-04-01/examples/ConnectedRegistryUpdate.json
if __name__ == "__main__":
    main()
