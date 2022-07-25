const { WorkloadsClient } = require("@azure/arm-workloads");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Delete PHP workload resource.
 *
 * @summary Delete PHP workload resource.
 * x-ms-original-file: specification/workloads/resource-manager/Microsoft.Workloads/preview/2021-12-01-preview/examples/phpworkloads/PhpWorkloads_Delete.json
 */
async function workloads() {
  const subscriptionId = "8e17e36c-42e9-4cd5-a078-7b44883414e0";
  const resourceGroupName = "test-rg";
  const phpWorkloadName = "wp39";
  const deleteInfra = "false";
  const options = { deleteInfra };
  const credential = new DefaultAzureCredential();
  const client = new WorkloadsClient(credential, subscriptionId);
  const result = await client.phpWorkloads.beginDeleteAndWait(
    resourceGroupName,
    phpWorkloadName,
    options
  );
  console.log(result);
}

workloads().catch(console.error);
