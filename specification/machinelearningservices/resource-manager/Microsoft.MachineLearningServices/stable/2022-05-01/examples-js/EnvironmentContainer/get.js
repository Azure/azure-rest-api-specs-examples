const { AzureMachineLearningWorkspaces } = require("@azure/arm-machinelearning");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get container.
 *
 * @summary Get container.
 * x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2022-05-01/examples/EnvironmentContainer/get.json
 */
async function getEnvironmentContainer() {
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = "testrg123";
  const workspaceName = "testworkspace";
  const name = "testEnvironment";
  const credential = new DefaultAzureCredential();
  const client = new AzureMachineLearningWorkspaces(credential, subscriptionId);
  const result = await client.environmentContainers.get(resourceGroupName, workspaceName, name);
  console.log(result);
}

getEnvironmentContainer().catch(console.error);
