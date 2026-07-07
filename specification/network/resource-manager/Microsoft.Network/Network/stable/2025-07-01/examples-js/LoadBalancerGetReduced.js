const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to gets the specified load balancer.
 *
 * @summary gets the specified load balancer.
 * x-ms-original-file: 2025-07-01/LoadBalancerGetReduced.json
 */
async function getLoadBalancerReduced() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.loadBalancers.get("rg-name", "lb-name", { detailLevel: "Reduced" });
  console.log(result);
}
