const { AzureMachineLearningWorkspaces } = require("@azure/arm-machinelearning");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Delete version.
 *
 * @summary Delete version.
 * x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2022-05-01/examples/ComponentVersion/delete.json
 */
async function deleteComponentVersion() {
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = "test-rg";
  const workspaceName = "my-aml-workspace";
  const name = "string";
  const version = "string";
  const credential = new DefaultAzureCredential();
  const client = new AzureMachineLearningWorkspaces(credential, subscriptionId);
  const result = await client.componentVersions.delete(
    resourceGroupName,
    workspaceName,
    name,
    version
  );
  console.log(result);
}

deleteComponentVersion().catch(console.error);
