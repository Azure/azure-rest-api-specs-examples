const { DevTestLabsClient } = require("@azure/arm-devtestlabs");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists all applicable schedules
 *
 * @summary Lists all applicable schedules
 * x-ms-original-file: specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/Schedules_ListApplicable.json
 */
async function schedulesListApplicable() {
  const subscriptionId = "{subscriptionId}";
  const resourceGroupName = "resourceGroupName";
  const labName = "{labName}";
  const name = "{scheduleName}";
  const credential = new DefaultAzureCredential();
  const client = new DevTestLabsClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.schedules.listApplicable(resourceGroupName, labName, name)) {
    resArray.push(item);
  }
  console.log(resArray);
}

schedulesListApplicable().catch(console.error);
