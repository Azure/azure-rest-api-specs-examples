const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to gets the specified load balancer.
 *
 * @summary gets the specified load balancer.
 * x-ms-original-file: 2025-07-01/LoadBalancerGet.json
 */
async function getLoadBalancer() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.loadBalancers.get("rg1", "lb");
  console.log(result);
}
