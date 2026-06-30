const { StorageManagementClient } = require("@azure/arm-storage");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to create or update the object replication policy of the storage account.
 *
 * @summary create or update the object replication policy of the storage account.
 * x-ms-original-file: 2026-04-01/StorageAccountCreateObjectReplicationPolicyOnDestination.json
 */
async function storageAccountCreateObjectReplicationPolicyOnDestination() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new StorageManagementClient(credential, subscriptionId);
  const result = await client.objectReplicationPolicies.createOrUpdate(
    "res7687",
    "dst112",
    "default",
    {
      destinationAccount: "dst112",
      metrics: { enabled: true },
      priorityReplication: { enabled: true },
      rules: [
        {
          destinationContainer: "dcont139",
          filters: { prefixMatch: ["blobA", "blobB"] },
          sourceContainer: "scont139",
        },
      ],
      sourceAccount: "src1122",
      tagsReplication: { enabled: true },
    },
  );
  console.log(result);
}
