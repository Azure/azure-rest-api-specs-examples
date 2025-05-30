const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Creates or updates a load balancer.
 *
 * @summary Creates or updates a load balancer.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2024-07-01/examples/LoadBalancerCreateGatewayLoadBalancerProviderWithOneBackendPool.json
 */
async function createLoadBalancerWithGatewayLoadBalancerProviderConfiguredWithOneBackendPool() {
  const subscriptionId = process.env["NETWORK_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["NETWORK_RESOURCE_GROUP"] || "rg1";
  const loadBalancerName = "lb";
  const parameters = {
    backendAddressPools: [
      {
        name: "be-lb",
        tunnelInterfaces: [
          { type: "Internal", identifier: 900, port: 15000, protocol: "VXLAN" },
          { type: "Internal", identifier: 901, port: 15001, protocol: "VXLAN" },
        ],
      },
    ],
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
        backendAddressPools: [
          {
            id: "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/loadBalancers/lb/backendAddressPools/be-lb",
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
        probeThreshold: 1,
        requestPath: "healthcheck.aspx",
        protocol: "Http",
      },
    ],
    sku: { name: "Gateway" },
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
