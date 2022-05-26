Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-recoveryservicesbackup_9.0.0/sdk/recoveryservicesbackup/arm-recoveryservicesbackup/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { RecoveryServicesBackupClient } = require("@azure/arm-recoveryservicesbackup");
const { DefaultAzureCredential } = require("@azure/identity");

async function createOrUpdateEnhancedAzureVMProtectionPolicyWithHourlyBackup() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const vaultName = "NetSDKTestRsVault";
  const resourceGroupName = "SwaggerTestRg";
  const policyName = "v2-daily-sample";
  const parameters = {
    properties: {
      backupManagementType: "AzureIaasVM",
      instantRpRetentionRangeInDays: 30,
      policyType: "V2",
      retentionPolicy: {
        dailySchedule: {
          retentionDuration: { count: 180, durationType: "Days" },
          retentionTimes: [new Date("2021-12-17T08:00:00+00:00")],
        },
        monthlySchedule: {
          retentionDuration: { count: 60, durationType: "Months" },
          retentionScheduleDaily: {},
          retentionScheduleFormatType: "Weekly",
          retentionScheduleWeekly: {
            daysOfTheWeek: ["Sunday"],
            weeksOfTheMonth: ["First"],
          },
          retentionTimes: [new Date("2021-12-17T08:00:00+00:00")],
        },
        retentionPolicyType: "LongTermRetentionPolicy",
        weeklySchedule: {
          daysOfTheWeek: ["Sunday"],
          retentionDuration: { count: 12, durationType: "Weeks" },
          retentionTimes: [new Date("2021-12-17T08:00:00+00:00")],
        },
        yearlySchedule: {
          monthsOfYear: ["January"],
          retentionDuration: { count: 10, durationType: "Years" },
          retentionScheduleDaily: {},
          retentionScheduleFormatType: "Weekly",
          retentionScheduleWeekly: {
            daysOfTheWeek: ["Sunday"],
            weeksOfTheMonth: ["First"],
          },
          retentionTimes: [new Date("2021-12-17T08:00:00+00:00")],
        },
      },
      schedulePolicy: {
        hourlySchedule: {
          interval: 4,
          scheduleWindowDuration: 16,
          scheduleWindowStartTime: new Date("2021-12-17T08:00:00Z"),
        },
        schedulePolicyType: "SimpleSchedulePolicyV2",
        scheduleRunFrequency: "Hourly",
      },
      timeZone: "India Standard Time",
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new RecoveryServicesBackupClient(credential, subscriptionId);
  const result = await client.protectionPolicies.createOrUpdate(
    vaultName,
    resourceGroupName,
    policyName,
    parameters
  );
  console.log(result);
}

createOrUpdateEnhancedAzureVMProtectionPolicyWithHourlyBackup().catch(console.error);
```
