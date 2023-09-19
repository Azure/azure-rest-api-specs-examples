from azure.identity import DefaultAzureCredential
from azure.mgmt.network import NetworkManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-network
# USAGE
    python private_link_service_delete.py

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

    client.private_link_services.begin_delete(
        resource_group_name="rg1",
        service_name="testPls",
    ).result()


# x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2023-05-01/examples/PrivateLinkServiceDelete.json
if __name__ == "__main__":
    main()
