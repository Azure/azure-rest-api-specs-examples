from azure.identity import DefaultAzureCredential
from azure.mgmt.fluidrelay import FluidRelayManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-fluidrelay
# USAGE
    python fluid_relay_servers_update.py

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

    response = client.fluid_relay_servers.update(
        resource_group="myResourceGroup",
        fluid_relay_server_name="myFluidRelayServer",
        resource={"tags": {"Category": "sales"}},
    )
    print(response)


# x-ms-original-file: specification/fluidrelay/resource-manager/Microsoft.FluidRelay/stable/2022-06-01/examples/FluidRelayServers_Update.json
if __name__ == "__main__":
    main()
