from azure.identity import DefaultAzureCredential

from azure.mgmt.network import NetworkManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-network
# USAGE
    python nat_gateway_create_or_update.py

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

    response = client.nat_gateways.begin_create_or_update(
        resource_group_name="rg1",
        nat_gateway_name="test-natgateway",
        parameters={
            "location": "westus",
            "properties": {
                "publicIpAddresses": [
                    {
                        "id": "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/publicIPAddresses/PublicIpAddress1"
                    }
                ],
                "publicIpPrefixes": [
                    {
                        "id": "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/publicIPPrefixes/PublicIpPrefix1"
                    }
                ],
            },
            "sku": {"name": "Standard"},
        },
    ).result()
    print(response)


# x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2024-05-01/examples/NatGatewayCreateOrUpdate.json
if __name__ == "__main__":
    main()
