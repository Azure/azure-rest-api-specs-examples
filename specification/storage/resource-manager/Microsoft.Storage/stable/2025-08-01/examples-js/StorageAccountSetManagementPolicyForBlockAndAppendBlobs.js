const { StorageManagementClient } = require("@azure/arm-storage");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to sets the managementpolicy to the specified storage account.
 *
 * @summary sets the managementpolicy to the specified storage account.
 * x-ms-original-file: 2025-08-01/StorageAccountSetManagementPolicyForBlockAndAppendBlobs.json
 */
async function storageAccountSetManagementPolicyForBlockAndAppendBlobs() {
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
  });
  console.log(result);
}
