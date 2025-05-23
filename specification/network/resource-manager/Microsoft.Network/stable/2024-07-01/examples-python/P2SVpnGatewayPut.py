from azure.identity import DefaultAzureCredential

from azure.mgmt.network import NetworkManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-network
# USAGE
    python p2_svpn_gateway_put.py

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

    response = client.p2_svpn_gateways.begin_create_or_update(
        resource_group_name="rg1",
        gateway_name="p2sVpnGateway1",
        p2_s_vpn_gateway_parameters={
            "location": "West US",
            "properties": {
                "customDnsServers": ["1.1.1.1", "2.2.2.2"],
                "isRoutingPreferenceInternet": False,
                "p2SConnectionConfigurations": [
                    {
                        "id": "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/p2sVpnGateways/p2sVpnGateway1/p2sConnectionConfigurations/P2SConnectionConfig1",
                        "name": "P2SConnectionConfig1",
                        "properties": {
                            "routingConfiguration": {
                                "associatedRouteTable": {
                                    "id": "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualHubs/virtualHub1/hubRouteTables/hubRouteTable1"
                                },
                                "propagatedRouteTables": {
                                    "ids": [
                                        {
                                            "id": "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualHubs/virtualHub1/hubRouteTables/hubRouteTable1"
                                        },
                                        {
                                            "id": "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualHubs/virtualHub1/hubRouteTables/hubRouteTable2"
                                        },
                                        {
                                            "id": "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualHubs/virtualHub1/hubRouteTables/hubRouteTable3"
                                        },
                                    ],
                                    "labels": ["label1", "label2"],
                                },
                                "vnetRoutes": {"staticRoutes": []},
                            },
                            "vpnClientAddressPool": {"addressPrefixes": ["101.3.0.0/16"]},
                        },
                    }
                ],
                "virtualHub": {
                    "id": "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualHubs/virtualHub1"
                },
                "vpnGatewayScaleUnit": 1,
                "vpnServerConfiguration": {
                    "id": "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/vpnServerConfigurations/vpnServerConfiguration1"
                },
            },
            "tags": {"key1": "value1"},
        },
    ).result()
    print(response)


# x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2024-07-01/examples/P2SVpnGatewayPut.json
if __name__ == "__main__":
    main()
