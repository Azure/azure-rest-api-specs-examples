Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-cdn_7.0.1/sdk/cdn/arm-cdn/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { CdnManagementClient } = require("@azure/arm-cdn");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get log report for AFD profile
 *
 * @summary Get log report for AFD profile
 * x-ms-original-file: specification/cdn/resource-manager/Microsoft.Cdn/stable/2021-06-01/examples/LogAnalytics_GetLogAnalyticsMetrics.json
 */
async function logAnalyticsGetLogAnalyticsMetrics() {
  const subscriptionId = "subid";
  const resourceGroupName = "RG";
  const profileName = "profile1";
  const metrics = ["clientRequestCount"];
  const dateTimeBegin = new Date("2020-11-04T04:30:00.000Z");
  const dateTimeEnd = new Date("2020-11-04T05:00:00.000Z");
  const granularity = "PT5M";
  const groupBy = ["protocol"];
  const customDomains = ["customdomain1.azurecdn.net", "customdomain2.azurecdn.net"];
  const protocols = ["https"];
  const options = { groupBy };
  const credential = new DefaultAzureCredential();
  const client = new CdnManagementClient(credential, subscriptionId);
  const result = await client.logAnalytics.getLogAnalyticsMetrics(
    resourceGroupName,
    profileName,
    metrics,
    dateTimeBegin,
    dateTimeEnd,
    granularity,
    customDomains,
    protocols,
    options
  );
  console.log(result);
}

logAnalyticsGetLogAnalyticsMetrics().catch(console.error);
```
