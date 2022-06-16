const { StorSimple8000SeriesManagementClient } = require("@azure/arm-storsimple8000series");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Cancels a job on the device.
 *
 * @summary Cancels a job on the device.
 * x-ms-original-file: specification/storsimple8000series/resource-manager/Microsoft.StorSimple/stable/2017-06-01/examples/JobsCancel.json
 */
async function jobsCancel() {
  const subscriptionId = "4385cf00-2d3a-425a-832f-f4285b1c9dce";
  const deviceName = "Device05ForSDKTest";
  const jobName = "993db21d-101b-41af-9e12-f593d78b99e9";
  const resourceGroupName = "ResourceGroupForSDKTest";
  const managerName = "ManagerForSDKTest1";
  const credential = new DefaultAzureCredential();
  const client = new StorSimple8000SeriesManagementClient(credential, subscriptionId);
  const result = await client.jobs.beginCancelAndWait(
    deviceName,
    jobName,
    resourceGroupName,
    managerName
  );
  console.log(result);
}

jobsCancel().catch(console.error);
