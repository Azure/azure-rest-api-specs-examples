const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets all the load balancer backed address pools.
 *
 * @summary Gets all the load balancer backed address pools.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-07-01/examples/LoadBalancerBackendAddressPoolList.json
 */
async function loadBalancerBackendAddressPoolList() {
  const subscriptionId = "subid";
  const resourceGroupName = "testrg";
  const loadBalancerName = "lb";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.loadBalancerBackendAddressPools.list(
    resourceGroupName,
    loadBalancerName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

loadBalancerBackendAddressPoolList().catch(console.error);
