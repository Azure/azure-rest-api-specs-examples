const { CustomerInsightsManagementClient } = require("@azure/arm-customerinsights");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets information about the specified relationship Link.
 *
 * @summary Gets information about the specified relationship Link.
 * x-ms-original-file: specification/customer-insights/resource-manager/Microsoft.CustomerInsights/stable/2017-04-26/examples/RelationshipLinksGet.json
 */
async function relationshipLinksGet() {
  const subscriptionId = "subid";
  const resourceGroupName = "TestHubRG";
  const hubName = "sdkTestHub";
  const relationshipLinkName = "Somelink";
  const credential = new DefaultAzureCredential();
  const client = new CustomerInsightsManagementClient(credential, subscriptionId);
  const result = await client.relationshipLinks.get(
    resourceGroupName,
    hubName,
    relationshipLinkName
  );
  console.log(result);
}

relationshipLinksGet().catch(console.error);
