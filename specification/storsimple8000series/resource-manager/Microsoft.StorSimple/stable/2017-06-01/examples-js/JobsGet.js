const { StorSimple8000SeriesManagementClient } = require("@azure/arm-storsimple8000series");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the details of the specified job name.
 *
 * @summary Gets the details of the specified job name.
 * x-ms-original-file: specification/storsimple8000series/resource-manager/Microsoft.StorSimple/stable/2017-06-01/examples/JobsGet.json
 */
async function jobsGet() {
  const subscriptionId = "d3ebfe71-b7a9-4c57-92b9-68a2afde4de5";
  const deviceName = "sca07forsdktest";
  const jobName = "70a29339-de6d-48e8-b24f-e25ee6769a00";
  const resourceGroupName = "ResourceGroupForSDKTest";
  const managerName = "ManagerForSDKTest1";
  const credential = new DefaultAzureCredential();
  const client = new StorSimple8000SeriesManagementClient(credential, subscriptionId);
  const result = await client.jobs.get(deviceName, jobName, resourceGroupName, managerName);
  console.log(result);
}

jobsGet().catch(console.error);
