const { SecurityInsights } = require("@azure/arm-securityinsight");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to delete the data connector.
 *
 * @summary delete the data connector.
 * x-ms-original-file: 2025-07-01-preview/dataConnectors/DeletePremiumMicrosoftDefenderForThreatIntelligenceDataConnector.json
 */
async function deletesAPremiumMicrosoftDefenderForThreatIntelligenceDataConnector() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "b66e5c69-e2eb-422a-81c3-002de57059f3";
  const client = new SecurityInsights(credential, subscriptionId);
  await client.dataConnectors.delete("myRg", "myWorkspace", "8c569548-a86c-4fb4-8ae4-d1e35a6146f8");
}
