Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-security_5.0.0/sdk/security/arm-security/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { SecurityCenter } = require("@azure/arm-security");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Use this method to get a single the aggregated alert of yours IoT Security solution. This aggregation is performed by alert name.
 *
 * @summary Use this method to get a single the aggregated alert of yours IoT Security solution. This aggregation is performed by alert name.
 * x-ms-original-file: specification/security/resource-manager/Microsoft.Security/stable/2019-08-01/examples/IoTSecuritySolutionsAnalytics/GetIoTSecuritySolutionsSecurityAggregatedAlert.json
 */
async function getTheAggregatedSecurityAnalyticsAlertOfYoursIoTSecuritySolutionThisAggregationIsPerformedByAlertName() {
  const subscriptionId = "20ff7fc3-e762-44dd-bd96-b71116dcdc23";
  const resourceGroupName = "MyGroup";
  const solutionName = "default";
  const aggregatedAlertName = "IoT_Bruteforce_Fail/2019-02-02";
  const credential = new DefaultAzureCredential();
  const client = new SecurityCenter(credential, subscriptionId);
  const result = await client.iotSecuritySolutionsAnalyticsAggregatedAlert.get(
    resourceGroupName,
    solutionName,
    aggregatedAlertName
  );
  console.log(result);
}

getTheAggregatedSecurityAnalyticsAlertOfYoursIoTSecuritySolutionThisAggregationIsPerformedByAlertName().catch(
  console.error
);
```
