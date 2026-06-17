const { SecurityInsights } = require("@azure/arm-securityinsight");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates or updates the data connector.
 *
 * @summary creates or updates the data connector.
 * x-ms-original-file: 2025-07-01-preview/dataConnectors/CreatePremiumMicrosoftDefenderForThreatIntelligenceDataConnector.json
 */
async function createsOrUpdatesAPremiumMicrosoftDefenderForThreatIntelligenceDataConnector() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "b66e5c69-e2eb-422a-81c3-002de57059f3";
  const client = new SecurityInsights(credential, subscriptionId);
  const result = await client.dataConnectors.createOrUpdate(
    "myRg",
    "myWorkspace",
    "8c569548-a86c-4fb4-8ae4-d1e35a6146f8",
    {
      kind: "PremiumMicrosoftDefenderForThreatIntelligence",
      dataTypes: { connector: { state: "Enabled" } },
      lookbackPeriod: new Date("1970-01-01T00:00:00.000Z"),
      tenantId: "e4afb3c4-813b-4e68-b6de-e5360866e798",
    },
  );
  console.log(result);
}
