Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-storage_17.2.0/sdk/storage/arm-storage/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { StorageManagementClient } = require("@azure/arm-storage");
const { DefaultAzureCredential } = require("@azure/identity");

async function versionLevelWormContainerMigration() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "res1782";
  const accountName = "sto7069";
  const containerName = "container6397";
  const credential = new DefaultAzureCredential();
  const client = new StorageManagementClient(credential, subscriptionId);
  const result = await client.blobContainers.beginObjectLevelWormAndWait(
    resourceGroupName,
    accountName,
    containerName
  );
  console.log(result);
}

versionLevelWormContainerMigration().catch(console.error);
```
