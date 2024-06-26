const { LogicManagementClient } = require("@azure/arm-logic");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Moves an existing workflow.
 *
 * @summary Moves an existing workflow.
 * x-ms-original-file: specification/logic/resource-manager/Microsoft.Logic/stable/2019-05-01/examples/Workflows_Move.json
 */
async function moveAWorkflow() {
  const subscriptionId =
    process.env["LOGIC_SUBSCRIPTION_ID"] || "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
  const resourceGroupName = process.env["LOGIC_RESOURCE_GROUP"] || "testResourceGroup";
  const workflowName = "testWorkflow";
  const move = {
    id: "subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/newResourceGroup/providers/Microsoft.Logic/workflows/newWorkflowName",
  };
  const credential = new DefaultAzureCredential();
  const client = new LogicManagementClient(credential, subscriptionId);
  const result = await client.workflows.beginMoveAndWait(resourceGroupName, workflowName, move);
  console.log(result);
}
