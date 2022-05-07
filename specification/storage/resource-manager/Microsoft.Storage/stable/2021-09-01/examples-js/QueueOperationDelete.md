Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-storage_17.2.0/sdk/storage/arm-storage/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { StorageManagementClient } = require("@azure/arm-storage");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes the queue with the specified queue name, under the specified account if it exists.
 *
 * @summary Deletes the queue with the specified queue name, under the specified account if it exists.
 * x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2021-09-01/examples/QueueOperationDelete.json
 */
async function queueOperationDelete() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "res3376";
  const accountName = "sto328";
  const queueName = "queue6185";
  const credential = new DefaultAzureCredential();
  const client = new StorageManagementClient(credential, subscriptionId);
  const result = await client.queue.delete(resourceGroupName, accountName, queueName);
  console.log(result);
}

queueOperationDelete().catch(console.error);
```
