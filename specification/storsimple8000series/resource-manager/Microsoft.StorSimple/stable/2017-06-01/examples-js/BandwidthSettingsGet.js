const { StorSimple8000SeriesManagementClient } = require("@azure/arm-storsimple8000series");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Returns the properties of the specified bandwidth setting name.
 *
 * @summary Returns the properties of the specified bandwidth setting name.
 * x-ms-original-file: specification/storsimple8000series/resource-manager/Microsoft.StorSimple/stable/2017-06-01/examples/BandwidthSettingsGet.json
 */
async function bandwidthSettingsGet() {
  const subscriptionId = "4385cf00-2d3a-425a-832f-f4285b1c9dce";
  const bandwidthSettingName = "bandwidthSetting1";
  const resourceGroupName = "ResourceGroupForSDKTest";
  const managerName = "ManagerForSDKTest1";
  const credential = new DefaultAzureCredential();
  const client = new StorSimple8000SeriesManagementClient(credential, subscriptionId);
  const result = await client.bandwidthSettings.get(
    bandwidthSettingName,
    resourceGroupName,
    managerName
  );
  console.log(result);
}

bandwidthSettingsGet().catch(console.error);
