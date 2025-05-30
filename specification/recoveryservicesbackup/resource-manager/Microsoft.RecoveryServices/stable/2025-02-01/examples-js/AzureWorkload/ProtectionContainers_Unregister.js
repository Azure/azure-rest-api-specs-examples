const { RecoveryServicesBackupClient } = require("@azure/arm-recoveryservicesbackup");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Unregisters the given container from your Recovery Services Vault. This is an asynchronous operation. To determine
whether the backend service has finished processing the request, call Get Container Operation Result API.
 *
 * @summary Unregisters the given container from your Recovery Services Vault. This is an asynchronous operation. To determine
whether the backend service has finished processing the request, call Get Container Operation Result API.
 * x-ms-original-file: specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2025-02-01/examples/AzureWorkload/ProtectionContainers_Unregister.json
 */
async function unregisterProtectionContainer() {
  const subscriptionId =
    process.env["RECOVERYSERVICESBACKUP_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const vaultName = "testVault";
  const resourceGroupName = process.env["RECOVERYSERVICESBACKUP_RESOURCE_GROUP"] || "testRg";
  const fabricName = "Azure";
  const containerName = "storagecontainer;Storage;test-rg;teststorage";
  const credential = new DefaultAzureCredential();
  const client = new RecoveryServicesBackupClient(credential, subscriptionId);
  const result = await client.protectionContainers.unregister(
    vaultName,
    resourceGroupName,
    fabricName,
    containerName,
  );
  console.log(result);
}
