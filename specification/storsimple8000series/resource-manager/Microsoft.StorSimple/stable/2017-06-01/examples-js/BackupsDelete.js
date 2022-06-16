const { StorSimple8000SeriesManagementClient } = require("@azure/arm-storsimple8000series");
const { DefaultAzureCredential } = require("@azure/identity");

async function backupsDelete() {
  const subscriptionId = "4385cf00-2d3a-425a-832f-f4285b1c9dce";
  const deviceName = "Device05ForSDKTest";
  const backupName = "880e1774-94a8-4f3e-85e6-a61e6b94a8b7";
  const resourceGroupName = "ResourceGroupForSDKTest";
  const managerName = "ManagerForSDKTest1";
  const credential = new DefaultAzureCredential();
  const client = new StorSimple8000SeriesManagementClient(credential, subscriptionId);
  const result = await client.backups.beginDeleteAndWait(
    deviceName,
    backupName,
    resourceGroupName,
    managerName
  );
  console.log(result);
}

backupsDelete().catch(console.error);
