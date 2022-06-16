```javascript
const { StorageManagementClient } = require("@azure/arm-storage");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Sets the blob inventory policy to the specified storage account.
 *
 * @summary Sets the blob inventory policy to the specified storage account.
 * x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2021-09-01/examples/StorageAccountSetBlobInventoryPolicyIncludeDeleteAndNewSchemaForNonHnsAccount.json
 */
async function storageAccountSetBlobInventoryPolicyIncludeDeleteAndNewSchemaForNonHnsAccount() {
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
              excludePrefix: ["excludeprefix1", "excludeprefix2"],
              includeBlobVersions: true,
              includeDeleted: true,
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
              "Tags",
              "ContentType",
              "ContentEncoding",
              "ContentLanguage",
              "ContentCRC64",
              "CacheControl",
              "Metadata",
              "Deleted",
              "RemainingRetentionDays",
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
              "Etag",
              "DefaultEncryptionScope",
              "DenyEncryptionScopeOverride",
              "ImmutableStorageWithVersioningEnabled",
              "Deleted",
              "Version",
              "DeletedTime",
              "RemainingRetentionDays",
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

storageAccountSetBlobInventoryPolicyIncludeDeleteAndNewSchemaForNonHnsAccount().catch(
  console.error
);
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-storage_17.2.1/sdk/storage/arm-storage/README.md) on how to add the SDK to your project and authenticate.
