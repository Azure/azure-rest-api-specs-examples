const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes the specified load balancer.
 *
 * @summary Deletes the specified load balancer.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-08-01/examples/LoadBalancerDelete.json
 */
async function deleteLoadBalancer() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const loadBalancerName = "lb";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.loadBalancers.beginDeleteAndWait(resourceGroupName, loadBalancerName);
  console.log(result);
}

deleteLoadBalancer().catch(console.error);
