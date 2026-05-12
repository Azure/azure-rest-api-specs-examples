const { RecoveryServicesBackupClient } = require("@azure/arm-recoveryservicesbackup");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates or modifies a backup policy. This is an asynchronous operation. Status of the operation can be fetched
 * using GetPolicyOperationResult API.
 *
 * @summary creates or modifies a backup policy. This is an asynchronous operation. Status of the operation can be fetched
 * using GetPolicyOperationResult API.
 * x-ms-original-file: 2026-01-31-preview/AzureStorage/ProtectionPolicies_CreateOrUpdate_Hardened.json
 */
async function createOrUpdateAzureStorageVaultStandardProtectionPolicy() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new RecoveryServicesBackupClient(credential, subscriptionId);
  const result = await client.protectionPolicies.createOrUpdate(
    "swaggertestvault",
    "SwaggerTestRg",
    "newPolicyV2",
    {
      properties: {
        backupManagementType: "AzureStorage",
        schedulePolicy: {
          schedulePolicyType: "SimpleSchedulePolicy",
          scheduleRunFrequency: "Daily",
          scheduleRunTimes: [new Date("2023-07-18T09:30:00.000Z")],
        },
        timeZone: "UTC",
        vaultRetentionPolicy: {
          snapshotRetentionInDays: 5,
          vaultRetention: {
            dailySchedule: {
              retentionDuration: { count: 30, durationType: "Days" },
              retentionTimes: [new Date("2023-07-18T09:30:00.000Z")],
            },
            monthlySchedule: {
              retentionDuration: { count: 60, durationType: "Months" },
              retentionScheduleFormatType: "Weekly",
              retentionScheduleWeekly: { daysOfTheWeek: ["Sunday"], weeksOfTheMonth: ["First"] },
              retentionTimes: [new Date("2023-07-18T09:30:00.000Z")],
            },
            retentionPolicyType: "LongTermRetentionPolicy",
            weeklySchedule: {
              daysOfTheWeek: ["Sunday"],
              retentionDuration: { count: 12, durationType: "Weeks" },
              retentionTimes: [new Date("2023-07-18T09:30:00.000Z")],
            },
            yearlySchedule: {
              monthsOfYear: ["January"],
              retentionDuration: { count: 10, durationType: "Years" },
              retentionScheduleFormatType: "Weekly",
              retentionScheduleWeekly: { daysOfTheWeek: ["Sunday"], weeksOfTheMonth: ["First"] },
              retentionTimes: [new Date("2023-07-18T09:30:00.000Z")],
            },
          },
        },
        workLoadType: "AzureFileShare",
      },
    },
  );
  console.log(result);
}
