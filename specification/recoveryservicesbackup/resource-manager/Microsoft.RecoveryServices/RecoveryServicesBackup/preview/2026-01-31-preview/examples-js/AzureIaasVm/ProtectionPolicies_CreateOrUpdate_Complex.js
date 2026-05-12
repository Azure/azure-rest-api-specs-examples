const { RecoveryServicesBackupClient } = require("@azure/arm-recoveryservicesbackup");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates or modifies a backup policy. This is an asynchronous operation. Status of the operation can be fetched
 * using GetPolicyOperationResult API.
 *
 * @summary creates or modifies a backup policy. This is an asynchronous operation. Status of the operation can be fetched
 * using GetPolicyOperationResult API.
 * x-ms-original-file: 2026-01-31-preview/AzureIaasVm/ProtectionPolicies_CreateOrUpdate_Complex.json
 */
async function createOrUpdateFullAzureVmProtectionPolicy() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new RecoveryServicesBackupClient(credential, subscriptionId);
  const result = await client.protectionPolicies.createOrUpdate(
    "NetSDKTestRsVault",
    "SwaggerTestRg",
    "testPolicy1",
    {
      properties: {
        backupManagementType: "AzureIaasVM",
        retentionPolicy: {
          monthlySchedule: {
            retentionDuration: { count: 2, durationType: "Months" },
            retentionScheduleFormatType: "Weekly",
            retentionScheduleWeekly: {
              daysOfTheWeek: ["Wednesday", "Thursday"],
              weeksOfTheMonth: ["First", "Third"],
            },
            retentionTimes: [new Date("2018-01-24T10:00:00Z")],
          },
          retentionPolicyType: "LongTermRetentionPolicy",
          weeklySchedule: {
            daysOfTheWeek: ["Monday", "Wednesday", "Thursday"],
            retentionDuration: { count: 1, durationType: "Weeks" },
            retentionTimes: [new Date("2018-01-24T10:00:00Z")],
          },
          yearlySchedule: {
            monthsOfYear: ["February", "November"],
            retentionDuration: { count: 4, durationType: "Years" },
            retentionScheduleFormatType: "Weekly",
            retentionScheduleWeekly: {
              daysOfTheWeek: ["Monday", "Thursday"],
              weeksOfTheMonth: ["Fourth"],
            },
            retentionTimes: [new Date("2018-01-24T10:00:00Z")],
          },
        },
        schedulePolicy: {
          schedulePolicyType: "SimpleSchedulePolicy",
          scheduleRunDays: ["Monday", "Wednesday", "Thursday"],
          scheduleRunFrequency: "Weekly",
          scheduleRunTimes: [new Date("2018-01-24T10:00:00Z")],
        },
        timeZone: "Pacific Standard Time",
      },
    },
  );
  console.log(result);
}
