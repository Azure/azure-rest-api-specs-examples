const { SecurityInsights } = require("@azure/arm-securityinsight");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to gets a data connector.
 *
 * @summary gets a data connector.
 * x-ms-original-file: 2025-07-01-preview/dataConnectors/GetPremiumMicrosoftDefenderForThreatIntelligenceById.json
 */
async function getAPremiumMicrosoftDefenderForThreatIntelligenceDataConnector() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "b66e5c69-e2eb-422a-81c3-002de57059f3";
  const client = new SecurityInsights(credential, subscriptionId);
  const result = await client.dataConnectors.get(
    "myRg",
    "myWorkspace",
    "8c569548-a86c-4fb4-8ae4-d1e35a6146f8",
  );
  console.log(result);
}
