Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-storsimple8000series_2.0.1/sdk/storsimple8000series/arm-storsimple8000series/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { StorSimple8000SeriesManagementClient } = require("@azure/arm-storsimple8000series");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Failovers a set of volume containers from a specified source device to a target device.
 *
 * @summary Failovers a set of volume containers from a specified source device to a target device.
 * x-ms-original-file: specification/storsimple8000series/resource-manager/Microsoft.StorSimple/stable/2017-06-01/examples/DevicesFailover.json
 */
async function devicesFailover() {
  const subscriptionId = "4385cf00-2d3a-425a-832f-f4285b1c9dce";
  const sourceDeviceName = "Device05ForSDKTest";
  const resourceGroupName = "ResourceGroupForSDKTest";
  const managerName = "ManagerForSDKTest1";
  const parameters = {
    targetDeviceId:
      "/subscriptions/4385cf00-2d3a-425a-832f-f4285b1c9dce/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/ManagerForSDKTest1/devices/jemdeviceforsdk",
    volumeContainers: [
      "/subscriptions/4385cf00-2d3a-425a-832f-f4285b1c9dce/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/ManagerForSDKTest1/devices/Device05ForSDKTest/volumeContainers/vcforsdktest",
    ],
  };
  const credential = new DefaultAzureCredential();
  const client = new StorSimple8000SeriesManagementClient(credential, subscriptionId);
  const result = await client.devices.beginFailoverAndWait(
    sourceDeviceName,
    resourceGroupName,
    managerName,
    parameters
  );
  console.log(result);
}

devicesFailover().catch(console.error);
```
