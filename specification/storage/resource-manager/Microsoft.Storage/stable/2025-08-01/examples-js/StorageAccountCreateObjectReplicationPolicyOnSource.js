const { StorageManagementClient } = require("@azure/arm-storage");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to create or update the object replication policy of the storage account.
 *
 * @summary create or update the object replication policy of the storage account.
 * x-ms-original-file: 2025-08-01/StorageAccountCreateObjectReplicationPolicyOnSource.json
 */
async function storageAccountCreateObjectReplicationPolicyOnSource() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new StorageManagementClient(credential, subscriptionId);
  const result = await client.objectReplicationPolicies.createOrUpdate(
    "res7687",
    "src1122",
    "2a20bb73-5717-4635-985a-5d4cf777438f",
    {
      destinationAccount: "dst112",
      metrics: { enabled: true },
      priorityReplication: { enabled: true },
      rules: [
        {
          destinationContainer: "dcont139",
          filters: { minCreationTime: "2020-02-19T16:05:00Z", prefixMatch: ["blobA", "blobB"] },
          ruleId: "d5d18a48-8801-4554-aeaa-74faf65f5ef9",
          sourceContainer: "scont139",
        },
      ],
      sourceAccount: "src1122",
      tagsReplication: { enabled: true },
    },
  );
  console.log(result);
}
