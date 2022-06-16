const { RecoveryServicesBackupClient } = require("@azure/arm-recoveryservicesbackup");
const { DefaultAzureCredential } = require("@azure/identity");

async function validateEnableProtectionOnAzureVM() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const azureRegion = "southeastasia";
  const parameters = {
    properties: "",
    resourceId:
      "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/arunaupgrade/providers/Microsoft.Compute/VirtualMachines/upgrade1",
    resourceType: "VM",
    vaultId:
      "/Subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRG/providers/Microsoft.RecoveryServices/Vaults/myVault",
  };
  const credential = new DefaultAzureCredential();
  const client = new RecoveryServicesBackupClient(credential, subscriptionId);
  const result = await client.protectionIntentOperations.validate(azureRegion, parameters);
  console.log(result);
}

validateEnableProtectionOnAzureVM().catch(console.error);
