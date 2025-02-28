const { ImpactClient } = require("@azure/arm-impactreporting");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to get Insight resources by workloadImpactName and insightName
 *
 * @summary get Insight resources by workloadImpactName and insightName
 * x-ms-original-file: 2024-05-01-preview/Insights_Get_diagnostics.json
 */
async function getInsightSampleForDiagnosticsCategory() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new ImpactClient(credential, subscriptionId);
  const result = await client.insights.get("impactid", "insight1");
  console.log(result);
}
