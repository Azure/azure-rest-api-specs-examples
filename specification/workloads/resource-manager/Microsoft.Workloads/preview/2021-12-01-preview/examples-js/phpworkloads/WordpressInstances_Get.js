const { WorkloadsClient } = require("@azure/arm-workloads");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the WordPress instance resource.
 *
 * @summary Gets the WordPress instance resource.
 * x-ms-original-file: specification/workloads/resource-manager/Microsoft.Workloads/preview/2021-12-01-preview/examples/phpworkloads/WordpressInstances_Get.json
 */
async function workloads() {
  const subscriptionId = "8e17e36c-42e9-4cd5-a078-7b44883414e0";
  const resourceGroupName = "test-rg";
  const phpWorkloadName = "wp39";
  const credential = new DefaultAzureCredential();
  const client = new WorkloadsClient(credential, subscriptionId);
  const result = await client.wordpressInstances.get(resourceGroupName, phpWorkloadName);
  console.log(result);
}

workloads().catch(console.error);
