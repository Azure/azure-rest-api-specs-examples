const { RecoveryServicesBackupClient } = require("@azure/arm-recoveryservicesbackup");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Fetches the result of any operation on the container.
 *
 * @summary Fetches the result of any operation on the container.
 * x-ms-original-file: specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2025-02-01/examples/AzureStorage/ProtectionContainers_Inquire_Result.json
 */
async function getAzureStorageProtectionContainerOperationResult() {
  const subscriptionId =
    process.env["RECOVERYSERVICESBACKUP_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const vaultName = "testvault";
  const resourceGroupName = process.env["RECOVERYSERVICESBACKUP_RESOURCE_GROUP"] || "test-rg";
  const fabricName = "Azure";
  const containerName = "VMAppContainer;Compute;testRG;testSQL";
  const operationId = "00000000-0000-0000-0000-000000000000";
  const credential = new DefaultAzureCredential();
  const client = new RecoveryServicesBackupClient(credential, subscriptionId);
  const result = await client.protectionContainerOperationResults.get(
    vaultName,
    resourceGroupName,
    fabricName,
    containerName,
    operationId,
  );
  console.log(result);
}
