const { MLTeamAccountManagementClient } = require("@azure/arm-machinelearningexperimentation");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists all the available machine learning projects under the specified workspace.
 *
 * @summary Lists all the available machine learning projects under the specified workspace.
 * x-ms-original-file: specification/machinelearningexperimentation/resource-manager/Microsoft.MachineLearningExperimentation/preview/2017-05-01-preview/examples/ProjectListByWorkspaces.json
 */
async function projectListByWorkspaces() {
  const subscriptionId =
    process.env["MACHINELEARNINGEXPERIMENTATION_SUBSCRIPTION_ID"] ||
    "00000000-1111-2222-3333-444444444444";
  const accountName = "testaccount";
  const workspaceName = "testworkspace";
  const resourceGroupName =
    process.env["MACHINELEARNINGEXPERIMENTATION_RESOURCE_GROUP"] || "testrg";
  const credential = new DefaultAzureCredential();
  const client = new MLTeamAccountManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.projects.listByWorkspace(
    accountName,
    workspaceName,
    resourceGroupName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}
