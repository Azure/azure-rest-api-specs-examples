const createNetworkManagementClient = require("@azure-rest/arm-network").default,
  { getLongRunningPoller } = require("@azure-rest/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv").config();

/**
 * This sample demonstrates how to Creates or updates a load balancer.
 *
 * @summary Creates or updates a load balancer.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-05-01/examples/LoadBalancerCreateGlobalTier.json
 */
async function createLoadBalancerWithGlobalTierAndOneRegionalLoadBalancerInItsBackendPool() {
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
          {
            name: "be-lb",
            properties: {
              loadBalancerBackendAddresses: [
                {
                  name: "regional-lb1-address",
                  properties: {
                    loadBalancerFrontendIPConfiguration: {
                      id: "/subscriptions/subid/resourceGroups/regional-lb-rg1/providers/Microsoft.Network/loadBalancers/regional-lb/frontendIPConfigurations/fe-rlb",
                    },
                  },
                },
              ],
            },
          },
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
        loadBalancingRules: [
          {
            name: "rulelb",
            properties: {
              backendAddressPool: {
                id: "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/loadBalancers/lb/backendAddressPools/be-lb",
              },
              backendPort: 80,
              enableFloatingIP: false,
              frontendIPConfiguration: {
                id: "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/loadBalancers/lb/frontendIPConfigurations/fe-lb",
              },
              frontendPort: 80,
              idleTimeoutInMinutes: 15,
              loadDistribution: "Default",
              probe: {
                id: "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/loadBalancers/lb/probes/probe-lb",
              },
              protocol: "Tcp",
            },
          },
        ],
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
      sku: { name: "Standard", tier: "Global" },
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

createLoadBalancerWithGlobalTierAndOneRegionalLoadBalancerInItsBackendPool().catch(console.error);
