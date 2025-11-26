const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Creates or updates a load balancer.
 *
 * @summary Creates or updates a load balancer.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/LoadBalancerCreateGatewayLoadBalancerConsumer.json
 */
async function createLoadBalancerWithGatewayLoadBalancerConsumerConfigured() {
  const subscriptionId = process.env["NETWORK_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["NETWORK_RESOURCE_GROUP"] || "rg1";
  const loadBalancerName = "lb";
  const parameters = {
    backendAddressPools: [{ name: "be-lb" }],
    frontendIPConfigurations: [
      {
        name: "fe-lb",
        gatewayLoadBalancer: {
          id: "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/loadBalancers/lb/frontendIPConfigurations/fe-lb-provider",
        },
        subnet: {
          id: "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/vnetlb/subnets/subnetlb",
        },
      },
    ],
    inboundNatPools: [],
    inboundNatRules: [
      {
        name: "in-nat-rule",
        backendPort: 3389,
        enableFloatingIP: true,
        frontendIPConfiguration: {
          id: "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/loadBalancers/lb/frontendIPConfigurations/fe-lb",
        },
        frontendPort: 3389,
        idleTimeoutInMinutes: 15,
        protocol: "Tcp",
      },
    ],
    loadBalancingRules: [
      {
        name: "rulelb",
        backendAddressPool: {
          id: "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/loadBalancers/lb/backendAddressPools/be-lb",
        },
        backendPort: 80,
        enableFloatingIP: true,
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
