const { StorSimple8000SeriesManagementClient } = require("@azure/arm-storsimple8000series");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Downloads and installs the updates on the device.
 *
 * @summary Downloads and installs the updates on the device.
 * x-ms-original-file: specification/storsimple8000series/resource-manager/Microsoft.StorSimple/stable/2017-06-01/examples/DevicesInstallUpdates.json
 */
async function devicesInstallUpdates() {
  const subscriptionId = "4385cf00-2d3a-425a-832f-f4285b1c9dce";
  const deviceName = "sugattdeviceforSDK";
  const resourceGroupName = "ResourceGroupForSDKTest";
  const managerName = "ManagerForSDKTest1";
  const credential = new DefaultAzureCredential();
  const client = new StorSimple8000SeriesManagementClient(credential, subscriptionId);
  const result = await client.devices.beginInstallUpdatesAndWait(
    deviceName,
    resourceGroupName,
    managerName
  );
  console.log(result);
}

devicesInstallUpdates().catch(console.error);
