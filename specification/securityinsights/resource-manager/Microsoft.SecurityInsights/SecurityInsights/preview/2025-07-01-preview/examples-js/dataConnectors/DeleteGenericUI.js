const { SecurityInsights } = require("@azure/arm-securityinsight");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to delete the data connector.
 *
 * @summary delete the data connector.
 * x-ms-original-file: 2025-07-01-preview/dataConnectors/DeleteGenericUI.json
 */
async function deleteAGenericUIDataConnector() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "d0cfe6b2-9ac0-4464-9919-dccaee2e48c0";
  const client = new SecurityInsights(credential, subscriptionId);
  await client.dataConnectors.delete("myRg", "myWorkspace", "316ec55e-7138-4d63-ab18-90c8a60fd1c8");
}
