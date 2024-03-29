const { StorageManagementClient } = require("@azure/arm-storage-profile-2020-09-01-hybrid");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Sets the blob inventory policy to the specified storage account.
 *
 * @summary Sets the blob inventory policy to the specified storage account.
 * x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2019-06-01/examples/StorageAccountSetBlobInventoryPolicy.json
 */
async function storageAccountSetBlobInventoryPolicy() {
  const subscriptionId = process.env["STORAGE_SUBSCRIPTION_ID"] || "{subscription-id}";
  const resourceGroupName = process.env["STORAGE_RESOURCE_GROUP"] || "res7687";
  const accountName = "sto9699";
  const blobInventoryPolicyName = "default";
  const properties = {
    policy: {
      type: "Inventory",
      destination: "containerName",
      enabled: true,
      rules: [
        {
          name: "inventoryPolicyRule1",
          definition: {
            filters: {
              blobTypes: ["blockBlob", "appendBlob", "pageBlob"],
              includeBlobVersions: true,
              includeSnapshots: true,
              prefixMatch: ["inventoryprefix1", "inventoryprefix2"],
            },
          },
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
