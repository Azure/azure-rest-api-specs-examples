const { RecoveryServicesBackupClient } = require("@azure/arm-recoveryservicesbackup");
const { DefaultAzureCredential } = require("@azure/identity");

async function listProtectionPoliciesWithBackupManagementTypeFilterAsAzureWorkload() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const vaultName = "NetSDKTestRsVault";
  const resourceGroupName = "SwaggerTestRg";
  const filter = "backupManagementType eq 'AzureWorkload'";
  const options = { filter: filter };
  const credential = new DefaultAzureCredential();
  const client = new RecoveryServicesBackupClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.backupPolicies.list(vaultName, resourceGroupName, options)) {
    resArray.push(item);
  }
  console.log(resArray);
}

listProtectionPoliciesWithBackupManagementTypeFilterAsAzureWorkload().catch(console.error);
