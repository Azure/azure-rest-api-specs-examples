const { CustomerInsightsManagementClient } = require("@azure/arm-customerinsights");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates a relationship or updates an existing relationship within a hub.
 *
 * @summary Creates a relationship or updates an existing relationship within a hub.
 * x-ms-original-file: specification/customer-insights/resource-manager/Microsoft.CustomerInsights/stable/2017-04-26/examples/RelationshipsCreateOrUpdate.json
 */
async function relationshipsCreateOrUpdate() {
  const subscriptionId = "subid";
  const resourceGroupName = "TestHubRG";
  const hubName = "sdkTestHub";
  const relationshipName = "SomeRelationship";
  const parameters = {
    description: { enUs: "Relationship Description" },
    cardinality: "OneToOne",
    displayName: { enUs: "Relationship DisplayName" },
    fields: [],
    profileType: "testProfile2326994",
    relatedProfileType: "testProfile2326994",
  };
  const credential = new DefaultAzureCredential();
  const client = new CustomerInsightsManagementClient(credential, subscriptionId);
  const result = await client.relationships.beginCreateOrUpdateAndWait(
    resourceGroupName,
    hubName,
    relationshipName,
    parameters
  );
  console.log(result);
}

relationshipsCreateOrUpdate().catch(console.error);
