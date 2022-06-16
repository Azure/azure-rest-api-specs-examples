const { StorSimpleManagementClient } = require("@azure/arm-storsimple1200series");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Returns the time settings of the specified device name.
 *
 * @summary Returns the time settings of the specified device name.
 * x-ms-original-file: specification/storSimple1200Series/resource-manager/Microsoft.StorSimple/stable/2016-10-01/examples/DevicesGetTimeSettings.json
 */
async function devicesGetTimeSettings() {
  const subscriptionId = "9eb689cd-7243-43b4-b6f6-5c65cb296641";
  const deviceName = "HSDK-T4ZA3EAJFR";
  const resourceGroupName = "ResourceGroupForSDKTest";
  const managerName = "hAzureSDKOperations";
  const credential = new DefaultAzureCredential();
  const client = new StorSimpleManagementClient(credential, subscriptionId);
  const result = await client.devices.getTimeSettings(deviceName, resourceGroupName, managerName);
  console.log(result);
}

devicesGetTimeSettings().catch(console.error);
