Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-storage_17.2.0/sdk/storage/arm-storage/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { StorageManagementClient } = require("@azure/arm-storage");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get the object replication policy of the storage account by policy ID.
 *
 * @summary Get the object replication policy of the storage account by policy ID.
 * x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2021-09-01/examples/StorageAccountGetObjectReplicationPolicy.json
 */
async function storageAccountGetObjectReplicationPolicies() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "res6977";
  const accountName = "sto2527";
  const objectReplicationPolicyId = "{objectReplicationPolicy-Id}";
  const credential = new DefaultAzureCredential();
  const client = new StorageManagementClient(credential, subscriptionId);
  const result = await client.objectReplicationPoliciesOperations.get(
    resourceGroupName,
    accountName,
    objectReplicationPolicyId
  );
  console.log(result);
}

storageAccountGetObjectReplicationPolicies().catch(console.error);
```
