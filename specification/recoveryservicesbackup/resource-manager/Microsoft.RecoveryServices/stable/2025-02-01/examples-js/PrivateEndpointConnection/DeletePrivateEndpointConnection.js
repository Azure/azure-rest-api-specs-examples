const { RecoveryServicesBackupClient } = require("@azure/arm-recoveryservicesbackup");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Delete Private Endpoint requests. This call is made by Backup Admin.
 *
 * @summary Delete Private Endpoint requests. This call is made by Backup Admin.
 * x-ms-original-file: specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2025-02-01/examples/PrivateEndpointConnection/DeletePrivateEndpointConnection.json
 */
async function deletePrivateEndpointConnection() {
  const subscriptionId =
    process.env["RECOVERYSERVICESBACKUP_SUBSCRIPTION_ID"] || "04cf684a-d41f-4550-9f70-7708a3a2283b";
  const vaultName = "gaallavaultbvtd2msi";
  const resourceGroupName = process.env["RECOVERYSERVICESBACKUP_RESOURCE_GROUP"] || "gaallaRG";
  const privateEndpointConnectionName = "gaallatestpe2.5704c932-249a-490b-a142-1396838cd3b";
  const credential = new DefaultAzureCredential();
  const client = new RecoveryServicesBackupClient(credential, subscriptionId);
  const result = await client.privateEndpointConnectionOperations.beginDeleteAndWait(
    vaultName,
    resourceGroupName,
    privateEndpointConnectionName,
  );
  console.log(result);
}
