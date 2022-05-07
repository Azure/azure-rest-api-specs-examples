Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-storsimple1200series_2.0.1/sdk/storsimple1200series/arm-storsimple1200series/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { StorSimpleManagementClient } = require("@azure/arm-storsimple1200series");
const { DefaultAzureCredential } = require("@azure/identity");

async function backupScheduleGroupsCreateOrUpdate() {
  const subscriptionId = "9eb689cd-7243-43b4-b6f6-5c65cb296641";
  const deviceName = "HSDK-4XY4FI2IVG";
  const scheduleGroupName = "BackupSchGroupForSDKTest";
  const resourceGroupName = "ResourceGroupForSDKTest";
  const managerName = "hAzureSDKOperations";
  const scheduleGroup = {
    name: "BackupSchGroupForSDKTest",
    startTime: { hour: 17, minute: 38 },
  };
  const credential = new DefaultAzureCredential();
  const client = new StorSimpleManagementClient(credential, subscriptionId);
  const result = await client.backupScheduleGroups.beginCreateOrUpdateAndWait(
    deviceName,
    scheduleGroupName,
    resourceGroupName,
    managerName,
    scheduleGroup
  );
  console.log(result);
}

backupScheduleGroupsCreateOrUpdate().catch(console.error);
```
