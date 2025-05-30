from azure.identity import DefaultAzureCredential

from azure.mgmt.network import NetworkManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-network
# USAGE
    python virtual_network_scalable_gateway_update.py

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

    response = client.virtual_network_gateways.begin_create_or_update(
        resource_group_name="rg1",
        virtual_network_gateway_name="ergw",
        parameters={
            "location": "centralus",
            "properties": {
                "activeActive": False,
                "allowRemoteVnetTraffic": False,
                "allowVirtualWanTraffic": False,
                "bgpSettings": None,
                "disableIPSecReplayProtection": False,
                "enableBgp": False,
                "enableBgpRouteTranslationForNat": False,
                "gatewayType": "ExpressRoute",
                "ipConfigurations": [
                    {
                        "name": "gwipconfig1",
                        "properties": {
                            "privateIPAllocationMethod": "Static",
                            "publicIPAddress": {
                                "id": "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/publicIPAddresses/gwpip"
                            },
                            "subnet": {
                                "id": "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/vnet1/subnets/GatewaySubnet"
                            },
                        },
                    }
                ],
                "natRules": [
                    {
                        "id": "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworkGateways/ergw/natRules/natRule1",
                        "name": "natRule1",
                        "properties": {
                            "externalMappings": [{"addressSpace": "50.0.0.0/24"}],
                            "internalMappings": [{"addressSpace": "10.10.0.0/24"}],
                            "ipConfigurationId": "",
                            "mode": "EgressSnat",
                            "type": "Static",
                        },
                    },
                    {
                        "id": "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworkGateways/ergw/natRules/natRule2",
                        "name": "natRule2",
                        "properties": {
                            "externalMappings": [{"addressSpace": "30.0.0.0/24"}],
                            "internalMappings": [{"addressSpace": "20.10.0.0/24"}],
                            "ipConfigurationId": "",
                            "mode": "IngressSnat",
                            "type": "Static",
                        },
                    },
                ],
                "sku": {"name": "ErGwScale", "tier": "ErGwScale"},
                "vpnClientConfiguration": None,
                "vpnType": "PolicyBased",
            },
        },
    ).result()
    print(response)


# x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2024-07-01/examples/VirtualNetworkScalableGatewayUpdate.json
if __name__ == "__main__":
    main()
