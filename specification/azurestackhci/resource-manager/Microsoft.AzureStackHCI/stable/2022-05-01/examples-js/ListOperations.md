Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-azurestackhci_3.0.0/sdk/azurestackhci/arm-azurestackhci/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { AzureStackHCIClient } = require("@azure/arm-azurestackhci");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List all available Microsoft.AzureStackHCI provider operations
 *
 * @summary List all available Microsoft.AzureStackHCI provider operations
 * x-ms-original-file: specification/azurestackhci/resource-manager/Microsoft.AzureStackHCI/stable/2022-05-01/examples/ListOperations.json
 */
async function createCluster() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const credential = new DefaultAzureCredential();
  const client = new AzureStackHCIClient(credential, subscriptionId);
  const result = await client.operations.list();
  console.log(result);
}

createCluster().catch(console.error);
```
