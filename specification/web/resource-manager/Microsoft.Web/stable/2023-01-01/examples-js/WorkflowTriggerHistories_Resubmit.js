const { WebSiteManagementClient } = require("@azure/arm-appservice");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Resubmits a workflow run based on the trigger history.
 *
 * @summary Resubmits a workflow run based on the trigger history.
 * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2023-01-01/examples/WorkflowTriggerHistories_Resubmit.json
 */
async function resubmitAWorkflowRunBasedOnTheTriggerHistory() {
  const subscriptionId =
    process.env["APPSERVICE_SUBSCRIPTION_ID"] || "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
  const resourceGroupName = process.env["APPSERVICE_RESOURCE_GROUP"] || "testResourceGroup";
  const name = "test-name";
  const workflowName = "testWorkflowName";
  const triggerName = "testTriggerName";
  const historyName = "testHistoryName";
  const credential = new DefaultAzureCredential();
  const client = new WebSiteManagementClient(credential, subscriptionId);
  const result = await client.workflowTriggerHistories.beginResubmitAndWait(
    resourceGroupName,
    name,
    workflowName,
    triggerName,
    historyName
  );
  console.log(result);
}
