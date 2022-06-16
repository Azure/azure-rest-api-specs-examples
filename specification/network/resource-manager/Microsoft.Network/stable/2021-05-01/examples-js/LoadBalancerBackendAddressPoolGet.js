const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets load balancer backend address pool.
 *
 * @summary Gets load balancer backend address pool.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/LoadBalancerBackendAddressPoolGet.json
 */
async function loadBalancerBackendAddressPoolGet() {
  const subscriptionId = "subid";
  const resourceGroupName = "testrg";
  const loadBalancerName = "lb";
  const backendAddressPoolName = "backend";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.loadBalancerBackendAddressPools.get(
    resourceGroupName,
    loadBalancerName,
    backendAddressPoolName
  );
  console.log(result);
}

loadBalancerBackendAddressPoolGet().catch(console.error);
