const { StorSimpleManagementClient } = require("@azure/arm-storsimple1200series");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Downloads updates on the device.
 *
 * @summary Downloads updates on the device.
 * x-ms-original-file: specification/storSimple1200Series/resource-manager/Microsoft.StorSimple/stable/2016-10-01/examples/DevicesDownloadUpdates.json
 */
async function devicesDownloadUpdates() {
  const subscriptionId = "4385cf00-2d3a-425a-832f-f4285b1c9dce";
  const deviceName = "sugattdeviceforSDK";
  const resourceGroupName = "ResourceGroupForSDKTest";
  const managerName = "ManagerForSDKTest1";
  const credential = new DefaultAzureCredential();
  const client = new StorSimpleManagementClient(credential, subscriptionId);
  const result = await client.devices.beginDownloadUpdatesAndWait(
    deviceName,
    resourceGroupName,
    managerName
  );
  console.log(result);
}

devicesDownloadUpdates().catch(console.error);
