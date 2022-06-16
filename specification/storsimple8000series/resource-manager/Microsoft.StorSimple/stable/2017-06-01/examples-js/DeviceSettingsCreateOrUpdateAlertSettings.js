const { StorSimple8000SeriesManagementClient } = require("@azure/arm-storsimple8000series");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates the alert settings of the specified device.
 *
 * @summary Creates or updates the alert settings of the specified device.
 * x-ms-original-file: specification/storsimple8000series/resource-manager/Microsoft.StorSimple/stable/2017-06-01/examples/DeviceSettingsCreateOrUpdateAlertSettings.json
 */
async function deviceSettingsCreateOrUpdateAlertSettings() {
  const subscriptionId = "4385cf00-2d3a-425a-832f-f4285b1c9dce";
  const deviceName = "Device05ForSDKTest";
  const resourceGroupName = "ResourceGroupForSDKTest";
  const managerName = "ManagerForSDKTest1";
  const parameters = {
    additionalRecipientEmailList: [],
    alertNotificationCulture: "en-US",
    emailNotification: "Enabled",
    notificationToServiceOwners: "Enabled",
  };
  const credential = new DefaultAzureCredential();
  const client = new StorSimple8000SeriesManagementClient(credential, subscriptionId);
  const result = await client.deviceSettings.beginCreateOrUpdateAlertSettingsAndWait(
    deviceName,
    resourceGroupName,
    managerName,
    parameters
  );
  console.log(result);
}

deviceSettingsCreateOrUpdateAlertSettings().catch(console.error);
