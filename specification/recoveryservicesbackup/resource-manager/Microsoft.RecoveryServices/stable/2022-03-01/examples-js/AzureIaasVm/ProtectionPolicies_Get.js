const { RecoveryServicesBackupClient } = require("@azure/arm-recoveryservicesbackup");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Provides the details of the backup policies associated to Recovery Services Vault. This is an asynchronous
operation. Status of the operation can be fetched using GetPolicyOperationResult API.
 *
 * @summary Provides the details of the backup policies associated to Recovery Services Vault. This is an asynchronous
operation. Status of the operation can be fetched using GetPolicyOperationResult API.
 * x-ms-original-file: specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2022-03-01/examples/AzureIaasVm/ProtectionPolicies_Get.json
 */
async function getAzureIaasVMProtectionPolicyDetails() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const vaultName = "NetSDKTestRsVault";
  const resourceGroupName = "SwaggerTestRg";
  const policyName = "testPolicy1";
  const credential = new DefaultAzureCredential();
  const client = new RecoveryServicesBackupClient(credential, subscriptionId);
  const result = await client.protectionPolicies.get(vaultName, resourceGroupName, policyName);
  console.log(result);
}

getAzureIaasVMProtectionPolicyDetails().catch(console.error);
