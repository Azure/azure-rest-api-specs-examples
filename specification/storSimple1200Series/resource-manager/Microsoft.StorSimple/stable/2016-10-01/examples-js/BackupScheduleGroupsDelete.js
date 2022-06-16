const { StorSimpleManagementClient } = require("@azure/arm-storsimple1200series");
const { DefaultAzureCredential } = require("@azure/identity");

async function backupScheduleGroupsDelete() {
  const subscriptionId = "9eb689cd-7243-43b4-b6f6-5c65cb296641";
  const deviceName = "HSDK-DMNJB2PET0";
  const scheduleGroupName = "BackupSchGroupForSDKTest";
  const resourceGroupName = "ResourceGroupForSDKTest";
  const managerName = "hAzureSDKOperations";
  const credential = new DefaultAzureCredential();
  const client = new StorSimpleManagementClient(credential, subscriptionId);
  const result = await client.backupScheduleGroups.beginDeleteAndWait(
    deviceName,
    scheduleGroupName,
    resourceGroupName,
    managerName
  );
  console.log(result);
}

backupScheduleGroupsDelete().catch(console.error);
