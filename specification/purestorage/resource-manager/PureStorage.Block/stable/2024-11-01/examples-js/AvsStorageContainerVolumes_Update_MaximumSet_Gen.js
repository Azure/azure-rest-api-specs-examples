const { BlockClient } = require("@azure/arm-purestorageblock");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to update a volume in an AVS storage container
 *
 * @summary update a volume in an AVS storage container
 * x-ms-original-file: 2024-11-01/AvsStorageContainerVolumes_Update_MaximumSet_Gen.json
 */
async function avsStorageContainerVolumesUpdate() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "BC47D6CC-AA80-4374-86F8-19D94EC70666";
  const client = new BlockClient(credential, subscriptionId);
  const result = await client.avsStorageContainerVolumes.update(
    "rgpurestorage",
    "storagePoolname",
    "name",
    "cbdec-ddbb",
    { properties: { softDeletion: { destroyed: true } } },
  );
  console.log(result);
}
