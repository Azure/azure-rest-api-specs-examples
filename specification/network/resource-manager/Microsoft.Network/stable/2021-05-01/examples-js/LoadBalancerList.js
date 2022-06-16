const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets all the load balancers in a resource group.
 *
 * @summary Gets all the load balancers in a resource group.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/LoadBalancerList.json
 */
async function listLoadBalancersInResourceGroup() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.loadBalancers.list(resourceGroupName)) {
    resArray.push(item);
  }
  console.log(resArray);
}

listLoadBalancersInResourceGroup().catch(console.error);
