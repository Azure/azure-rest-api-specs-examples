const { StorSimple8000SeriesManagementClient } = require("@azure/arm-storsimple8000series");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to sync Remote management Certificate between appliance and Service
 *
 * @summary sync Remote management Certificate between appliance and Service
 * x-ms-original-file: specification/storsimple8000series/resource-manager/Microsoft.StorSimple/stable/2017-06-01/examples/DeviceSettingsSyncRemotemanagementCertificate.json
 */
async function deviceSettingsSyncRemotemanagementCertificate() {
  const subscriptionId = "4385cf00-2d3a-425a-832f-f4285b1c9dce";
  const deviceName = "Device05ForSDKTest";
  const resourceGroupName = "ResourceGroupForSDKTest";
  const managerName = "ManagerForSDKTest1";
  const credential = new DefaultAzureCredential();
  const client = new StorSimple8000SeriesManagementClient(credential, subscriptionId);
  const result = await client.deviceSettings.beginSyncRemotemanagementCertificateAndWait(
    deviceName,
    resourceGroupName,
    managerName
  );
  console.log(result);
}

deviceSettingsSyncRemotemanagementCertificate().catch(console.error);
