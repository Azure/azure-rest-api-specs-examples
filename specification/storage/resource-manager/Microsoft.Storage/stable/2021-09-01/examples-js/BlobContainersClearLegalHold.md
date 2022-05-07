Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-storage_17.2.0/sdk/storage/arm-storage/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { StorageManagementClient } = require("@azure/arm-storage");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Clears legal hold tags. Clearing the same or non-existent tag results in an idempotent operation. ClearLegalHold clears out only the specified tags in the request.
 *
 * @summary Clears legal hold tags. Clearing the same or non-existent tag results in an idempotent operation. ClearLegalHold clears out only the specified tags in the request.
 * x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2021-09-01/examples/BlobContainersClearLegalHold.json
 */
async function clearLegalHoldContainers() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "res4303";
  const accountName = "sto7280";
  const containerName = "container8723";
  const legalHold = { tags: ["tag1", "tag2", "tag3"] };
  const credential = new DefaultAzureCredential();
  const client = new StorageManagementClient(credential, subscriptionId);
  const result = await client.blobContainers.clearLegalHold(
    resourceGroupName,
    accountName,
    containerName,
    legalHold
  );
  console.log(result);
}

clearLegalHoldContainers().catch(console.error);
```
