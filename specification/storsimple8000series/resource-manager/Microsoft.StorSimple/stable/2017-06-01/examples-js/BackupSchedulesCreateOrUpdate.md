```javascript
const { StorSimple8000SeriesManagementClient } = require("@azure/arm-storsimple8000series");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates the backup schedule.
 *
 * @summary Creates or updates the backup schedule.
 * x-ms-original-file: specification/storsimple8000series/resource-manager/Microsoft.StorSimple/stable/2017-06-01/examples/BackupSchedulesCreateOrUpdate.json
 */
async function backupSchedulesCreateOrUpdate() {
  const subscriptionId = "4385cf00-2d3a-425a-832f-f4285b1c9dce";
  const deviceName = "Device05ForSDKTest";
  const backupPolicyName = "BkUpPolicy01ForSDKTest";
  const backupScheduleName = "schedule2";
  const resourceGroupName = "ResourceGroupForSDKTest";
  const managerName = "ManagerForSDKTest1";
  const parameters = {
    backupType: "CloudSnapshot",
    kind: "Series8000",
    retentionCount: 1,
    scheduleRecurrence: {
      recurrenceType: "Weekly",
      recurrenceValue: 1,
      weeklyDaysList: ["Friday", "Thursday", "Monday"],
    },
    scheduleStatus: "Enabled",
    startTime: new Date("2017-06-24T01:00:00Z"),
  };
  const credential = new DefaultAzureCredential();
  const client = new StorSimple8000SeriesManagementClient(credential, subscriptionId);
  const result = await client.backupSchedules.beginCreateOrUpdateAndWait(
    deviceName,
    backupPolicyName,
    backupScheduleName,
    resourceGroupName,
    managerName,
    parameters
  );
  console.log(result);
}

backupSchedulesCreateOrUpdate().catch(console.error);
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-storsimple8000series_2.0.1/sdk/storsimple8000series/arm-storsimple8000series/README.md) on how to add the SDK to your project and authenticate.
