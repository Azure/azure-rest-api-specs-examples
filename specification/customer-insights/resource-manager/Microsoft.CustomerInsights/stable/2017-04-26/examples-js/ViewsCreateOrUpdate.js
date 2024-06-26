const { CustomerInsightsManagementClient } = require("@azure/arm-customerinsights");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates a view or updates an existing view in the hub.
 *
 * @summary Creates a view or updates an existing view in the hub.
 * x-ms-original-file: specification/customer-insights/resource-manager/Microsoft.CustomerInsights/stable/2017-04-26/examples/ViewsCreateOrUpdate.json
 */
async function viewsCreateOrUpdate() {
  const subscriptionId = "subid";
  const resourceGroupName = "TestHubRG";
  const hubName = "sdkTestHub";
  const viewName = "testView";
  const parameters = {
    definition: '{\\"isProfileType\\":false,\\"profileTypes\\":[],\\"widgets\\":[],\\"style\\":[]}',
    displayName: { en: "some name" },
    userId: "testUser",
  };
  const credential = new DefaultAzureCredential();
  const client = new CustomerInsightsManagementClient(credential, subscriptionId);
  const result = await client.views.createOrUpdate(
    resourceGroupName,
    hubName,
    viewName,
    parameters
  );
  console.log(result);
}

viewsCreateOrUpdate().catch(console.error);
