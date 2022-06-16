const { StorSimple8000SeriesManagementClient } = require("@azure/arm-storsimple8000series");
const { DefaultAzureCredential } = require("@azure/identity");

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
