const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets all the load balancer frontend IP configurations.
 *
 * @summary Gets all the load balancer frontend IP configurations.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-07-01/examples/LoadBalancerFrontendIPConfigurationList.json
 */
async function loadBalancerFrontendIPConfigurationList() {
  const subscriptionId = "subid";
  const resourceGroupName = "testrg";
  const loadBalancerName = "lb";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.loadBalancerFrontendIPConfigurations.list(
    resourceGroupName,
    loadBalancerName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

loadBalancerFrontendIPConfigurationList().catch(console.error);
