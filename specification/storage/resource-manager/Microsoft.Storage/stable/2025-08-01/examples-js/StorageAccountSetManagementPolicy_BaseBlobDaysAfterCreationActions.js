const { StorageManagementClient } = require("@azure/arm-storage");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to sets the managementpolicy to the specified storage account.
 *
 * @summary sets the managementpolicy to the specified storage account.
 * x-ms-original-file: 2025-08-01/StorageAccountSetManagementPolicy_BaseBlobDaysAfterCreationActions.json
 */
async function storageAccountSetManagementPolicyBaseBlobDaysAfterCreationActions() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new StorageManagementClient(credential, subscriptionId);
  const result = await client.managementPolicies.createOrUpdate("res7687", "sto9699", "default", {
    policy: {
      rules: [
        {
          name: "olcmtest1",
          type: "Lifecycle",
          definition: {
            actions: {
              baseBlob: {
                delete: { daysAfterCreationGreaterThan: 1000 },
                tierToArchive: { daysAfterCreationGreaterThan: 90 },
                tierToCool: { daysAfterCreationGreaterThan: 30 },
              },
            },
            filters: { blobTypes: ["blockBlob"], prefixMatch: ["olcmtestcontainer1"] },
          },
          enabled: true,
        },
      ],
    },
  });
  console.log(result);
}
