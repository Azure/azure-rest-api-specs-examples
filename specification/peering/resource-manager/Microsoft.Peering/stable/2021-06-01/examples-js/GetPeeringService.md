Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-peering_2.0.1/sdk/peering/arm-peering/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { PeeringManagementClient } = require("@azure/arm-peering");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets an existing peering service with the specified name under the given subscription and resource group.
 *
 * @summary Gets an existing peering service with the specified name under the given subscription and resource group.
 * x-ms-original-file: specification/peering/resource-manager/Microsoft.Peering/stable/2021-06-01/examples/GetPeeringService.json
 */
async function getAPeeringService() {
  const subscriptionId = "subId";
  const resourceGroupName = "rgName";
  const peeringServiceName = "peeringServiceName";
  const credential = new DefaultAzureCredential();
  const client = new PeeringManagementClient(credential, subscriptionId);
  const result = await client.peeringServices.get(resourceGroupName, peeringServiceName);
  console.log(result);
}

getAPeeringService().catch(console.error);
```
