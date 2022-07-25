const { MLTeamAccountManagementClient } = require("@azure/arm-machinelearningexperimentation");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates a project with the specified parameters.
 *
 * @summary Creates or updates a project with the specified parameters.
 * x-ms-original-file: specification/machinelearningexperimentation/resource-manager/Microsoft.MachineLearningExperimentation/preview/2017-05-01-preview/examples/CreateProject.json
 */
async function createProject() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "myResourceGroup";
  const accountName = "testaccount";
  const workspaceName = "testworkspace";
  const projectName = "testProject";
  const parameters = {
    friendlyName: "testName",
    gitrepo: "https://github/abc",
    location: "East US",
    tags: { tagKey1: "TagValue1" },
  };
  const credential = new DefaultAzureCredential();
  const client = new MLTeamAccountManagementClient(credential, subscriptionId);
  const result = await client.projects.createOrUpdate(
    resourceGroupName,
    accountName,
    workspaceName,
    projectName,
    parameters
  );
  console.log(result);
}

createProject().catch(console.error);
