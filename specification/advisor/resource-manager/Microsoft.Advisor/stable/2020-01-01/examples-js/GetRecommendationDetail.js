const { AdvisorManagementClient } = require("@azure/arm-advisor");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Obtains details of a cached recommendation.
 *
 * @summary Obtains details of a cached recommendation.
 * x-ms-original-file: specification/advisor/resource-manager/Microsoft.Advisor/stable/2020-01-01/examples/GetRecommendationDetail.json
 */
async function getRecommendationDetail() {
  const resourceUri = "resourceUri";
  const recommendationId = "recommendationId";
  const credential = new DefaultAzureCredential();
  const client = new AdvisorManagementClient(credential);
  const result = await client.recommendations.get(resourceUri, recommendationId);
  console.log(result);
}
