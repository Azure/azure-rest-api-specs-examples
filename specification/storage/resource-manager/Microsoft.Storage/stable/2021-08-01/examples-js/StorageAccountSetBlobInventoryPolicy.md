Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-storage_17.2.0/sdk/storage/arm-storage/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { StorageManagementClient } = require("@azure/arm-storage");
const { DefaultAzureCredential } = require("@azure/identity");

async function storageAccountSetBlobInventoryPolicy() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "res7687";
  const accountName = "sto9699";
  const blobInventoryPolicyName = "default";
  const properties = {
    policy: {
      type: "Inventory",
      enabled: true,
      rules: [
        {
          name: "inventoryPolicyRule1",
          definition: {
            format: "Csv",
            filters: {
              blobTypes: ["blockBlob", "appendBlob", "pageBlob"],
              includeBlobVersions: true,
              includeSnapshots: true,
              prefixMatch: ["inventoryprefix1", "inventoryprefix2"],
            },
            objectType: "Blob",
            schedule: "Daily",
            schemaFields: [
              "Name",
              "Creation-Time",
              "Last-Modified",
              "Content-Length",
              "Content-MD5",
              "BlobType",
              "AccessTier",
              "AccessTierChangeTime",
              "Snapshot",
              "VersionId",
              "IsCurrentVersion",
              "Metadata",
            ],
          },
          destination: "container1",
          enabled: true,
        },
        {
          name: "inventoryPolicyRule2",
          definition: {
            format: "Parquet",
            objectType: "Container",
            schedule: "Weekly",
            schemaFields: [
              "Name",
              "Last-Modified",
              "Metadata",
              "LeaseStatus",
              "LeaseState",
              "LeaseDuration",
              "PublicAccess",
              "HasImmutabilityPolicy",
              "HasLegalHold",
            ],
          },
          destination: "container2",
          enabled: true,
        },
      ],
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new StorageManagementClient(credential, subscriptionId);
  const result = await client.blobInventoryPolicies.createOrUpdate(
    resourceGroupName,
    accountName,
    blobInventoryPolicyName,
    properties
  );
  console.log(result);
}

storageAccountSetBlobInventoryPolicy().catch(console.error);
```
