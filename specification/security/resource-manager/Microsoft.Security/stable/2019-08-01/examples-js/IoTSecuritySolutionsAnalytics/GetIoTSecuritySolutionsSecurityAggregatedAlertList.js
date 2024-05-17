const { SecurityCenter } = require("@azure/arm-security");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Use this method to get the aggregated alert list of yours IoT Security solution.
 *
 * @summary Use this method to get the aggregated alert list of yours IoT Security solution.
 * x-ms-original-file: specification/security/resource-manager/Microsoft.Security/stable/2019-08-01/examples/IoTSecuritySolutionsAnalytics/GetIoTSecuritySolutionsSecurityAggregatedAlertList.json
 */
async function getTheAggregatedAlertListOfYoursIoTSecuritySolution() {
  const subscriptionId =
    process.env["SECURITY_SUBSCRIPTION_ID"] || "20ff7fc3-e762-44dd-bd96-b71116dcdc23";
  const resourceGroupName = process.env["SECURITY_RESOURCE_GROUP"] || "MyGroup";
  const solutionName = "default";
  const credential = new DefaultAzureCredential();
  const client = new SecurityCenter(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.iotSecuritySolutionsAnalyticsAggregatedAlert.list(
    resourceGroupName,
    solutionName,
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}
