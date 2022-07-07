const { WebSiteManagementClient } = require("@azure/arm-appservice");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List a workflow run repetition request history.
 *
 * @summary List a workflow run repetition request history.
 * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2022-03-01/examples/WorkflowRunActionRepetitionsRequestHistories_List.json
 */
async function listRepetitionRequestHistory() {
  const subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
  const resourceGroupName = "test-resource-group";
  const name = "test-name";
  const workflowName = "test-workflow";
  const runName = "08586776228332053161046300351";
  const actionName = "HTTP_Webhook";
  const repetitionName = "000001";
  const credential = new DefaultAzureCredential();
  const client = new WebSiteManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.workflowRunActionRepetitionsRequestHistories.list(
    resourceGroupName,
    name,
    workflowName,
    runName,
    actionName,
    repetitionName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

listRepetitionRequestHistory().catch(console.error);
