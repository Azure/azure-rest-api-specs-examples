const { AdvisorManagementClient } = require("@azure/arm-advisor");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to obtains details of a cached recommendation.
 *
 * @summary obtains details of a cached recommendation.
 * x-ms-original-file: 2026-02-01-preview/GetRecommendationDetailServiceGroupResourceUri.json
 */
async function getRecommendationDetailServiceGroupResourceUri() {
  const credential = new DefaultAzureCredential();
  const client = new AdvisorManagementClient(credential);
  const result = await client.recommendations.get(
    "providers/microsoft.management/serviceGroup/serviceGroupXYZ",
    "37c93209-4bfb-4f3b-8874-ccc718f7a467",
  );
  console.log(result);
}
