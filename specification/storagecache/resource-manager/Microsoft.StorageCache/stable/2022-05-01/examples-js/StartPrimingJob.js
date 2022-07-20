const { StorageCacheManagementClient } = require("@azure/arm-storagecache");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create a priming job. This operation is only allowed when the cache is healthy.
 *
 * @summary Create a priming job. This operation is only allowed when the cache is healthy.
 * x-ms-original-file: specification/storagecache/resource-manager/Microsoft.StorageCache/stable/2022-05-01/examples/StartPrimingJob.json
 */
async function startPrimingJob() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "scgroup";
  const cacheName = "sc1";
  const primingjob = {
    primingJobName: "contosoJob",
    primingManifestUrl:
      "https://contosostorage.blob.core.windows.net/contosoblob/00000000_00000000000000000000000000000000.00000000000.FFFFFFFF.00000000?sp=r&st=2021-08-11T19:33:35Z&se=2021-08-12T03:33:35Z&spr=https&sv=2020-08-04&sr=b&sig=<secret-value-from-key>",
  };
  const options = { primingjob };
  const credential = new DefaultAzureCredential();
  const client = new StorageCacheManagementClient(credential, subscriptionId);
  const result = await client.caches.beginStartPrimingJobAndWait(
    resourceGroupName,
    cacheName,
    options
  );
  console.log(result);
}

startPrimingJob().catch(console.error);
