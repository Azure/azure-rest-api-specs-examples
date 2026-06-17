const { SecurityInsights } = require("@azure/arm-securityinsight");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates or updates the data connector.
 *
 * @summary creates or updates the data connector.
 * x-ms-original-file: 2025-07-01-preview/dataConnectors/CreateMicrosoftThreatProtectionDataConnetor.json
 */
async function createsOrUpdatesAMicrosoftThreatProtectionDataConnector() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "d0cfe6b2-9ac0-4464-9919-dccaee2e48c0";
  const client = new SecurityInsights(credential, subscriptionId);
  const result = await client.dataConnectors.createOrUpdate(
    "myRg",
    "myWorkspace",
    "73e01a99-5cd7-4139-a149-9f2736ff2ab5",
    {
      etag: '"0300bf09-0000-0000-0000-5c37296e0000"',
      kind: "MicrosoftThreatProtection",
      dataTypes: { alerts: { state: "Enabled" }, incidents: { state: "Disabled" } },
      filteredProviders: { alerts: ["microsoftDefenderForCloudApps"] },
      tenantId: "178265c4-3136-4ff6-8ed1-b5b62b4cb5f5",
    },
  );
  console.log(result);
}
