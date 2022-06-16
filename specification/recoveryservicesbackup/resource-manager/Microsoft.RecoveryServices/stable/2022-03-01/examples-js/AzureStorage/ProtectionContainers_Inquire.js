const { RecoveryServicesBackupClient } = require("@azure/arm-recoveryservicesbackup");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to This is an async operation and the results should be tracked using location header or Azure-async-url.
 *
 * @summary This is an async operation and the results should be tracked using location header or Azure-async-url.
 * x-ms-original-file: specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2022-03-01/examples/AzureStorage/ProtectionContainers_Inquire.json
 */
async function inquireAzureStorageProtectionContainers() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const vaultName = "testvault";
  const resourceGroupName = "test-rg";
  const fabricName = "Azure";
  const containerName = "storagecontainer;Storage;test-rg;teststorage";
  const credential = new DefaultAzureCredential();
  const client = new RecoveryServicesBackupClient(credential, subscriptionId);
  const result = await client.protectionContainers.inquire(
    vaultName,
    resourceGroupName,
    fabricName,
    containerName
  );
  console.log(result);
}

inquireAzureStorageProtectionContainers().catch(console.error);
