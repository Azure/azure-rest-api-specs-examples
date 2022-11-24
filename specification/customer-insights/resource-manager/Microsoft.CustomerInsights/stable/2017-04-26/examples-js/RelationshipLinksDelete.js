const { CustomerInsightsManagementClient } = require("@azure/arm-customerinsights");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes a relationship link within a hub.
 *
 * @summary Deletes a relationship link within a hub.
 * x-ms-original-file: specification/customer-insights/resource-manager/Microsoft.CustomerInsights/stable/2017-04-26/examples/RelationshipLinksDelete.json
 */
async function relationshipLinksDelete() {
  const subscriptionId = "subid";
  const resourceGroupName = "TestHubRG";
  const hubName = "sdkTestHub";
  const relationshipLinkName = "Somelink";
  const credential = new DefaultAzureCredential();
  const client = new CustomerInsightsManagementClient(credential, subscriptionId);
  const result = await client.relationshipLinks.beginDeleteAndWait(
    resourceGroupName,
    hubName,
    relationshipLinkName
  );
  console.log(result);
}

relationshipLinksDelete().catch(console.error);
