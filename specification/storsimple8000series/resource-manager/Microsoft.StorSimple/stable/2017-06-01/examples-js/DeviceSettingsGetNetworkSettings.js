const { StorSimple8000SeriesManagementClient } = require("@azure/arm-storsimple8000series");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the network settings of the specified device.
 *
 * @summary Gets the network settings of the specified device.
 * x-ms-original-file: specification/storsimple8000series/resource-manager/Microsoft.StorSimple/stable/2017-06-01/examples/DeviceSettingsGetNetworkSettings.json
 */
async function deviceSettingsGetNetworkSettings() {
  const subscriptionId = "4385cf00-2d3a-425a-832f-f4285b1c9dce";
  const deviceName = "Device05ForSDKTest";
  const resourceGroupName = "ResourceGroupForSDKTest";
  const managerName = "ManagerForSDKTest1";
  const credential = new DefaultAzureCredential();
  const client = new StorSimple8000SeriesManagementClient(credential, subscriptionId);
  const result = await client.deviceSettings.getNetworkSettings(
    deviceName,
    resourceGroupName,
    managerName
  );
  console.log(result);
}

deviceSettingsGetNetworkSettings().catch(console.error);
