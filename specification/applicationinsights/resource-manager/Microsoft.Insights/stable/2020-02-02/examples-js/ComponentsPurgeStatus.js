const { ApplicationInsightsManagementClient } = require("@azure/arm-appinsights");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Get status for an ongoing purge operation.
 *
 * @summary Get status for an ongoing purge operation.
 * x-ms-original-file: specification/applicationinsights/resource-manager/Microsoft.Insights/stable/2020-02-02/examples/ComponentsPurgeStatus.json
 */
async function componentPurge() {
  const subscriptionId =
    process.env["APPLICATIONINSIGHTS_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-00000000000";
  const resourceGroupName = process.env["APPLICATIONINSIGHTS_RESOURCE_GROUP"] || "OIAutoRest5123";
  const resourceName = "aztest5048";
  const purgeId = "purge-970318e7-b859-4edb-8903-83b1b54d0b74";
  const credential = new DefaultAzureCredential();
  const client = new ApplicationInsightsManagementClient(credential, subscriptionId);
  const result = await client.components.getPurgeStatus(resourceGroupName, resourceName, purgeId);
  console.log(result);
}
