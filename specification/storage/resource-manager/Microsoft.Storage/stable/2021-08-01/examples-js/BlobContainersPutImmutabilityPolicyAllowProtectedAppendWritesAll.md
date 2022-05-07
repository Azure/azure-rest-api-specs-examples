Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-storage_17.2.0/sdk/storage/arm-storage/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { StorageManagementClient } = require("@azure/arm-storage");
const { DefaultAzureCredential } = require("@azure/identity");

async function createOrUpdateImmutabilityPolicyWithAllowProtectedAppendWritesAll() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "res1782";
  const accountName = "sto7069";
  const containerName = "container6397";
  const parameters = {
    allowProtectedAppendWritesAll: true,
    immutabilityPeriodSinceCreationInDays: 3,
  };
  const options = { parameters: parameters };
  const credential = new DefaultAzureCredential();
  const client = new StorageManagementClient(credential, subscriptionId);
  const result = await client.blobContainers.createOrUpdateImmutabilityPolicy(
    resourceGroupName,
    accountName,
    containerName,
    options
  );
  console.log(result);
}

createOrUpdateImmutabilityPolicyWithAllowProtectedAppendWritesAll().catch(console.error);
```
