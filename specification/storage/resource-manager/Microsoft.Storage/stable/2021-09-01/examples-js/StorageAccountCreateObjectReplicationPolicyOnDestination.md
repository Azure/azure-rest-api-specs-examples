```javascript
const { StorageManagementClient } = require("@azure/arm-storage");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create or update the object replication policy of the storage account.
 *
 * @summary Create or update the object replication policy of the storage account.
 * x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2021-09-01/examples/StorageAccountCreateObjectReplicationPolicyOnDestination.json
 */
async function storageAccountCreateObjectReplicationPolicyOnDestination() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "res7687";
  const accountName = "dst112";
  const objectReplicationPolicyId = "default";
  const properties = {
    destinationAccount: "dst112",
    rules: [
      {
        destinationContainer: "dcont139",
        filters: { prefixMatch: ["blobA", "blobB"] },
        sourceContainer: "scont139",
      },
    ],
    sourceAccount: "src1122",
  };
  const credential = new DefaultAzureCredential();
  const client = new StorageManagementClient(credential, subscriptionId);
  const result = await client.objectReplicationPoliciesOperations.createOrUpdate(
    resourceGroupName,
    accountName,
    objectReplicationPolicyId,
    properties
  );
  console.log(result);
}

storageAccountCreateObjectReplicationPolicyOnDestination().catch(console.error);
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-storage_17.2.1/sdk/storage/arm-storage/README.md) on how to add the SDK to your project and authenticate.
