const { AzureMachineLearningWorkspaces } = require("@azure/arm-machinelearning");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List versions.
 *
 * @summary List versions.
 * x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2022-05-01/examples/EnvironmentVersion/list.json
 */
async function listEnvironmentVersion() {
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = "test-rg";
  const workspaceName = "my-aml-workspace";
  const name = "string";
  const orderBy = "string";
  const top = 1;
  const options = { orderBy, top };
  const credential = new DefaultAzureCredential();
  const client = new AzureMachineLearningWorkspaces(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.environmentVersions.list(
    resourceGroupName,
    workspaceName,
    name,
    options
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

listEnvironmentVersion().catch(console.error);
