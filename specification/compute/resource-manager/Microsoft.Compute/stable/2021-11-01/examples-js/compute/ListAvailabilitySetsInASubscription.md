Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-compute_17.3.1/sdk/compute/arm-compute/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists all availability sets in a subscription.
 *
 * @summary Lists all availability sets in a subscription.
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-11-01/examples/compute/ListAvailabilitySetsInASubscription.json
 */
async function listAvailabilitySetsInASubscription() {
  const subscriptionId = "{subscriptionId}";
  const expand = "virtualMachines$ref";
  const options = { expand };
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.availabilitySets.listBySubscription(options)) {
    resArray.push(item);
  }
  console.log(resArray);
}

listAvailabilitySetsInASubscription().catch(console.error);
```
