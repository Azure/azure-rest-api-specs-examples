const { BlockClient } = require("@azure/arm-purestorageblock");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to update a storage pool
 *
 * @summary update a storage pool
 * x-ms-original-file: 2024-11-01-preview/StoragePools_Update_MaximumSet_Gen.json
 */
async function storagePoolsUpdate() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "BC47D6CC-AA80-4374-86F8-19D94EC70666";
  const client = new BlockClient(credential, subscriptionId);
  const result = await client.storagePools.update("rgpurestorage", "storagePoolname", {
    identity: { type: "None", userAssignedIdentities: { key4211: {} } },
    tags: { key9065: "ebgmkwxqewe" },
    properties: { provisionedBandwidthMbPerSec: 23 },
  });
  console.log(result);
}
