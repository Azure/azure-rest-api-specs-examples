const { WebSiteManagementClient } = require("@azure/arm-appservice");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get the callback URL for a workflow trigger.
 *
 * @summary Get the callback URL for a workflow trigger.
 * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2022-03-01/examples/WorkflowTriggers_ListCallbackUrl.json
 */
async function getTheCallbackUrlForATrigger() {
  const subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
  const resourceGroupName = "test-resource-group";
  const name = "test-name";
  const workflowName = "test-workflow";
  const triggerName = "manual";
  const credential = new DefaultAzureCredential();
  const client = new WebSiteManagementClient(credential, subscriptionId);
  const result = await client.workflowTriggers.listCallbackUrl(
    resourceGroupName,
    name,
    workflowName,
    triggerName
  );
  console.log(result);
}

getTheCallbackUrlForATrigger().catch(console.error);
