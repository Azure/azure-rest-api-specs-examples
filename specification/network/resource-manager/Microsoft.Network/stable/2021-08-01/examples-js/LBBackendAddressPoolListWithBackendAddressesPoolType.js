const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets all the load balancer backed address pools.
 *
 * @summary Gets all the load balancer backed address pools.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-08-01/examples/LBBackendAddressPoolListWithBackendAddressesPoolType.json
 */
async function loadBalancerWithBackendAddressPoolContainingBackendAddresses() {
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

loadBalancerWithBackendAddressPoolContainingBackendAddresses().catch(console.error);
