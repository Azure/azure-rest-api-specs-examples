const { RecoveryServicesBackupClient } = require("@azure/arm-recoveryservicesbackup");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or modifies a backup policy. This is an asynchronous operation. Status of the operation can be fetched
using GetPolicyOperationResult API.
 *
 * @summary Creates or modifies a backup policy. This is an asynchronous operation. Status of the operation can be fetched
using GetPolicyOperationResult API.
 * x-ms-original-file: specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2023-06-01/examples/AzureIaasVm/V2Policy/IaaS_v2_daily.json
 */
async function createOrUpdateEnhancedAzureVMProtectionPolicyWithDailyBackup() {
  const subscriptionId =
    process.env["RECOVERYSERVICESBACKUP_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const vaultName = "NetSDKTestRsVault";
  const resourceGroupName = process.env["RECOVERYSERVICESBACKUP_RESOURCE_GROUP"] || "SwaggerTestRg";
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
        dailySchedule: { scheduleRunTimes: [new Date("2018-01-24T10:00:00Z")] },
        schedulePolicyType: "SimpleSchedulePolicyV2",
        scheduleRunFrequency: "Daily",
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
    parameters,
  );
  console.log(result);
}
