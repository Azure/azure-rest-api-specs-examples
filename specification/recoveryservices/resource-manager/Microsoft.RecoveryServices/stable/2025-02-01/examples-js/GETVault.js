const { RecoveryServicesClient } = require("@azure/arm-recoveryservices");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Get the Vault details.
 *
 * @summary Get the Vault details.
 * x-ms-original-file: specification/recoveryservices/resource-manager/Microsoft.RecoveryServices/stable/2025-02-01/examples/GETVault.json
 */
async function getRecoveryServicesResource() {
  const subscriptionId =
    process.env["RECOVERYSERVICES_SUBSCRIPTION_ID"] || "77777777-b0c6-47a2-b37c-d8e65a629c18";
  const resourceGroupName =
    process.env["RECOVERYSERVICES_RESOURCE_GROUP"] || "Default-RecoveryServices-ResourceGroup";
  const vaultName = "swaggerExample";
  const credential = new DefaultAzureCredential();
  const client = new RecoveryServicesClient(credential, subscriptionId);
  const result = await client.vaults.get(resourceGroupName, vaultName);
  console.log(result);
}
