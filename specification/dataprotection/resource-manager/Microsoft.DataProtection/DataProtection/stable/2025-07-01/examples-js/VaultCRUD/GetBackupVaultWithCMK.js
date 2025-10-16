const { DataProtectionClient } = require("@azure/arm-dataprotection");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to returns a resource belonging to a resource group.
 *
 * @summary returns a resource belonging to a resource group.
 * x-ms-original-file: 2025-07-01/VaultCRUD/GetBackupVaultWithCMK.json
 */
async function getBackupVaultWithCMK() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "0b352192-dcac-4cc7-992e-a96190ccc68c";
  const client = new DataProtectionClient(credential, subscriptionId);
  const result = await client.backupVaults.get("SampleResourceGroup", "swaggerExample");
  console.log(result);
}
