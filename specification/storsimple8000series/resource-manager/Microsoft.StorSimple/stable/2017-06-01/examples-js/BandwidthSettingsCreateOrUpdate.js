const { StorSimple8000SeriesManagementClient } = require("@azure/arm-storsimple8000series");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates the bandwidth setting
 *
 * @summary Creates or updates the bandwidth setting
 * x-ms-original-file: specification/storsimple8000series/resource-manager/Microsoft.StorSimple/stable/2017-06-01/examples/BandwidthSettingsCreateOrUpdate.json
 */
async function bandwidthSettingsCreateOrUpdate() {
  const subscriptionId = "4385cf00-2d3a-425a-832f-f4285b1c9dce";
  const bandwidthSettingName = "BWSForTest";
  const resourceGroupName = "ResourceGroupForSDKTest";
  const managerName = "ManagerForSDKTest1";
  const parameters = {
    schedules: [
      {
        days: ["Saturday", "Sunday"],
        rateInMbps: 10,
        start: { hours: 10, minutes: 0, seconds: 0 },
        stop: { hours: 20, minutes: 0, seconds: 0 },
      },
    ],
  };
  const credential = new DefaultAzureCredential();
  const client = new StorSimple8000SeriesManagementClient(credential, subscriptionId);
  const result = await client.bandwidthSettings.beginCreateOrUpdateAndWait(
    bandwidthSettingName,
    resourceGroupName,
    managerName,
    parameters
  );
  console.log(result);
}

bandwidthSettingsCreateOrUpdate().catch(console.error);
