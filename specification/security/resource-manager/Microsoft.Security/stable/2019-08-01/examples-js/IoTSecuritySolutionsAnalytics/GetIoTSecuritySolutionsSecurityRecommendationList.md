```javascript
const { SecurityCenter } = require("@azure/arm-security");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Use this method to get the list of aggregated security analytics recommendations of yours IoT Security solution.
 *
 * @summary Use this method to get the list of aggregated security analytics recommendations of yours IoT Security solution.
 * x-ms-original-file: specification/security/resource-manager/Microsoft.Security/stable/2019-08-01/examples/IoTSecuritySolutionsAnalytics/GetIoTSecuritySolutionsSecurityRecommendationList.json
 */
async function getTheListOfAggregatedSecurityAnalyticsRecommendationsOfYoursIoTSecuritySolution() {
  const subscriptionId = "075423e9-7d33-4166-8bdf-3920b04e3735";
  const resourceGroupName = "IoTEdgeResources";
  const solutionName = "default";
  const credential = new DefaultAzureCredential();
  const client = new SecurityCenter(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.iotSecuritySolutionsAnalyticsRecommendation.list(
    resourceGroupName,
    solutionName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

getTheListOfAggregatedSecurityAnalyticsRecommendationsOfYoursIoTSecuritySolution().catch(
  console.error
);
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-security_5.0.0/sdk/security/arm-security/README.md) on how to add the SDK to your project and authenticate.
