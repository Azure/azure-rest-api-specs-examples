const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets load balancer frontend IP configuration.
 *
 * @summary Gets load balancer frontend IP configuration.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2023-05-01/examples/LoadBalancerFrontendIPConfigurationGet.json
 */
async function loadBalancerFrontendIPConfigurationGet() {
  const subscriptionId = process.env["NETWORK_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["NETWORK_RESOURCE_GROUP"] || "testrg";
  const loadBalancerName = "lb";
  const frontendIPConfigurationName = "frontend";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.loadBalancerFrontendIPConfigurations.get(
    resourceGroupName,
    loadBalancerName,
    frontendIPConfigurationName
  );
  console.log(result);
}
