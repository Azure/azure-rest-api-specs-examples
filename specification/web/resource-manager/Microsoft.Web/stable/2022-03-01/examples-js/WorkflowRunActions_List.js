const { WebSiteManagementClient } = require("@azure/arm-appservice");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a list of workflow run actions.
 *
 * @summary Gets a list of workflow run actions.
 * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2022-03-01/examples/WorkflowRunActions_List.json
 */
async function listAWorkflowRunActions() {
  const subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
  const resourceGroupName = "test-resource-group";
  const name = "test-name";
  const workflowName = "test-workflow";
  const runName = "08586676746934337772206998657CU22";
  const credential = new DefaultAzureCredential();
  const client = new WebSiteManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.workflowRunActions.list(
    resourceGroupName,
    name,
    workflowName,
    runName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

listAWorkflowRunActions().catch(console.error);
