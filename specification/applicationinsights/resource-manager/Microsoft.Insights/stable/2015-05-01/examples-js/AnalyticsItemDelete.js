const { ApplicationInsightsManagementClient } = require("@azure/arm-appinsights");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Deletes a specific Analytics Items defined within an Application Insights component.
 *
 * @summary Deletes a specific Analytics Items defined within an Application Insights component.
 * x-ms-original-file: specification/applicationinsights/resource-manager/Microsoft.Insights/stable/2015-05-01/examples/AnalyticsItemDelete.json
 */
async function analyticsItemDelete() {
  const subscriptionId = process.env["APPLICATIONINSIGHTS_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName =
    process.env["APPLICATIONINSIGHTS_RESOURCE_GROUP"] || "my-resource-group";
  const resourceName = "my-component";
  const scopePath = "analyticsItems";
  const id = "3466c160-4a10-4df8-afdf-0007f3f6dee5";
  const options = { id };
  const credential = new DefaultAzureCredential();
  const client = new ApplicationInsightsManagementClient(credential, subscriptionId);
  const result = await client.analyticsItems.delete(
    resourceGroupName,
    resourceName,
    scopePath,
    options,
  );
  console.log(result);
}
