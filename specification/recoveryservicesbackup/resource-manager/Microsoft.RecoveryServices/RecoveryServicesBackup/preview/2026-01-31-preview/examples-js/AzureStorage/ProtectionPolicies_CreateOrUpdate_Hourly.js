const { RecoveryServicesBackupClient } = require("@azure/arm-recoveryservicesbackup");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates or modifies a backup policy. This is an asynchronous operation. Status of the operation can be fetched
 * using GetPolicyOperationResult API.
 *
 * @summary creates or modifies a backup policy. This is an asynchronous operation. Status of the operation can be fetched
 * using GetPolicyOperationResult API.
 * x-ms-original-file: 2026-01-31-preview/AzureStorage/ProtectionPolicies_CreateOrUpdate_Hourly.json
 */
async function createOrUpdateHourlyAzureStorageProtectionPolicy() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new RecoveryServicesBackupClient(credential, subscriptionId);
  const result = await client.protectionPolicies.createOrUpdate(
    "swaggertestvault",
    "SwaggerTestRg",
    "newPolicy2",
    {
      properties: {
        backupManagementType: "AzureStorage",
        retentionPolicy: {
          dailySchedule: { retentionDuration: { count: 5, durationType: "Days" } },
          monthlySchedule: {
            retentionDuration: { count: 60, durationType: "Months" },
            retentionScheduleFormatType: "Weekly",
            retentionScheduleWeekly: { daysOfTheWeek: ["Sunday"], weeksOfTheMonth: ["First"] },
          },
          retentionPolicyType: "LongTermRetentionPolicy",
          weeklySchedule: {
            daysOfTheWeek: ["Sunday"],
            retentionDuration: { count: 12, durationType: "Weeks" },
          },
          yearlySchedule: {
            monthsOfYear: ["January"],
            retentionDuration: { count: 10, durationType: "Years" },
            retentionScheduleFormatType: "Weekly",
            retentionScheduleWeekly: { daysOfTheWeek: ["Sunday"], weeksOfTheMonth: ["First"] },
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
    },
  );
  console.log(result);
}
