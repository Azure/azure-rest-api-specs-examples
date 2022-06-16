const { StorSimpleManagementClient } = require("@azure/arm-storsimple1200series");
const { DefaultAzureCredential } = require("@azure/identity");

async function backupScheduleGroupsGet() {
  const subscriptionId = "9eb689cd-7243-43b4-b6f6-5c65cb296641";
  const deviceName = "HSDK-4XY4FI2IVG";
  const scheduleGroupName = "BackupSchGroupForSDKTest";
  const resourceGroupName = "ResourceGroupForSDKTest";
  const managerName = "hAzureSDKOperations";
  const credential = new DefaultAzureCredential();
  const client = new StorSimpleManagementClient(credential, subscriptionId);
  const result = await client.backupScheduleGroups.get(
    deviceName,
    scheduleGroupName,
    resourceGroupName,
    managerName
  );
  console.log(result);
}

backupScheduleGroupsGet().catch(console.error);
