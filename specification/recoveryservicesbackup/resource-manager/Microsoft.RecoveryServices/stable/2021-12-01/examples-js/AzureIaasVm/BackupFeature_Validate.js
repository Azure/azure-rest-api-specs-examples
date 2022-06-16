const { RecoveryServicesBackupClient } = require("@azure/arm-recoveryservicesbackup");
const { DefaultAzureCredential } = require("@azure/identity");

async function checkAzureVMBackupFeatureSupport() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const azureRegion = "southeastasia";
  const parameters = {
    featureType: "AzureVMResourceBackup",
    vmSize: "Basic_A0",
    vmSku: "Premium",
  };
  const credential = new DefaultAzureCredential();
  const client = new RecoveryServicesBackupClient(credential, subscriptionId);
  const result = await client.featureSupport.validate(azureRegion, parameters);
  console.log(result);
}

checkAzureVMBackupFeatureSupport().catch(console.error);
