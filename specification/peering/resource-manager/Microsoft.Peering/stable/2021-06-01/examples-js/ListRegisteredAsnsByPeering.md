Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-peering_2.0.1/sdk/peering/arm-peering/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { PeeringManagementClient } = require("@azure/arm-peering");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists all registered ASNs under the given subscription, resource group and peering.
 *
 * @summary Lists all registered ASNs under the given subscription, resource group and peering.
 * x-ms-original-file: specification/peering/resource-manager/Microsoft.Peering/stable/2021-06-01/examples/ListRegisteredAsnsByPeering.json
 */
async function listAllTheRegisteredAsNsAssociatedWithThePeering() {
  const subscriptionId = "subId";
  const resourceGroupName = "rgName";
  const peeringName = "peeringName";
  const credential = new DefaultAzureCredential();
  const client = new PeeringManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.registeredAsns.listByPeering(resourceGroupName, peeringName)) {
    resArray.push(item);
  }
  console.log(resArray);
}

listAllTheRegisteredAsNsAssociatedWithThePeering().catch(console.error);
```
