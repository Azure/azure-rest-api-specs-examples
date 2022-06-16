const { StorageCacheManagementClient } = require("@azure/arm-storagecache");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Resumes client access to a previously suspended storage target.
 *
 * @summary Resumes client access to a previously suspended storage target.
 * x-ms-original-file: specification/storagecache/resource-manager/Microsoft.StorageCache/stable/2022-01-01/examples/StorageTargets_Resume.json
 */
async function storageTargetsResume() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "scgroup";
  const cacheName = "sc";
  const storageTargetName = "st1";
  const credential = new DefaultAzureCredential();
  const client = new StorageCacheManagementClient(credential, subscriptionId);
  const result = await client.storageTargetOperations.beginResumeAndWait(
    resourceGroupName,
    cacheName,
    storageTargetName
  );
  console.log(result);
}

storageTargetsResume().catch(console.error);
