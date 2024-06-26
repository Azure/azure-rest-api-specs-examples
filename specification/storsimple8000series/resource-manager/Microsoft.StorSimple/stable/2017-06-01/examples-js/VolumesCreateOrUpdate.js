const { StorSimple8000SeriesManagementClient } = require("@azure/arm-storsimple8000series");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates the volume.
 *
 * @summary Creates or updates the volume.
 * x-ms-original-file: specification/storsimple8000series/resource-manager/Microsoft.StorSimple/stable/2017-06-01/examples/VolumesCreateOrUpdate.json
 */
async function volumesCreateOrUpdate() {
  const subscriptionId = "4385cf00-2d3a-425a-832f-f4285b1c9dce";
  const deviceName = "Device05ForSDKTest";
  const volumeContainerName = "VolumeContainerForSDKTest";
  const volumeName = "Volume1ForSDKTest";
  const resourceGroupName = "ResourceGroupForSDKTest";
  const managerName = "ManagerForSDKTest1";
  const parameters = {
    accessControlRecordIds: [
      "/subscriptions/4385cf00-2d3a-425a-832f-f4285b1c9dce/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/ManagerForSDKTest1/accessControlRecords/ACR2",
    ],
    monitoringStatus: "Enabled",
    sizeInBytes: 5368709120,
    volumeStatus: "Offline",
    volumeType: "Tiered",
  };
  const credential = new DefaultAzureCredential();
  const client = new StorSimple8000SeriesManagementClient(credential, subscriptionId);
  const result = await client.volumes.beginCreateOrUpdateAndWait(
    deviceName,
    volumeContainerName,
    volumeName,
    resourceGroupName,
    managerName,
    parameters
  );
  console.log(result);
}

volumesCreateOrUpdate().catch(console.error);
