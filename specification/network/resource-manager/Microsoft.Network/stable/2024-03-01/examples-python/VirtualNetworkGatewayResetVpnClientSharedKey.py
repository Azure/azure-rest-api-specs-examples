from azure.identity import DefaultAzureCredential

from azure.mgmt.network import NetworkManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-network
# USAGE
    python virtual_network_gateway_reset_vpn_client_shared_key.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = NetworkManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="subid",
    )

    client.virtual_network_gateways.begin_reset_vpn_client_shared_key(
        resource_group_name="rg1",
        virtual_network_gateway_name="vpngw",
    ).result()


# x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2024-03-01/examples/VirtualNetworkGatewayResetVpnClientSharedKey.json
if __name__ == "__main__":
    main()
