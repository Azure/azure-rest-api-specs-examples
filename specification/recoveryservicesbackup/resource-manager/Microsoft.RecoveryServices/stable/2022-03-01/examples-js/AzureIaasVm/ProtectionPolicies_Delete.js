const { RecoveryServicesBackupClient } = require("@azure/arm-recoveryservicesbackup");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes specified backup policy from your Recovery Services Vault. This is an asynchronous operation. Status of the
operation can be fetched using GetProtectionPolicyOperationResult API.
 *
 * @summary Deletes specified backup policy from your Recovery Services Vault. This is an asynchronous operation. Status of the
operation can be fetched using GetProtectionPolicyOperationResult API.
 * x-ms-original-file: specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2022-03-01/examples/AzureIaasVm/ProtectionPolicies_Delete.json
 */
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
