const { WebSiteManagementClient } = require("@azure/arm-appservice");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Gets a list of workflow trigger histories.
 *
 * @summary Gets a list of workflow trigger histories.
 * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2024-11-01/examples/WorkflowTriggerHistories_List.json
 */
async function listAWorkflowTriggerHistory() {
  const subscriptionId =
    process.env["APPSERVICE_SUBSCRIPTION_ID"] || "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
  const resourceGroupName = process.env["APPSERVICE_RESOURCE_GROUP"] || "testResourceGroup";
  const name = "test-name";
  const workflowName = "testWorkflowName";
  const triggerName = "testTriggerName";
  const credential = new DefaultAzureCredential();
  const client = new WebSiteManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (const item of client.workflowTriggerHistories.list(
    resourceGroupName,
    name,
    workflowName,
    triggerName,
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}
