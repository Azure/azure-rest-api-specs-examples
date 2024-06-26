const { CustomerInsightsManagementClient } = require("@azure/arm-customerinsights");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets information about the specified relationship.
 *
 * @summary Gets information about the specified relationship.
 * x-ms-original-file: specification/customer-insights/resource-manager/Microsoft.CustomerInsights/stable/2017-04-26/examples/RelationshipsGet.json
 */
async function relationshipsGet() {
  const subscriptionId = "subid";
  const resourceGroupName = "TestHubRG";
  const hubName = "sdkTestHub";
  const relationshipName = "SomeRelationship";
  const credential = new DefaultAzureCredential();
  const client = new CustomerInsightsManagementClient(credential, subscriptionId);
  const result = await client.relationships.get(resourceGroupName, hubName, relationshipName);
  console.log(result);
}

relationshipsGet().catch(console.error);
