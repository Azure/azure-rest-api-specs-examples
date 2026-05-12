const { RecoveryServicesBackupClient } = require("@azure/arm-recoveryservicesbackup");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates or modifies a backup policy. This is an asynchronous operation. Status of the operation can be fetched
 * using GetPolicyOperationResult API.
 *
 * @summary creates or modifies a backup policy. This is an asynchronous operation. Status of the operation can be fetched
 * using GetPolicyOperationResult API.
 * x-ms-original-file: 2026-01-31-preview/AzureWorkload/ProtectionPolicies_CreateOrUpdate_SapHanaDBInstance.json
 */
async function createOrUpdateSapHanaDBInstanceWorkloadProtectionPolicy() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new RecoveryServicesBackupClient(credential, subscriptionId);
  const result = await client.protectionPolicies.createOrUpdate(
    "HanaTestRsVault",
    "SwaggerTestRg",
    "testHanaSnapshotV2Policy1",
    {
      properties: {
        backupManagementType: "AzureWorkload",
        workLoadType: "SAPHanaDBInstance",
        vmWorkloadPolicyType: "SnapshotV2",
        settings: { timeZone: "UTC", issqlcompression: false, isCompression: false },
        subProtectionPolicy: [
          {
            policyType: "SnapshotFull",
            schedulePolicy: {
              schedulePolicyType: "SimpleSchedulePolicy",
              scheduleRunFrequency: "Daily",
              scheduleRunTimes: [new Date("2024-10-01T03:30:00.000Z")],
            },
            retentionPolicy: {
              retentionPolicyType: "LongTermRetentionPolicy",
              dailySchedule: {
                retentionTimes: [new Date("2023-12-19T20:00:00.000Z")],
                retentionDuration: { count: 30, durationType: "Days" },
              },
              weeklySchedule: {
                daysOfTheWeek: ["Sunday"],
                retentionTimes: [new Date("2023-12-19T20:00:00.000Z")],
                retentionDuration: { count: 10, durationType: "Weeks" },
              },
              monthlySchedule: {
                retentionScheduleFormatType: "Weekly",
                retentionScheduleWeekly: { daysOfTheWeek: ["Sunday"], weeksOfTheMonth: ["Second"] },
                retentionTimes: [new Date("2023-12-15T20:00:00.000Z")],
                retentionDuration: { count: 6, durationType: "Months" },
              },
              yearlySchedule: {
                retentionScheduleFormatType: "Weekly",
                monthsOfYear: ["January"],
                retentionScheduleWeekly: { daysOfTheWeek: ["Sunday"], weeksOfTheMonth: ["Last"] },
                retentionTimes: [new Date("2023-12-19T20:00:00.000Z")],
                retentionDuration: { count: 2, durationType: "Years" },
              },
            },
            snapshotBackupAdditionalDetails: {
              instantRpRetentionRangeInDays: 5,
              instantRPDetails: "SwaggerTestRG",
              userAssignedManagedIdentityDetails: {
                identityArmId:
                  "/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/SwaggerMsiRG/providers/Microsoft.ManagedIdentity/userAssignedIdentities/SwaggerUMI",
                identityName: "SwaggerUMI",
                userAssignedIdentityProperties: {
                  clientId: "00000000-0000-0000-0000-000000000000",
                  principalId: "00000000-0000-0000-0000-000000000000",
                },
              },
            },
          },
        ],
        protectedItemsCount: 0,
      },
    },
  );
  console.log(result);
}
