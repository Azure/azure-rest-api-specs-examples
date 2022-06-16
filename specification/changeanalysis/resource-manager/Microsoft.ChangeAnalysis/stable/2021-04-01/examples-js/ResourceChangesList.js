const { AzureChangeAnalysisManagementClient } = require("@azure/arm-changeanalysis");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List the changes of a resource within the specified time range. Customer data will be masked if the user doesn't have access.
 *
 * @summary List the changes of a resource within the specified time range. Customer data will be masked if the user doesn't have access.
 * x-ms-original-file: specification/changeanalysis/resource-manager/Microsoft.ChangeAnalysis/stable/2021-04-01/examples/ResourceChangesList.json
 */
async function resourceChangesList() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceId =
    "subscriptions/4d962866-1e3f-47f2-bd18-450c08f914c1/resourceGroups/MyResourceGroup/providers/Microsoft.Web/sites/mysite";
  const startTime = new Date("2021-04-25T12:09:03.141Z");
  const endTime = new Date("2021-04-26T12:09:03.141Z");
  const credential = new DefaultAzureCredential();
  const client = new AzureChangeAnalysisManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.resourceChanges.list(resourceId, startTime, endTime)) {
    resArray.push(item);
  }
  console.log(resArray);
}

resourceChangesList().catch(console.error);
