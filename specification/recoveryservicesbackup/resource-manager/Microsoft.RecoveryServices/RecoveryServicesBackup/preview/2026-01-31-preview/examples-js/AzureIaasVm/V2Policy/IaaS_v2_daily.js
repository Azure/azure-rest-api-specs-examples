const { RecoveryServicesBackupClient } = require("@azure/arm-recoveryservicesbackup");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates or modifies a backup policy. This is an asynchronous operation. Status of the operation can be fetched
 * using GetPolicyOperationResult API.
 *
 * @summary creates or modifies a backup policy. This is an asynchronous operation. Status of the operation can be fetched
 * using GetPolicyOperationResult API.
 * x-ms-original-file: 2026-01-31-preview/AzureIaasVm/V2Policy/IaaS_v2_daily.json
 */
async function createOrUpdateEnhancedAzureVmProtectionPolicyWithDailyBackup() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new RecoveryServicesBackupClient(credential, subscriptionId);
  const result = await client.protectionPolicies.createOrUpdate(
    "NetSDKTestRsVault",
    "SwaggerTestRg",
    "v2-daily-sample",
    {
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
            retentionScheduleFormatType: "Weekly",
            retentionScheduleWeekly: { daysOfTheWeek: ["Sunday"], weeksOfTheMonth: ["First"] },
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
            retentionScheduleFormatType: "Weekly",
            retentionScheduleWeekly: { daysOfTheWeek: ["Sunday"], weeksOfTheMonth: ["First"] },
            retentionTimes: [new Date("2021-12-17T08:00:00+00:00")],
          },
        },
        schedulePolicy: {
          dailySchedule: { scheduleRunTimes: [new Date("2018-01-24T10:00:00Z")] },
          schedulePolicyType: "SimpleSchedulePolicyV2",
          scheduleRunFrequency: "Daily",
        },
        snapshotConsistencyType: "OnlyCrashConsistent",
        timeZone: "India Standard Time",
      },
    },
  );
  console.log(result);
}
