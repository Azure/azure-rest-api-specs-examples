from azure.identity import DefaultAzureCredential

from azure.mgmt.network import NetworkManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-network
# USAGE
    python private_link_service_delete_private_endpoint_connection.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = NetworkManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="subId",
    )

    client.private_link_services.begin_delete_private_endpoint_connection(
        resource_group_name="rg1",
        service_name="testPls",
        pe_connection_name="testPlePeConnection",
    ).result()


# x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2024-07-01/examples/PrivateLinkServiceDeletePrivateEndpointConnection.json
if __name__ == "__main__":
    main()
