from azure.identity import DefaultAzureCredential

from azure.mgmt.network import NetworkManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-network
# USAGE
    python network_interface_create_gateway_load_balancer_consumer.py

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

    response = client.network_interfaces.begin_create_or_update(
        resource_group_name="rg1",
        network_interface_name="test-nic",
        parameters={
            "location": "eastus",
            "properties": {
                "enableAcceleratedNetworking": True,
                "ipConfigurations": [
                    {
                        "name": "ipconfig1",
                        "properties": {
                            "gatewayLoadBalancer": {
                                "id": "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/loadBalancers/lb/frontendIPConfigurations/fe-lb-provider"
                            },
                            "publicIPAddress": {
                                "id": "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/publicIPAddresses/test-ip"
                            },
                            "subnet": {
                                "id": "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/rg1-vnet/subnets/default"
                            },
                        },
                    }
                ],
            },
        },
    ).result()
    print(response)


# x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2024-07-01/examples/NetworkInterfaceCreateGatewayLoadBalancerConsumer.json
if __name__ == "__main__":
    main()
