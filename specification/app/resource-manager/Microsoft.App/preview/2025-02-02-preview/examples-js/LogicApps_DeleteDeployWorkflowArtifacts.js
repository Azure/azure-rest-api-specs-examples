const { ContainerAppsAPIClient } = require("@azure/arm-appcontainers");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Creates or updates the artifacts for the logic app
 *
 * @summary Creates or updates the artifacts for the logic app
 * x-ms-original-file: specification/app/resource-manager/Microsoft.App/preview/2025-02-02-preview/examples/LogicApps_DeleteDeployWorkflowArtifacts.json
 */
async function deleteWorkflowArtifacts() {
  const subscriptionId =
    process.env["APPCONTAINERS_SUBSCRIPTION_ID"] || "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
  const resourceGroupName = process.env["APPCONTAINERS_RESOURCE_GROUP"] || "testrg123";
  const containerAppName = "testapp2";
  const logicAppName = "testapp2";
  const workflowArtifacts = {
    filesToDelete: ["test/workflow.json", "test/"],
  };
  const options = {
    workflowArtifacts,
  };
  const credential = new DefaultAzureCredential();
  const client = new ContainerAppsAPIClient(credential, subscriptionId);
  const result = await client.logicApps.deployWorkflowArtifacts(
    resourceGroupName,
    containerAppName,
    logicAppName,
    options,
  );
  console.log(result);
}
