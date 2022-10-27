const createNetworkManagementClient = require("@azure-rest/arm-network").default,
  { getLongRunningPoller } = require("@azure-rest/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv").config();

/**
 * This sample demonstrates how to Creates or updates a load balancer.
 *
 * @summary Creates or updates a load balancer.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-05-01/examples/LoadBalancerCreateGatewayLoadBalancerProviderWithTwoBackendPool.json
 */
async function createLoadBalancerWithGatewayLoadBalancerProviderConfiguredWithTwoBackendPool() {
  const credential = new DefaultAzureCredential();
  const client = createNetworkManagementClient(credential);
  const subscriptionId = "";
  const resourceGroupName = "rg1";
  const loadBalancerName = "lb";
  const options = {
    body: {
      location: "eastus",
      properties: {
        backendAddressPools: [
          { name: "be-lb1", properties: {} },
          { name: "be-lb2", properties: {} },
        ],
        frontendIPConfigurations: [
          {
            name: "fe-lb",
            properties: {
              subnet: {
                id: "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/vnetlb/subnets/subnetlb",
              },
            },
          },
        ],
        inboundNatPools: [],
        loadBalancingRules: [
          {
            name: "rulelb",
            properties: {
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
          },
        ],
        outboundRules: [],
        probes: [
          {
            name: "probe-lb",
            properties: {
              intervalInSeconds: 15,
              numberOfProbes: 2,
              port: 80,
              probeThreshold: 1,
              requestPath: "healthcheck.aspx",
              protocol: "Http",
            },
          },
        ],
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

createLoadBalancerWithGatewayLoadBalancerProviderConfiguredWithTwoBackendPool().catch(
  console.error
);
