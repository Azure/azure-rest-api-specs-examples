const { RecoveryServicesBackupClient } = require("@azure/arm-recoveryservicesbackup");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or modifies a backup policy. This is an asynchronous operation. Status of the operation can be fetched
using GetPolicyOperationResult API.
 *
 * @summary Creates or modifies a backup policy. This is an asynchronous operation. Status of the operation can be fetched
using GetPolicyOperationResult API.
 * x-ms-original-file: specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2023-06-01/examples/AzureIaasVm/ProtectionPolicies_CreateOrUpdate_Complex.json
 */
async function createOrUpdateFullAzureVMProtectionPolicy() {
  const subscriptionId =
    process.env["RECOVERYSERVICESBACKUP_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const vaultName = "NetSDKTestRsVault";
  const resourceGroupName = process.env["RECOVERYSERVICESBACKUP_RESOURCE_GROUP"] || "SwaggerTestRg";
  const policyName = "testPolicy1";
  const parameters = {
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
