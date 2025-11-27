const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Creates or updates a load balancer.
 *
 * @summary Creates or updates a load balancer.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/LoadBalancerCreateWithInboundNatPool.json
 */
async function createLoadBalancerWithInboundNatPool() {
  const subscriptionId = process.env["NETWORK_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["NETWORK_RESOURCE_GROUP"] || "rg1";
  const loadBalancerName = "lb";
  const parameters = {
    backendAddressPools: [],
    frontendIPConfigurations: [
      {
        name: "test",
        id: "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/loadBalancers/lb/frontendIPConfigurations/test",
        privateIPAllocationMethod: "Dynamic",
        subnet: {
          id: "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/lbvnet/subnets/lbsubnet",
        },
        zones: [],
      },
    ],
    inboundNatPools: [
      {
        name: "test",
        backendPort: 8888,
        enableFloatingIP: true,
        enableTcpReset: true,
        frontendIPConfiguration: {
          id: "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/loadBalancers/lb/frontendIPConfigurations/test",
        },
        frontendPortRangeEnd: 8085,
        frontendPortRangeStart: 8080,
        id: "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/loadBalancers/lb/inboundNatPools/test",
        idleTimeoutInMinutes: 10,
        protocol: "Tcp",
      },
    ],
    inboundNatRules: [],
    loadBalancingRules: [],
    location: "eastus",
    outboundRules: [],
    probes: [],
    sku: { name: "Standard" },
  };
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.loadBalancers.beginCreateOrUpdateAndWait(
    resourceGroupName,
    loadBalancerName,
    parameters,
  );
  console.log(result);
}
