const { RecoveryServicesBackupClient } = require("@azure/arm-recoveryservicesbackup");
const { DefaultAzureCredential } = require("@azure/identity");

async function createOrUpdateFullAzureWorkloadProtectionPolicy() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const vaultName = "NetSDKTestRsVault";
  const resourceGroupName = "SwaggerTestRg";
  const policyName = "testPolicy1";
  const parameters = {
    properties: {
      backupManagementType: "AzureWorkload",
      settings: { issqlcompression: false, timeZone: "Pacific Standard Time" },
      subProtectionPolicy: [
        {
          policyType: "Full",
          retentionPolicy: {
            monthlySchedule: {
              retentionDuration: { count: 1, durationType: "Months" },
              retentionScheduleFormatType: "Weekly",
              retentionScheduleWeekly: {
                daysOfTheWeek: ["Sunday"],
                weeksOfTheMonth: ["Second"],
              },
              retentionTimes: [new Date("2018-01-24T10:00:00Z")],
            },
            retentionPolicyType: "LongTermRetentionPolicy",
            weeklySchedule: {
              daysOfTheWeek: ["Sunday", "Tuesday"],
              retentionDuration: { count: 2, durationType: "Weeks" },
              retentionTimes: [new Date("2018-01-24T10:00:00Z")],
            },
            yearlySchedule: {
              monthsOfYear: ["January", "June", "December"],
              retentionDuration: { count: 1, durationType: "Years" },
              retentionScheduleFormatType: "Weekly",
              retentionScheduleWeekly: {
                daysOfTheWeek: ["Sunday"],
                weeksOfTheMonth: ["Last"],
              },
              retentionTimes: [new Date("2018-01-24T10:00:00Z")],
            },
          },
          schedulePolicy: {
            schedulePolicyType: "SimpleSchedulePolicy",
            scheduleRunDays: ["Sunday", "Tuesday"],
            scheduleRunFrequency: "Weekly",
            scheduleRunTimes: [new Date("2018-01-24T10:00:00Z")],
          },
        },
        {
          policyType: "Differential",
          retentionPolicy: {
            retentionDuration: { count: 8, durationType: "Days" },
            retentionPolicyType: "SimpleRetentionPolicy",
          },
          schedulePolicy: {
            schedulePolicyType: "SimpleSchedulePolicy",
            scheduleRunDays: ["Friday"],
            scheduleRunFrequency: "Weekly",
            scheduleRunTimes: [new Date("2018-01-24T10:00:00Z")],
          },
        },
        {
          policyType: "Log",
          retentionPolicy: {
            retentionDuration: { count: 7, durationType: "Days" },
            retentionPolicyType: "SimpleRetentionPolicy",
          },
          schedulePolicy: {
            scheduleFrequencyInMins: 60,
            schedulePolicyType: "LogSchedulePolicy",
          },
        },
      ],
      workLoadType: "SQLDataBase",
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

createOrUpdateFullAzureWorkloadProtectionPolicy().catch(console.error);
