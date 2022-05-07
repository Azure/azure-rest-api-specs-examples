Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-storage_17.2.0/sdk/storage/arm-storage/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { StorageManagementClient } = require("@azure/arm-storage");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Sets the ImmutabilityPolicy to Locked state. The only action allowed on a Locked policy is ExtendImmutabilityPolicy action. ETag in If-Match is required for this operation.
 *
 * @summary Sets the ImmutabilityPolicy to Locked state. The only action allowed on a Locked policy is ExtendImmutabilityPolicy action. ETag in If-Match is required for this operation.
 * x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2021-09-01/examples/BlobContainersLockImmutabilityPolicy.json
 */
async function lockImmutabilityPolicy() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "res2702";
  const accountName = "sto5009";
  const containerName = "container1631";
  const ifMatch = '"8d59f825b721dd3"';
  const credential = new DefaultAzureCredential();
  const client = new StorageManagementClient(credential, subscriptionId);
  const result = await client.blobContainers.lockImmutabilityPolicy(
    resourceGroupName,
    accountName,
    containerName,
    ifMatch
  );
  console.log(result);
}

lockImmutabilityPolicy().catch(console.error);
```
