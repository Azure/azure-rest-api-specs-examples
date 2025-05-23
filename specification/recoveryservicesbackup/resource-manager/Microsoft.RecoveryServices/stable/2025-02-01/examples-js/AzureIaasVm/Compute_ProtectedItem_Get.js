const { RecoveryServicesBackupClient } = require("@azure/arm-recoveryservicesbackup");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Provides the details of the backed up item. This is an asynchronous operation. To know the status of the operation,
call the GetItemOperationResult API.
 *
 * @summary Provides the details of the backed up item. This is an asynchronous operation. To know the status of the operation,
call the GetItemOperationResult API.
 * x-ms-original-file: specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2025-02-01/examples/AzureIaasVm/Compute_ProtectedItem_Get.json
 */
async function getProtectedVirtualMachineDetails() {
  const subscriptionId =
    process.env["RECOVERYSERVICESBACKUP_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const vaultName = "PySDKBackupTestRsVault";
  const resourceGroupName =
    process.env["RECOVERYSERVICESBACKUP_RESOURCE_GROUP"] || "PythonSDKBackupTestRg";
  const fabricName = "Azure";
  const containerName = "iaasvmcontainer;iaasvmcontainerv2;iaasvm-rg;iaasvm-1";
  const protectedItemName = "vm;iaasvmcontainerv2;iaasvm-rg;iaasvm-1";
  const credential = new DefaultAzureCredential();
  const client = new RecoveryServicesBackupClient(credential, subscriptionId);
  const result = await client.protectedItems.get(
    vaultName,
    resourceGroupName,
    fabricName,
    containerName,
    protectedItemName,
  );
  console.log(result);
}
