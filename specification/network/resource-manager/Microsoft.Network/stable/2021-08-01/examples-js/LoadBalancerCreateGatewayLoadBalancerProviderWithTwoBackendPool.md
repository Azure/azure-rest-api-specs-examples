```javascript
const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates a load balancer.
 *
 * @summary Creates or updates a load balancer.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-08-01/examples/LoadBalancerCreateGatewayLoadBalancerProviderWithTwoBackendPool.json
 */
async function createLoadBalancerWithGatewayLoadBalancerProviderConfiguredWithTwoBackendPool() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const loadBalancerName = "lb";
  const parameters = {
    backendAddressPools: [{ name: "be-lb1" }, { name: "be-lb2" }],
    frontendIPConfigurations: [
      {
        name: "fe-lb",
        subnet: {
          id: "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/vnetlb/subnets/subnetlb",
        },
      },
    ],
    inboundNatPools: [],
    loadBalancingRules: [
      {
        name: "rulelb",
        backendAddressPool: {},
        backendAddressPools: [
          {
            id: "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/loadBalancers/lb/backendAddressPools/be-lb1",
          },
          {
            id: "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/loadBalancers/lb/backendAddressPools/be-lb2",
          },
        ],
        backendPort: 0,
        enableFloatingIP: true,
        frontendIPConfiguration: {
          id: "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/loadBalancers/lb/frontendIPConfigurations/fe-lb",
        },
        frontendPort: 0,
        idleTimeoutInMinutes: 15,
        loadDistribution: "Default",
        probe: {
          id: "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/loadBalancers/lb/probes/probe-lb",
        },
        protocol: "All",
      },
    ],
    location: "eastus",
    outboundRules: [],
    probes: [
      {
        name: "probe-lb",
        intervalInSeconds: 15,
        numberOfProbes: 2,
        port: 80,
        requestPath: "healthcheck.aspx",
        protocol: "Http",
      },
    ],
    sku: { name: "Premium" },
  };
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.loadBalancers.beginCreateOrUpdateAndWait(
    resourceGroupName,
    loadBalancerName,
    parameters
  );
  console.log(result);
}

createLoadBalancerWithGatewayLoadBalancerProviderConfiguredWithTwoBackendPool().catch(
  console.error
);
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-network_28.0.0/sdk/network/arm-network/README.md) on how to add the SDK to your project and authenticate.
