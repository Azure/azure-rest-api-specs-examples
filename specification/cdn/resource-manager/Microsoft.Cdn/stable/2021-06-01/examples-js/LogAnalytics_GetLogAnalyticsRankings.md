Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-cdn_7.0.0/sdk/cdn/arm-cdn/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { CdnManagementClient } = require("@azure/arm-cdn");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get log analytics ranking report for AFD profile
 *
 * @summary Get log analytics ranking report for AFD profile
 * x-ms-original-file: specification/cdn/resource-manager/Microsoft.Cdn/stable/2021-06-01/examples/LogAnalytics_GetLogAnalyticsRankings.json
 */
async function logAnalyticsGetLogAnalyticsRankings() {
  const subscriptionId = "subid";
  const resourceGroupName = "RG";
  const profileName = "profile1";
  const rankings = ["url"];
  const metrics = ["clientRequestCount"];
  const maxRanking = 5;
  const dateTimeBegin = new Date("2020-11-04T06:49:27.554Z");
  const dateTimeEnd = new Date("2020-11-04T09:49:27.554Z");
  const credential = new DefaultAzureCredential();
  const client = new CdnManagementClient(credential, subscriptionId);
  const result = await client.logAnalytics.getLogAnalyticsRankings(
    resourceGroupName,
    profileName,
    rankings,
    metrics,
    maxRanking,
    dateTimeBegin,
    dateTimeEnd
  );
  console.log(result);
}

logAnalyticsGetLogAnalyticsRankings().catch(console.error);
```
