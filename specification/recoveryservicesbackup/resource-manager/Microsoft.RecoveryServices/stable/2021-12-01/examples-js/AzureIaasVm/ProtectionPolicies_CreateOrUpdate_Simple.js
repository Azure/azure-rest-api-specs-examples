const { RecoveryServicesBackupClient } = require("@azure/arm-recoveryservicesbackup");
const { DefaultAzureCredential } = require("@azure/identity");

async function createOrUpdateSimpleAzureVMProtectionPolicy() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const vaultName = "NetSDKTestRsVault";
  const resourceGroupName = "SwaggerTestRg";
  const policyName = "testPolicy1";
  const parameters = {
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

createOrUpdateSimpleAzureVMProtectionPolicy().catch(console.error);
