const createNetworkManagementClient = require("@azure-rest/arm-network").default,
  { getLongRunningPoller } = require("@azure-rest/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv").config();

/**
 * This sample demonstrates how to Creates or updates a load balancer.
 *
 * @summary Creates or updates a load balancer.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-05-01/examples/LoadBalancerCreateWithInboundNatPool.json
 */
async function createLoadBalancerWithInboundNatPool() {
  const credential = new DefaultAzureCredential();
  const client = createNetworkManagementClient(credential);
  const subscriptionId = "";
  const resourceGroupName = "rg1";
  const loadBalancerName = "lb";
  const options = {
    body: {
      location: "eastus",
      properties: {
        backendAddressPools: [],
        frontendIPConfigurations: [
          {
            name: "test",
            id: "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/loadBalancers/lb/frontendIPConfigurations/test",
            properties: {
              privateIPAllocationMethod: "Dynamic",
              subnet: {
                id: "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/lbvnet/subnets/lbsubnet",
              },
            },
            zones: [],
          },
        ],
        inboundNatPools: [
          {
            name: "test",
            id: "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/loadBalancers/lb/inboundNatPools/test",
            properties: {
              backendPort: 8888,
              enableFloatingIP: true,
              enableTcpReset: true,
              frontendIPConfiguration: {
                id: "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/loadBalancers/lb/frontendIPConfigurations/test",
              },
              frontendPortRangeEnd: 8085,
              frontendPortRangeStart: 8080,
              idleTimeoutInMinutes: 10,
              protocol: "Tcp",
            },
          },
        ],
        inboundNatRules: [],
        loadBalancingRules: [],
        outboundRules: [],
        probes: [],
      },
      sku: { name: "Standard" },
    },
    queryParameters: { "api-version": "2022-05-01" },
  };
  const initialResponse = await client
    .path(
      "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/loadBalancers/{loadBalancerName}",
      subscriptionId,
      resourceGroupName,
      loadBalancerName
    )
    .put(options);
  const poller = getLongRunningPoller(client, initialResponse);
  const result = await poller.pollUntilDone();
  console.log(result);
}

createLoadBalancerWithInboundNatPool().catch(console.error);
