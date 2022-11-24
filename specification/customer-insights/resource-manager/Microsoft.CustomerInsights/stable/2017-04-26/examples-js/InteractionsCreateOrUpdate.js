const { CustomerInsightsManagementClient } = require("@azure/arm-customerinsights");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates an interaction or updates an existing interaction within a hub.
 *
 * @summary Creates an interaction or updates an existing interaction within a hub.
 * x-ms-original-file: specification/customer-insights/resource-manager/Microsoft.CustomerInsights/stable/2017-04-26/examples/InteractionsCreateOrUpdate.json
 */
async function interactionsCreateOrUpdate() {
  const subscriptionId = "subid";
  const resourceGroupName = "TestHubRG";
  const hubName = "sdkTestHub";
  const interactionName = "TestProfileType396";
  const parameters = {
    apiEntitySetName: "TestInteractionType6358",
    fields: [
      {
        fieldName: "TestInteractionType6358",
        fieldType: "Edm.String",
        isArray: false,
        isRequired: true,
      },
      { fieldName: "profile1", fieldType: "Edm.String" },
    ],
    idPropertyNames: ["TestInteractionType6358"],
    largeImage: "\\\\Images\\\\LargeImage",
    mediumImage: "\\\\Images\\\\MediumImage",
    primaryParticipantProfilePropertyName: "profile1",
    smallImage: "\\\\Images\\\\smallImage",
  };
  const credential = new DefaultAzureCredential();
  const client = new CustomerInsightsManagementClient(credential, subscriptionId);
  const result = await client.interactions.beginCreateOrUpdateAndWait(
    resourceGroupName,
    hubName,
    interactionName,
    parameters
  );
  console.log(result);
}

interactionsCreateOrUpdate().catch(console.error);
