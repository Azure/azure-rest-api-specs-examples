const { RecoveryServicesBackupClient } = require("@azure/arm-recoveryservicesbackup");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to It will validate followings
1. Vault capacity
2. VM is already protected
3. Any VM related configuration passed in properties.
 *
 * @summary It will validate followings
1. Vault capacity
2. VM is already protected
3. Any VM related configuration passed in properties.
 * x-ms-original-file: specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2025-02-01/examples/AzureIaasVm/ProtectionIntent_Validate.json
 */
async function validateEnableProtectionOnAzureVM() {
  const subscriptionId =
    process.env["RECOVERYSERVICESBACKUP_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
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
