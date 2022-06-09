```javascript
const { PeeringManagementClient } = require("@azure/arm-peering");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes an existing registered prefix with the specified name under the given subscription, resource group and peering.
 *
 * @summary Deletes an existing registered prefix with the specified name under the given subscription, resource group and peering.
 * x-ms-original-file: specification/peering/resource-manager/Microsoft.Peering/stable/2021-06-01/examples/DeleteRegisteredPrefix.json
 */
async function deletesARegisteredPrefixAssociatedWithThePeering() {
  const subscriptionId = "subId";
  const resourceGroupName = "rgName";
  const peeringName = "peeringName";
  const registeredPrefixName = "registeredPrefixName";
  const credential = new DefaultAzureCredential();
  const client = new PeeringManagementClient(credential, subscriptionId);
  const result = await client.registeredPrefixes.delete(
    resourceGroupName,
    peeringName,
    registeredPrefixName
  );
  console.log(result);
}

deletesARegisteredPrefixAssociatedWithThePeering().catch(console.error);
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-peering_2.0.1/sdk/peering/arm-peering/README.md) on how to add the SDK to your project and authenticate.
