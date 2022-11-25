const { CustomerInsightsManagementClient } = require("@azure/arm-customerinsights");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates a link or updates an existing link in the hub.
 *
 * @summary Creates a link or updates an existing link in the hub.
 * x-ms-original-file: specification/customer-insights/resource-manager/Microsoft.CustomerInsights/stable/2017-04-26/examples/LinksCreateOrUpdate.json
 */
async function linksCreateOrUpdate() {
  const subscriptionId = "subid";
  const resourceGroupName = "TestHubRG";
  const hubName = "sdkTestHub";
  const linkName = "linkTest4806";
  const parameters = {
    description: { enUs: "Link Description" },
    displayName: { enUs: "Link DisplayName" },
    linkName: "linkTest4806",
    mappings: [
      {
        linkType: "UpdateAlways",
        sourcePropertyName: "testInteraction1949",
        targetPropertyName: "testProfile1446",
      },
    ],
    participantPropertyReferences: [
      {
        sourcePropertyName: "testInteraction1949",
        targetPropertyName: "ProfileId",
      },
    ],
    sourceEntityType: "Interaction",
    sourceEntityTypeName: "testInteraction1949",
    targetEntityType: "Profile",
    targetEntityTypeName: "testProfile1446",
  };
  const credential = new DefaultAzureCredential();
  const client = new CustomerInsightsManagementClient(credential, subscriptionId);
  const result = await client.links.beginCreateOrUpdateAndWait(
    resourceGroupName,
    hubName,
    linkName,
    parameters
  );
  console.log(result);
}

linksCreateOrUpdate().catch(console.error);
