const { StorSimple8000SeriesManagementClient } = require("@azure/arm-storsimple8000series");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes the backup schedule.
 *
 * @summary Deletes the backup schedule.
 * x-ms-original-file: specification/storsimple8000series/resource-manager/Microsoft.StorSimple/stable/2017-06-01/examples/BackupSchedulesDelete.json
 */
async function backupSchedulesDelete() {
  const subscriptionId = "4385cf00-2d3a-425a-832f-f4285b1c9dce";
  const deviceName = "Device05ForSDKTest";
  const backupPolicyName = "BkUpPolicy01ForSDKTest";
  const backupScheduleName = "schedule1";
  const resourceGroupName = "ResourceGroupForSDKTest";
  const managerName = "ManagerForSDKTest1";
  const credential = new DefaultAzureCredential();
  const client = new StorSimple8000SeriesManagementClient(credential, subscriptionId);
  const result = await client.backupSchedules.beginDeleteAndWait(
    deviceName,
    backupPolicyName,
    backupScheduleName,
    resourceGroupName,
    managerName
  );
  console.log(result);
}

backupSchedulesDelete().catch(console.error);
