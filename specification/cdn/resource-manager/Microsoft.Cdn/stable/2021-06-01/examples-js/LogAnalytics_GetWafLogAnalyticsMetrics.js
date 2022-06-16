const { CdnManagementClient } = require("@azure/arm-cdn");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get Waf related log analytics report for AFD profile.
 *
 * @summary Get Waf related log analytics report for AFD profile.
 * x-ms-original-file: specification/cdn/resource-manager/Microsoft.Cdn/stable/2021-06-01/examples/LogAnalytics_GetWafLogAnalyticsMetrics.json
 */
async function logAnalyticsGetWafLogAnalyticsMetrics() {
  const subscriptionId = "subid";
  const resourceGroupName = "RG";
  const profileName = "profile1";
  const metrics = ["clientRequestCount"];
  const dateTimeBegin = new Date("2020-11-04T06:49:27.554Z");
  const dateTimeEnd = new Date("2020-11-04T09:49:27.554Z");
  const granularity = "PT5M";
  const actions = ["block", "log"];
  const options = {
    actions,
  };
  const credential = new DefaultAzureCredential();
  const client = new CdnManagementClient(credential, subscriptionId);
  const result = await client.logAnalytics.getWafLogAnalyticsMetrics(
    resourceGroupName,
    profileName,
    metrics,
    dateTimeBegin,
    dateTimeEnd,
    granularity,
    options
  );
  console.log(result);
}

logAnalyticsGetWafLogAnalyticsMetrics().catch(console.error);
