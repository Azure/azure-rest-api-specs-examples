const { LogicManagementClient } = require("@azure/arm-logic");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Disables a workflow.
 *
 * @summary Disables a workflow.
 * x-ms-original-file: specification/logic/resource-manager/Microsoft.Logic/stable/2019-05-01/examples/Workflows_Disable.json
 */
async function disableAWorkflow() {
  const subscriptionId =
    process.env["LOGIC_SUBSCRIPTION_ID"] || "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
  const resourceGroupName = process.env["LOGIC_RESOURCE_GROUP"] || "test-resource-group";
  const workflowName = "test-workflow";
  const credential = new DefaultAzureCredential();
  const client = new LogicManagementClient(credential, subscriptionId);
  const result = await client.workflows.disable(resourceGroupName, workflowName);
  console.log(result);
}
