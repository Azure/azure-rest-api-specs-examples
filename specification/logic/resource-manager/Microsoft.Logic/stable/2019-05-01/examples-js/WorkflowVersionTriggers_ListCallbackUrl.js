const { LogicManagementClient } = require("@azure/arm-logic");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get the callback url for a trigger of a workflow version.
 *
 * @summary Get the callback url for a trigger of a workflow version.
 * x-ms-original-file: specification/logic/resource-manager/Microsoft.Logic/stable/2019-05-01/examples/WorkflowVersionTriggers_ListCallbackUrl.json
 */
async function getTheCallbackUrlForATriggerOfAWorkflowVersion() {
  const subscriptionId =
    process.env["LOGIC_SUBSCRIPTION_ID"] || "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
  const resourceGroupName = process.env["LOGIC_RESOURCE_GROUP"] || "testResourceGroup";
  const workflowName = "testWorkflowName";
  const versionId = "testWorkflowVersionId";
  const triggerName = "testTriggerName";
  const parameters = {
    keyType: "Primary",
    notAfter: new Date("2017-03-05T08:00:00Z"),
  };
  const options = {
    parameters,
  };
  const credential = new DefaultAzureCredential();
  const client = new LogicManagementClient(credential, subscriptionId);
  const result = await client.workflowVersionTriggers.listCallbackUrl(
    resourceGroupName,
    workflowName,
    versionId,
    triggerName,
    options
  );
  console.log(result);
}
