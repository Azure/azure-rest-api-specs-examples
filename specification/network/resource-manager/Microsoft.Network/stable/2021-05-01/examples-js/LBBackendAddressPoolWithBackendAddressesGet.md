Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-network_27.0.0/sdk/network/arm-network/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets load balancer backend address pool.
 *
 * @summary Gets load balancer backend address pool.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/LBBackendAddressPoolWithBackendAddressesGet.json
 */
async function loadBalancerWithBackendAddressPoolWithBackendAddresses() {
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

loadBalancerWithBackendAddressPoolWithBackendAddresses().catch(console.error);
```
