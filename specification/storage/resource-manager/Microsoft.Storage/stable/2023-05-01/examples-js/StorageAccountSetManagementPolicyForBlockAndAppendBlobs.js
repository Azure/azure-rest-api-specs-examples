const { StorageManagementClient } = require("@azure/arm-storage");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Sets the managementpolicy to the specified storage account.
 *
 * @summary Sets the managementpolicy to the specified storage account.
 * x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2023-05-01/examples/StorageAccountSetManagementPolicyForBlockAndAppendBlobs.json
 */
async function storageAccountSetManagementPolicyForBlockAndAppendBlobs() {
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
              baseBlob: { delete: { daysAfterModificationGreaterThan: 90 } },
              snapshot: { delete: { daysAfterCreationGreaterThan: 90 } },
              version: { delete: { daysAfterCreationGreaterThan: 90 } },
            },
            filters: {
              blobTypes: ["blockBlob", "appendBlob"],
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
    properties,
  );
  console.log(result);
}
