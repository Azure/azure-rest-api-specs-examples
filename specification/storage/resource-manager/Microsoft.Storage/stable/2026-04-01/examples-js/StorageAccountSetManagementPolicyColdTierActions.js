const { StorageManagementClient } = require("@azure/arm-storage");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to sets the managementpolicy to the specified storage account.
 *
 * @summary sets the managementpolicy to the specified storage account.
 * x-ms-original-file: 2026-04-01/StorageAccountSetManagementPolicyColdTierActions.json
 */
async function storageAccountSetManagementPolicyColdTierActions() {
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
                delete: { daysAfterModificationGreaterThan: 1000 },
                tierToArchive: { daysAfterModificationGreaterThan: 90 },
                tierToCold: { daysAfterModificationGreaterThan: 30 },
                tierToCool: { daysAfterModificationGreaterThan: 30 },
              },
              snapshot: {
                delete: { daysAfterCreationGreaterThan: 30 },
                tierToCold: { daysAfterCreationGreaterThan: 30 },
              },
              version: {
                delete: { daysAfterCreationGreaterThan: 30 },
                tierToCold: { daysAfterCreationGreaterThan: 30 },
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
