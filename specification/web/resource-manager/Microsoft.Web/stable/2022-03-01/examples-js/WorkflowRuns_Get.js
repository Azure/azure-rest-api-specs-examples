const { WebSiteManagementClient } = require("@azure/arm-appservice");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a workflow run.
 *
 * @summary Gets a workflow run.
 * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2022-03-01/examples/WorkflowRuns_Get.json
 */
async function getARunForAWorkflow() {
  const subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
  const resourceGroupName = "test-resource-group";
  const name = "test-name";
  const workflowName = "test-workflow";
  const runName = "08586676746934337772206998657CU22";
  const credential = new DefaultAzureCredential();
  const client = new WebSiteManagementClient(credential, subscriptionId);
  const result = await client.workflowRuns.get(resourceGroupName, name, workflowName, runName);
  console.log(result);
}

getARunForAWorkflow().catch(console.error);
