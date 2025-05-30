from azure.identity import DefaultAzureCredential

from azure.mgmt.network import NetworkManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-network
# USAGE
    python load_balancer_create_gateway_load_balancer_provider_with_two_backend_pool.py

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

    response = client.load_balancers.begin_create_or_update(
        resource_group_name="rg1",
        load_balancer_name="lb",
        parameters={
            "location": "eastus",
            "properties": {
                "backendAddressPools": [{"name": "be-lb1", "properties": {}}, {"name": "be-lb2", "properties": {}}],
                "frontendIPConfigurations": [
                    {
                        "name": "fe-lb",
                        "properties": {
                            "subnet": {
                                "id": "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/vnetlb/subnets/subnetlb"
                            }
                        },
                    }
                ],
                "inboundNatPools": [],
                "loadBalancingRules": [
                    {
                        "name": "rulelb",
                        "properties": {
                            "backendAddressPool": {},
                            "backendAddressPools": [
                                {
                                    "id": "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/loadBalancers/lb/backendAddressPools/be-lb1"
                                },
                                {
                                    "id": "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/loadBalancers/lb/backendAddressPools/be-lb2"
                                },
                            ],
                            "backendPort": 0,
                            "enableFloatingIP": True,
                            "frontendIPConfiguration": {
                                "id": "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/loadBalancers/lb/frontendIPConfigurations/fe-lb"
                            },
                            "frontendPort": 0,
                            "idleTimeoutInMinutes": 15,
                            "loadDistribution": "Default",
                            "probe": {
                                "id": "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/loadBalancers/lb/probes/probe-lb"
                            },
                            "protocol": "All",
                        },
                    }
                ],
                "outboundRules": [],
                "probes": [
                    {
                        "name": "probe-lb",
                        "properties": {
                            "intervalInSeconds": 15,
                            "numberOfProbes": 2,
                            "port": 80,
                            "probeThreshold": 1,
                            "protocol": "Http",
                            "requestPath": "healthcheck.aspx",
                        },
                    }
                ],
            },
            "sku": {"name": "Gateway"},
        },
    ).result()
    print(response)


# x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2024-07-01/examples/LoadBalancerCreateGatewayLoadBalancerProviderWithTwoBackendPool.json
if __name__ == "__main__":
    main()
