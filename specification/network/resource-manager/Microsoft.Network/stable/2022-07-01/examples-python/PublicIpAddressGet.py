from azure.identity import DefaultAzureCredential
from azure.mgmt.network import NetworkManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-network
# USAGE
    python public_ip_address_get.py

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

    response = client.public_ip_addresses.get(
        resource_group_name="rg1",
        public_ip_address_name="testDNS-ip",
    )
    print(response)


# x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-07-01/examples/PublicIpAddressGet.json
if __name__ == "__main__":
    main()
