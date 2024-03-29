const { LogicManagementClient } = require("@azure/arm-logic");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Generates the upgraded definition for a workflow.
 *
 * @summary Generates the upgraded definition for a workflow.
 * x-ms-original-file: specification/logic/resource-manager/Microsoft.Logic/stable/2019-05-01/examples/Workflows_GenerateUpgradedDefinition.json
 */
async function generateAnUpgradedDefinition() {
  const subscriptionId =
    process.env["LOGIC_SUBSCRIPTION_ID"] || "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
  const resourceGroupName = process.env["LOGIC_RESOURCE_GROUP"] || "test-resource-group";
  const workflowName = "test-workflow";
  const parameters = {
    targetSchemaVersion: "2016-06-01",
  };
  const credential = new DefaultAzureCredential();
  const client = new LogicManagementClient(credential, subscriptionId);
  const result = await client.workflows.generateUpgradedDefinition(
    resourceGroupName,
    workflowName,
    parameters
  );
  console.log(result);
}
