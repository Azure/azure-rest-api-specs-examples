const { StorageManagementClient } = require("@azure/arm-storage");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to sets the managementpolicy to the specified storage account.
 *
 * @summary sets the managementpolicy to the specified storage account.
 * x-ms-original-file: 2026-04-01/StorageAccountSetManagementPolicy_LastTierChangeTimeActions.json
 */
async function storageAccountSetManagementPolicyLastTierChangeTimeActions() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new StorageManagementClient(credential, subscriptionId);
  const result = await client.managementPolicies.createOrUpdate("res7687", "sto9699", "default", {
    policy: {
      rules: [
        {
          name: "olcmtest",
          type: "Lifecycle",
          definition: {
            actions: {
              baseBlob: {
                delete: { daysAfterModificationGreaterThan: 1000 },
                tierToArchive: {
                  daysAfterLastTierChangeGreaterThan: 120,
                  daysAfterModificationGreaterThan: 90,
                },
                tierToCool: { daysAfterModificationGreaterThan: 30 },
              },
              snapshot: {
                tierToArchive: {
                  daysAfterCreationGreaterThan: 30,
                  daysAfterLastTierChangeGreaterThan: 90,
                },
              },
              version: {
                tierToArchive: {
                  daysAfterCreationGreaterThan: 30,
                  daysAfterLastTierChangeGreaterThan: 90,
                },
              },
            },
            filters: { blobTypes: ["blockBlob"], prefixMatch: ["olcmtestcontainer"] },
          },
          enabled: true,
        },
      ],
    },
  });
  console.log(result);
}
