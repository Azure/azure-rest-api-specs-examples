const { StorageManagementClient } = require("@azure/arm-storage");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Sets the blob inventory policy to the specified storage account.
 *
 * @summary Sets the blob inventory policy to the specified storage account.
 * x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2023-01-01/examples/StorageAccountSetBlobInventoryPolicy.json
 */
async function storageAccountSetBlobInventoryPolicy() {
  const subscriptionId = process.env["STORAGE_SUBSCRIPTION_ID"] || "{subscription-id}";
  const resourceGroupName = process.env["STORAGE_RESOURCE_GROUP"] || "res7687";
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
              creationTime: { lastNDays: 1000 },
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
