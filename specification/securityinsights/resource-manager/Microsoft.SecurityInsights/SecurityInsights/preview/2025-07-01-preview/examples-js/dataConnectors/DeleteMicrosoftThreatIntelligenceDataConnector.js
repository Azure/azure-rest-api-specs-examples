const { SecurityInsights } = require("@azure/arm-securityinsight");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to delete the data connector.
 *
 * @summary delete the data connector.
 * x-ms-original-file: 2025-07-01-preview/dataConnectors/DeleteMicrosoftThreatIntelligenceDataConnector.json
 */
async function deleteAnMicrosoftThreatIntelligenceDataConnector() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "d0cfe6b2-9ac0-4464-9919-dccaee2e48c0";
  const client = new SecurityInsights(credential, subscriptionId);
  await client.dataConnectors.delete("myRg", "myWorkspace", "c345bf40-8509-4ed2-b947-50cb773aaf04");
}
