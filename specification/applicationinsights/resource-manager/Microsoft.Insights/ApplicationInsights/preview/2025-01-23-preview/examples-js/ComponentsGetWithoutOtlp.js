const { ApplicationInsightsManagementClient } = require("@azure/arm-appinsights");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to returns an Application Insights component.
 *
 * @summary returns an Application Insights component.
 * x-ms-original-file: 2025-01-23-preview/ComponentsGetWithoutOtlp.json
 */
async function componentGetWithoutOtlp() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new ApplicationInsightsManagementClient(credential, subscriptionId);
  const result = await client.components.get("my-resource-group", "my-component");
  console.log(result);
}
