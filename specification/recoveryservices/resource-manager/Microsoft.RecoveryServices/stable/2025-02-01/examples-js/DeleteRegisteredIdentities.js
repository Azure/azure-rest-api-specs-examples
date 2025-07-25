const { RecoveryServicesClient } = require("@azure/arm-recoveryservices");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Unregisters the given container from your Recovery Services vault.
 *
 * @summary Unregisters the given container from your Recovery Services vault.
 * x-ms-original-file: specification/recoveryservices/resource-manager/Microsoft.RecoveryServices/stable/2025-02-01/examples/DeleteRegisteredIdentities.json
 */
async function deleteRegisteredIdentity() {
  const subscriptionId =
    process.env["RECOVERYSERVICES_SUBSCRIPTION_ID"] || "77777777-d41f-4550-9f70-7708a3a2283b";
  const resourceGroupName = process.env["RECOVERYSERVICES_RESOURCE_GROUP"] || "BCDRIbzRG";
  const vaultName = "BCDRIbzVault";
  const identityName = "dpmcontainer01";
  const credential = new DefaultAzureCredential();
  const client = new RecoveryServicesClient(credential, subscriptionId);
  const result = await client.registeredIdentities.delete(
    resourceGroupName,
    vaultName,
    identityName,
  );
  console.log(result);
}
