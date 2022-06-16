const { RecoveryServicesBackupClient } = require("@azure/arm-recoveryservicesbackup");
const { DefaultAzureCredential } = require("@azure/identity");

async function getAzureIaasVMEnhancedProtectionPolicyDetails() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const vaultName = "NetSDKTestRsVault";
  const resourceGroupName = "SwaggerTestRg";
  const policyName = "v2-daily-sample";
  const credential = new DefaultAzureCredential();
  const client = new RecoveryServicesBackupClient(credential, subscriptionId);
  const result = await client.protectionPolicies.get(vaultName, resourceGroupName, policyName);
  console.log(result);
}

getAzureIaasVMEnhancedProtectionPolicyDetails().catch(console.error);
