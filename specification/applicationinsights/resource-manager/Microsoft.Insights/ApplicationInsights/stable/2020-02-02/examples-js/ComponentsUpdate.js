const { ApplicationInsightsManagementClient } = require("@azure/arm-appinsights");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates (or updates) an Application Insights component. Note: You cannot specify a different value for InstrumentationKey nor AppId in the Put operation.
 *
 * @summary creates (or updates) an Application Insights component. Note: You cannot specify a different value for InstrumentationKey nor AppId in the Put operation.
 * x-ms-original-file: 2020-02-02/ComponentsUpdate.json
 */
async function componentUpdate() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "subid";
  const client = new ApplicationInsightsManagementClient(credential, subscriptionId);
  const result = await client.components.createOrUpdate("my-resource-group", "my-component", {
    kind: "web",
    location: "South Central US",
    tags: { ApplicationGatewayType: "Internal-Only", BillingEntity: "Self" },
  });
  console.log(result);
}
