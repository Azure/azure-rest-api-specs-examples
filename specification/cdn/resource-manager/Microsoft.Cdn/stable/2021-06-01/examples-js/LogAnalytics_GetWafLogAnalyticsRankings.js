const { CdnManagementClient } = require("@azure/arm-cdn");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get WAF log analytics charts for AFD profile
 *
 * @summary Get WAF log analytics charts for AFD profile
 * x-ms-original-file: specification/cdn/resource-manager/Microsoft.Cdn/stable/2021-06-01/examples/LogAnalytics_GetWafLogAnalyticsRankings.json
 */
async function logAnalyticsGetWafLogAnalyticsRankings() {
  const subscriptionId = "subid";
  const resourceGroupName = "RG";
  const profileName = "profile1";
  const metrics = ["clientRequestCount"];
  const dateTimeBegin = new Date("2020-11-04T06:49:27.554Z");
  const dateTimeEnd = new Date("2020-11-04T09:49:27.554Z");
  const maxRanking = 5;
  const rankings = ["ruleId"];
  const credential = new DefaultAzureCredential();
  const client = new CdnManagementClient(credential, subscriptionId);
  const result = await client.logAnalytics.getWafLogAnalyticsRankings(
    resourceGroupName,
    profileName,
    metrics,
    dateTimeBegin,
    dateTimeEnd,
    maxRanking,
    rankings
  );
  console.log(result);
}

logAnalyticsGetWafLogAnalyticsRankings().catch(console.error);
