const { LogicManagementClient } = require("@azure/arm-logic");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Sets the state of a workflow trigger.
 *
 * @summary Sets the state of a workflow trigger.
 * x-ms-original-file: specification/logic/resource-manager/Microsoft.Logic/stable/2019-05-01/examples/WorkflowTriggers_SetState.json
 */
async function setTriggerState() {
  const subscriptionId =
    process.env["LOGIC_SUBSCRIPTION_ID"] || "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
  const resourceGroupName = process.env["LOGIC_RESOURCE_GROUP"] || "testResourceGroup";
  const workflowName = "testWorkflow";
  const triggerName = "testTrigger";
  const setState = {
    source: {
      id: "subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/sourceResGroup/providers/Microsoft.Logic/workflows/sourceWorkflow/triggers/sourceTrigger",
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new LogicManagementClient(credential, subscriptionId);
  const result = await client.workflowTriggers.setState(
    resourceGroupName,
    workflowName,
    triggerName,
    setState
  );
  console.log(result);
}
