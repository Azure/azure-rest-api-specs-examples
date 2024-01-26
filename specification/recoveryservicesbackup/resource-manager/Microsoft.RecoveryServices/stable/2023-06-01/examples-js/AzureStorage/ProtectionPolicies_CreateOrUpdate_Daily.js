const { RecoveryServicesBackupClient } = require("@azure/arm-recoveryservicesbackup");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or modifies a backup policy. This is an asynchronous operation. Status of the operation can be fetched
using GetPolicyOperationResult API.
 *
 * @summary Creates or modifies a backup policy. This is an asynchronous operation. Status of the operation can be fetched
using GetPolicyOperationResult API.
 * x-ms-original-file: specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2023-06-01/examples/AzureStorage/ProtectionPolicies_CreateOrUpdate_Daily.json
 */
async function createOrUpdateDailyAzureStorageProtectionPolicy() {
  const subscriptionId =
    process.env["RECOVERYSERVICESBACKUP_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const vaultName = "swaggertestvault";
  const resourceGroupName = process.env["RECOVERYSERVICESBACKUP_RESOURCE_GROUP"] || "SwaggerTestRg";
  const policyName = "dailyPolicy2";
  const parameters = {
    properties: {
      backupManagementType: "AzureStorage",
      retentionPolicy: {
        dailySchedule: {
          retentionDuration: { count: 5, durationType: "Days" },
          retentionTimes: [new Date("2021-09-29T08:00:00.000Z")],
        },
        monthlySchedule: {
          retentionDuration: { count: 60, durationType: "Months" },
          retentionScheduleDaily: {},
          retentionScheduleFormatType: "Weekly",
          retentionScheduleWeekly: {
            daysOfTheWeek: ["Sunday"],
            weeksOfTheMonth: ["First"],
          },
          retentionTimes: [new Date("2021-09-29T08:00:00.000Z")],
        },
        retentionPolicyType: "LongTermRetentionPolicy",
        weeklySchedule: {
          daysOfTheWeek: ["Sunday"],
          retentionDuration: { count: 12, durationType: "Weeks" },
          retentionTimes: [new Date("2021-09-29T08:00:00.000Z")],
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
          retentionTimes: [new Date("2021-09-29T08:00:00.000Z")],
        },
      },
      schedulePolicy: {
        schedulePolicyType: "SimpleSchedulePolicy",
        scheduleRunFrequency: "Daily",
        scheduleRunTimes: [new Date("2021-09-29T08:00:00.000Z")],
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
    parameters,
  );
  console.log(result);
}
