const { RecoveryServicesBackupClient } = require("@azure/arm-recoveryservicesbackup");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates or modifies a backup policy. This is an asynchronous operation. Status of the operation can be fetched
 * using GetPolicyOperationResult API.
 *
 * @summary creates or modifies a backup policy. This is an asynchronous operation. Status of the operation can be fetched
 * using GetPolicyOperationResult API.
 * x-ms-original-file: 2026-01-31-preview/AzureIaasVm/ProtectionPolicies_CreateOrUpdate_Simple.json
 */
async function createOrUpdateSimpleAzureVmProtectionPolicy() {
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
          dailySchedule: {
            retentionDuration: { count: 1, durationType: "Days" },
            retentionTimes: [new Date("2018-01-24T02:00:00Z")],
          },
          retentionPolicyType: "LongTermRetentionPolicy",
        },
        schedulePolicy: {
          schedulePolicyType: "SimpleSchedulePolicy",
          scheduleRunFrequency: "Daily",
          scheduleRunTimes: [new Date("2018-01-24T02:00:00Z")],
        },
        timeZone: "Pacific Standard Time",
      },
    },
  );
  console.log(result);
}
