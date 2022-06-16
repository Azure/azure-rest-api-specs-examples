const { RecoveryServicesBackupClient } = require("@azure/arm-recoveryservicesbackup");
const { DefaultAzureCredential } = require("@azure/identity");

async function createOrUpdateHourlyAzureStorageProtectionPolicy() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const vaultName = "swaggertestvault";
  const resourceGroupName = "SwaggerTestRg";
  const policyName = "newPolicy2";
  const parameters = {
    properties: {
      backupManagementType: "AzureStorage",
      retentionPolicy: {
        dailySchedule: {
          retentionDuration: { count: 5, durationType: "Days" },
          retentionTimes: [],
        },
        monthlySchedule: {
          retentionDuration: { count: 60, durationType: "Months" },
          retentionScheduleDaily: {},
          retentionScheduleFormatType: "Weekly",
          retentionScheduleWeekly: {
            daysOfTheWeek: ["Sunday"],
            weeksOfTheMonth: ["First"],
          },
          retentionTimes: [],
        },
        retentionPolicyType: "LongTermRetentionPolicy",
        weeklySchedule: {
          daysOfTheWeek: ["Sunday"],
          retentionDuration: { count: 12, durationType: "Weeks" },
          retentionTimes: [],
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
          retentionTimes: [],
        },
      },
      schedulePolicy: {
        hourlySchedule: {
          interval: 4,
          scheduleWindowDuration: 12,
          scheduleWindowStartTime: new Date("2021-09-29T08:00:00.000Z"),
        },
        schedulePolicyType: "SimpleSchedulePolicy",
        scheduleRunFrequency: "Hourly",
      },
      timeZone: "UTC",
      workLoadType: "AzureFileShare",
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

createOrUpdateHourlyAzureStorageProtectionPolicy().catch(console.error);
