const { RecoveryServicesClient } = require("@azure/arm-recoveryservices");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Unregisters the given container from your Recovery Services vault.
 *
 * @summary Unregisters the given container from your Recovery Services vault.
 * x-ms-original-file: specification/recoveryservices/resource-manager/Microsoft.RecoveryServices/stable/2022-04-01/examples/DeleteRegisteredIdentities.json
 */
async function deleteRegisteredIdentity() {
  const subscriptionId = "77777777-d41f-4550-9f70-7708a3a2283b";
  const resourceGroupName = "BCDRIbzRG";
  const vaultName = "BCDRIbzVault";
  const identityName = "dpmcontainer01";
  const credential = new DefaultAzureCredential();
  const client = new RecoveryServicesClient(credential, subscriptionId);
  const result = await client.registeredIdentities.delete(
    resourceGroupName,
    vaultName,
    identityName
  );
  console.log(result);
}

deleteRegisteredIdentity().catch(console.error);
