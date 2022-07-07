const { WebSiteManagementClient } = require("@azure/arm-appservice");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a workflow run repetition request history.
 *
 * @summary Gets a workflow run repetition request history.
 * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2022-03-01/examples/WorkflowRunActionRepetitionsRequestHistories_Get.json
 */
async function getARepetitionRequestHistory() {
  const subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
  const resourceGroupName = "test-resource-group";
  const name = "test-name";
  const workflowName = "test-workflow";
  const runName = "08586776228332053161046300351";
  const actionName = "HTTP_Webhook";
  const repetitionName = "000001";
  const requestHistoryName = "08586611142732800686";
  const credential = new DefaultAzureCredential();
  const client = new WebSiteManagementClient(credential, subscriptionId);
  const result = await client.workflowRunActionRepetitionsRequestHistories.get(
    resourceGroupName,
    name,
    workflowName,
    runName,
    actionName,
    repetitionName,
    requestHistoryName
  );
  console.log(result);
}

getARepetitionRequestHistory().catch(console.error);
