from azure.identity import DefaultAzureCredential
from azure.mgmt.fluidrelay import FluidRelayManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-fluidrelay
# USAGE
    python fluid_relay_servers_create_with_ami.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = FluidRelayManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="xxxx-xxxx-xxxx-xxxx",
    )

    response = client.fluid_relay_servers.create_or_update(
        resource_group="myResourceGroup",
        fluid_relay_server_name="myFluidRelayServer",
        resource={
            "identity": {
                "type": "SystemAssigned, UserAssigned",
                "userAssignedIdentities": {
                    "/subscriptions/xxxx-xxxx-xxxx-xxxx/resourceGroups/myResourceGroup/providers/Microsoft.ManagedIdentity/userAssignedIdentities/id1": {},
                    "/subscriptions/xxxx-xxxx-xxxx-xxxx/resourceGroups/myResourceGroup/providers/Microsoft.ManagedIdentity/userAssignedIdentities/id2": {},
                },
            },
            "location": "west-us",
            "properties": {"storagesku": "basic"},
            "tags": {"Category": "sales"},
        },
    )
    print(response)


# x-ms-original-file: specification/fluidrelay/resource-manager/Microsoft.FluidRelay/stable/2022-06-01/examples/FluidRelayServers_CreateWithAmi.json
if __name__ == "__main__":
    main()
