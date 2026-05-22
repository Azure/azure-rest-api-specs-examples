const { ApplicationInsightsManagementClient } = require("@azure/arm-appinsights");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates (or updates) an Application Insights component. Note: You cannot specify a different value for InstrumentationKey nor AppId in the Put operation.
 *
 * @summary creates (or updates) an Application Insights component. Note: You cannot specify a different value for InstrumentationKey nor AppId in the Put operation.
 * x-ms-original-file: 2020-02-02/ComponentsCreate.json
 */
async function componentCreate() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "subid";
  const client = new ApplicationInsightsManagementClient(credential, subscriptionId);
  const result = await client.components.createOrUpdate("my-resource-group", "my-component", {
    kind: "web",
    location: "South Central US",
    applicationType: "web",
    flowType: "Bluefield",
    requestSource: "rest",
    workspaceResourceId:
      "/subscriptions/subid/resourcegroups/my-resource-group/providers/microsoft.operationalinsights/workspaces/my-workspace",
  });
  console.log(result);
}
