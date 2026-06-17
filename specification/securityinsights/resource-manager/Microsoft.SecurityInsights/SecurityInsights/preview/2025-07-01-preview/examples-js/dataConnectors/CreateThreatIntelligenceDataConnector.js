const { SecurityInsights } = require("@azure/arm-securityinsight");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates or updates the data connector.
 *
 * @summary creates or updates the data connector.
 * x-ms-original-file: 2025-07-01-preview/dataConnectors/CreateThreatIntelligenceDataConnector.json
 */
async function createsOrUpdatesAnThreatIntelligencePlatformDataConnector() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "d0cfe6b2-9ac0-4464-9919-dccaee2e48c0";
  const client = new SecurityInsights(credential, subscriptionId);
  const result = await client.dataConnectors.createOrUpdate(
    "myRg",
    "myWorkspace",
    "73e01a99-5cd7-4139-a149-9f2736ff2ab5",
    {
      kind: "ThreatIntelligence",
      dataTypes: { indicators: { state: "Enabled" } },
      tenantId: "06b3ccb8-1384-4bcc-aec7-852f6d57161b",
      tipLookbackPeriod: new Date("2020-01-01T13:00:30.123Z"),
    },
  );
  console.log(result);
}
