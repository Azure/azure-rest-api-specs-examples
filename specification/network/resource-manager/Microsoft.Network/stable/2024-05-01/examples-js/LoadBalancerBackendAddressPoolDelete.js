const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes the specified load balancer backend address pool.
 *
 * @summary Deletes the specified load balancer backend address pool.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2024-05-01/examples/LoadBalancerBackendAddressPoolDelete.json
 */
async function backendAddressPoolDelete() {
  const subscriptionId = process.env["NETWORK_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["NETWORK_RESOURCE_GROUP"] || "testrg";
  const loadBalancerName = "lb";
  const backendAddressPoolName = "backend";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.loadBalancerBackendAddressPools.beginDeleteAndWait(
    resourceGroupName,
    loadBalancerName,
    backendAddressPoolName,
  );
  console.log(result);
}
