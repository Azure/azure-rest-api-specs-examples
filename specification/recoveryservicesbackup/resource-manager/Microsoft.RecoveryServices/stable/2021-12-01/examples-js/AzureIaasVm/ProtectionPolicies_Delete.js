const { RecoveryServicesBackupClient } = require("@azure/arm-recoveryservicesbackup");
const { DefaultAzureCredential } = require("@azure/identity");

async function deleteAzureVMProtectionPolicy() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const vaultName = "NetSDKTestRsVault";
  const resourceGroupName = "SwaggerTestRg";
  const policyName = "testPolicy1";
  const credential = new DefaultAzureCredential();
  const client = new RecoveryServicesBackupClient(credential, subscriptionId);
  const result = await client.protectionPolicies.beginDeleteAndWait(
    vaultName,
    resourceGroupName,
    policyName
  );
  console.log(result);
}

deleteAzureVMProtectionPolicy().catch(console.error);
