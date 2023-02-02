const { LoadTestClient } = require("@azure/arm-loadtesting");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists loadtest resources in a resource group.
 *
 * @summary Lists loadtest resources in a resource group.
 * x-ms-original-file: specification/loadtestservice/resource-manager/Microsoft.LoadTestService/stable/2022-12-01/examples/LoadTests_ListByResourceGroup.json
 */
async function loadTestsListByResourceGroup() {
  const subscriptionId =
    process.env["LOADTESTSERVICE_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["LOADTESTSERVICE_RESOURCE_GROUP"] || "dummyrg";
  const credential = new DefaultAzureCredential();
  const client = new LoadTestClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.loadTests.listByResourceGroup(resourceGroupName)) {
    resArray.push(item);
  }
  console.log(resArray);
}
