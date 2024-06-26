from azure.identity import DefaultAzureCredential
from azure.mgmt.databricks import AzureDatabricksManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-databricks
# USAGE
    python workspace_virtual_network_peering_create_or_update.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = AzureDatabricksManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="subid",
    )

    response = client.vnet_peering.begin_create_or_update(
        resource_group_name="rg",
        workspace_name="myWorkspace",
        peering_name="vNetPeeringTest",
        virtual_network_peering_parameters={
            "properties": {
                "allowForwardedTraffic": False,
                "allowGatewayTransit": False,
                "allowVirtualNetworkAccess": True,
                "remoteVirtualNetwork": {
                    "id": "/subscriptions/0140911e-1040-48da-8bc9-b99fb3dd88a6/resourceGroups/subramantest/providers/Microsoft.Network/virtualNetworks/subramanvnet"
                },
                "useRemoteGateways": False,
            }
        },
    ).result()
    print(response)


# x-ms-original-file: specification/databricks/resource-manager/Microsoft.Databricks/stable/2023-02-01/examples/WorkspaceVirtualNetworkPeeringCreateOrUpdate.json
if __name__ == "__main__":
    main()
