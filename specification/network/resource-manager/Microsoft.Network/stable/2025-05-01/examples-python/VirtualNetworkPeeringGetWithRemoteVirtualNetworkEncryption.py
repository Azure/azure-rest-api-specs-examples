from azure.identity import DefaultAzureCredential

from azure.mgmt.network import NetworkManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-network
# USAGE
    python virtual_network_peering_get_with_remote_virtual_network_encryption.py

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

    response = client.virtual_network_peerings.get(
        resource_group_name="peerTest",
        virtual_network_name="vnet1",
        virtual_network_peering_name="peer",
    )
    print(response)


# x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/VirtualNetworkPeeringGetWithRemoteVirtualNetworkEncryption.json
if __name__ == "__main__":
    main()
