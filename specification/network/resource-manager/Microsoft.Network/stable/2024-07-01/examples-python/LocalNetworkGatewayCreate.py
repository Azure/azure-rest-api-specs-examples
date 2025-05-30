from azure.identity import DefaultAzureCredential

from azure.mgmt.network import NetworkManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-network
# USAGE
    python local_network_gateway_create.py

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

    response = client.local_network_gateways.begin_create_or_update(
        resource_group_name="rg1",
        local_network_gateway_name="localgw",
        parameters={
            "location": "Central US",
            "properties": {
                "fqdn": "site1.contoso.com",
                "gatewayIpAddress": "11.12.13.14",
                "localNetworkAddressSpace": {"addressPrefixes": ["10.1.0.0/16"]},
            },
        },
    ).result()
    print(response)


# x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2024-07-01/examples/LocalNetworkGatewayCreate.json
if __name__ == "__main__":
    main()
