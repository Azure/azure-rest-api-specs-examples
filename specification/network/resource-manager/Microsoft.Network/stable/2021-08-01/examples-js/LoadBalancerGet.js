const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the specified load balancer.
 *
 * @summary Gets the specified load balancer.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-08-01/examples/LoadBalancerGet.json
 */
async function getLoadBalancer() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const loadBalancerName = "lb";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.loadBalancers.get(resourceGroupName, loadBalancerName);
  console.log(result);
}

getLoadBalancer().catch(console.error);
