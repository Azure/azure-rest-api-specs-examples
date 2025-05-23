const { RecoveryServicesBackupClient } = require("@azure/arm-recoveryservicesbackup");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Create Intent for Enabling backup of an item. This is a synchronous operation.
 *
 * @summary Create Intent for Enabling backup of an item. This is a synchronous operation.
 * x-ms-original-file: specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2025-02-01/examples/AzureIaasVm/ProtectionIntent_CreateOrUpdate.json
 */
async function createOrUpdateAzureVMProtectionIntent() {
  const subscriptionId =
    process.env["RECOVERYSERVICESBACKUP_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const vaultName = "myVault";
  const resourceGroupName = process.env["RECOVERYSERVICESBACKUP_RESOURCE_GROUP"] || "myRG";
  const fabricName = "Azure";
  const intentObjectName = "vm;iaasvmcontainerv2;chamsrgtest;chamscandel";
  const parameters = {
    properties: {
      policyId:
        "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRG/providers/Microsoft.RecoveryServices/vaults/myVault/backupPolicies/myPolicy",
      protectionIntentItemType: "AzureResourceItem",
      sourceResourceId:
        "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/chamsrgtest/providers/Microsoft.Compute/virtualMachines/chamscandel",
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new RecoveryServicesBackupClient(credential, subscriptionId);
  const result = await client.protectionIntentOperations.createOrUpdate(
    vaultName,
    resourceGroupName,
    fabricName,
    intentObjectName,
    parameters,
  );
  console.log(result);
}
