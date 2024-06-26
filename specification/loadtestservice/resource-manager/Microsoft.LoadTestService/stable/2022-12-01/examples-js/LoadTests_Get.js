const { LoadTestClient } = require("@azure/arm-loadtesting");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get a LoadTest resource.
 *
 * @summary Get a LoadTest resource.
 * x-ms-original-file: specification/loadtestservice/resource-manager/Microsoft.LoadTestService/stable/2022-12-01/examples/LoadTests_Get.json
 */
async function loadTestsGet() {
  const subscriptionId =
    process.env["LOADTESTSERVICE_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["LOADTESTSERVICE_RESOURCE_GROUP"] || "dummyrg";
  const loadTestName = "myLoadTest";
  const credential = new DefaultAzureCredential();
  const client = new LoadTestClient(credential, subscriptionId);
  const result = await client.loadTests.get(resourceGroupName, loadTestName);
  console.log(result);
}
