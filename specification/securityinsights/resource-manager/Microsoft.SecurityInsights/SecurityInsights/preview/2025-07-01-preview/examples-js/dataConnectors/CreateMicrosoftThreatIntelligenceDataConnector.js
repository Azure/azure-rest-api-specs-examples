const { SecurityInsights } = require("@azure/arm-securityinsight");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates or updates the data connector.
 *
 * @summary creates or updates the data connector.
 * x-ms-original-file: 2025-07-01-preview/dataConnectors/CreateMicrosoftThreatIntelligenceDataConnector.json
 */
async function createsOrUpdatesAMicrosoftThreatIntelligenceDataConnector() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "d0cfe6b2-9ac0-4464-9919-dccaee2e48c0";
  const client = new SecurityInsights(credential, subscriptionId);
  const result = await client.dataConnectors.createOrUpdate(
    "myRg",
    "myWorkspace",
    "c345bf40-8509-4ed2-b947-50cb773aaf04",
    {
      kind: "MicrosoftThreatIntelligence",
      dataTypes: {
        microsoftEmergingThreatFeed: {
          lookbackPeriod: new Date("1970-01-01T00:00:00.000Z"),
          state: "Enabled",
        },
      },
      tenantId: "06b3ccb8-1384-4bcc-aec7-852f6d57161b",
    },
  );
  console.log(result);
}
