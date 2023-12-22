const { DataProtectionClient } = require("@azure/arm-dataprotection");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Returns a resource belonging to a resource group.
 *
 * @summary Returns a resource belonging to a resource group.
 * x-ms-original-file: specification/dataprotection/resource-manager/Microsoft.DataProtection/stable/2023-11-01/examples/VaultCRUD/GetBackupVault.json
 */
async function getBackupVault() {
  const subscriptionId =
    process.env["DATAPROTECTION_SUBSCRIPTION_ID"] || "0b352192-dcac-4cc7-992e-a96190ccc68c";
  const resourceGroupName = process.env["DATAPROTECTION_RESOURCE_GROUP"] || "SampleResourceGroup";
  const vaultName = "swaggerExample";
  const credential = new DefaultAzureCredential();
  const client = new DataProtectionClient(credential, subscriptionId);
  const result = await client.backupVaults.get(resourceGroupName, vaultName);
  console.log(result);
}
