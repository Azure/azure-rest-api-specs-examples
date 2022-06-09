```javascript
const { StorSimple8000SeriesManagementClient } = require("@azure/arm-storsimple8000series");
const { DefaultAzureCredential } = require("@azure/identity");

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
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-storsimple8000series_2.0.1/sdk/storsimple8000series/arm-storsimple8000series/README.md) on how to add the SDK to your project and authenticate.
