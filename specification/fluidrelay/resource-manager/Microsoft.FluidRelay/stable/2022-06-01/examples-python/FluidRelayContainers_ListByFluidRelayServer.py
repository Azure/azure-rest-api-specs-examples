from azure.identity import DefaultAzureCredential
from azure.mgmt.fluidrelay import FluidRelayManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-fluidrelay
# USAGE
    python fluid_relay_containers_list_by_fluid_relay_server.py

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

    response = client.fluid_relay_containers.list_by_fluid_relay_servers(
        resource_group="myResourceGroup",
        fluid_relay_server_name="myFluidRelayServer",
    )
    for item in response:
        print(item)


# x-ms-original-file: specification/fluidrelay/resource-manager/Microsoft.FluidRelay/stable/2022-06-01/examples/FluidRelayContainers_ListByFluidRelayServer.json
if __name__ == "__main__":
    main()
