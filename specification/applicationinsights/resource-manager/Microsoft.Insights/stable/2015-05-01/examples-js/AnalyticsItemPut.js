const { ApplicationInsightsManagementClient } = require("@azure/arm-appinsights");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Adds or Updates a specific Analytics Item within an Application Insights component.
 *
 * @summary Adds or Updates a specific Analytics Item within an Application Insights component.
 * x-ms-original-file: specification/applicationinsights/resource-manager/Microsoft.Insights/stable/2015-05-01/examples/AnalyticsItemPut.json
 */
async function analyticsItemPut() {
  const subscriptionId = process.env["APPLICATIONINSIGHTS_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName =
    process.env["APPLICATIONINSIGHTS_RESOURCE_GROUP"] || "my-resource-group";
  const resourceName = "my-component";
  const scopePath = "analyticsItems";
  const itemProperties = {
    content:
      "let newExceptionsTimeRange = 1d;\nlet timeRangeToCheckBefore = 7d;\nexceptions\n| where timestamp < ago(timeRangeToCheckBefore)\n| summarize count() by problemId\n| join kind= rightanti (\nexceptions\n| where timestamp >= ago(newExceptionsTimeRange)\n| extend stack = tostring(details[0].rawStack)\n| summarize count(), dcount(user_AuthenticatedId), min(timestamp), max(timestamp), any(stack) by problemId  \n) on problemId \n| order by  count_ desc\n",
    name: "Exceptions - New in the last 24 hours",
    scope: "shared",
    type: "query",
  };
  const credential = new DefaultAzureCredential();
  const client = new ApplicationInsightsManagementClient(credential, subscriptionId);
  const result = await client.analyticsItems.put(
    resourceGroupName,
    resourceName,
    scopePath,
    itemProperties,
  );
  console.log(result);
}
