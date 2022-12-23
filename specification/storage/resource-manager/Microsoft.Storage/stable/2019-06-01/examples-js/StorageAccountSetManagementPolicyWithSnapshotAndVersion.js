const { StorageManagementClient } = require("@azure/arm-storage-profile-2020-09-01-hybrid");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Sets the managementpolicy to the specified storage account.
 *
 * @summary Sets the managementpolicy to the specified storage account.
 * x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2019-06-01/examples/StorageAccountSetManagementPolicyWithSnapshotAndVersion.json
 */
async function storageAccountSetManagementPolicyWithSnapshotAndVersion() {
  const subscriptionId = process.env["STORAGE_SUBSCRIPTION_ID"] || "{subscription-id}";
  const resourceGroupName = process.env["STORAGE_RESOURCE_GROUP"] || "res7687";
  const accountName = "sto9699";
  const managementPolicyName = "default";
  const properties = {
    policy: {
      rules: [
        {
          name: "olcmtest1",
          type: "Lifecycle",
          definition: {
            actions: {
              baseBlob: {
                delete: { daysAfterModificationGreaterThan: 1000 },
                tierToArchive: { daysAfterModificationGreaterThan: 90 },
                tierToCool: { daysAfterModificationGreaterThan: 30 },
              },
              snapshot: {
                delete: { daysAfterCreationGreaterThan: 1000 },
                tierToArchive: { daysAfterCreationGreaterThan: 90 },
                tierToCool: { daysAfterCreationGreaterThan: 30 },
              },
              version: {
                delete: { daysAfterCreationGreaterThan: 1000 },
                tierToArchive: { daysAfterCreationGreaterThan: 90 },
                tierToCool: { daysAfterCreationGreaterThan: 30 },
              },
            },
            filters: {
              blobTypes: ["blockBlob"],
              prefixMatch: ["olcmtestcontainer1"],
            },
          },
          enabled: true,
        },
      ],
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new StorageManagementClient(credential, subscriptionId);
  const result = await client.managementPolicies.createOrUpdate(
    resourceGroupName,
    accountName,
    managementPolicyName,
    properties
  );
  console.log(result);
}
