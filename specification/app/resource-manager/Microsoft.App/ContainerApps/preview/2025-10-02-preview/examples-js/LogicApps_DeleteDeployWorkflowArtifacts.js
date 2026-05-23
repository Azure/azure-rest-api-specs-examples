const { ContainerAppsAPIClient } = require("@azure/arm-appcontainers");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates or updates the artifacts for the logic app
 *
 * @summary creates or updates the artifacts for the logic app
 * x-ms-original-file: 2025-10-02-preview/LogicApps_DeleteDeployWorkflowArtifacts.json
 */
async function deleteWorkflowArtifacts() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
  const client = new ContainerAppsAPIClient(credential, subscriptionId);
  await client.logicApps.deployWorkflowArtifacts("testrg123", "testapp2", "testapp2", {
    workflowArtifacts: { filesToDelete: ["test/workflow.json", "test/"] },
  });
}
