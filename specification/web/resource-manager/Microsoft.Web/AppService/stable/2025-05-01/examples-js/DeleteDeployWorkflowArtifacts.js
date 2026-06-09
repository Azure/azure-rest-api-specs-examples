const { WebSiteManagementClient } = require("@azure/arm-appservice");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to description for Creates the artifacts for web site, or a deployment slot.
 *
 * @summary description for Creates the artifacts for web site, or a deployment slot.
 * x-ms-original-file: 2025-05-01/DeleteDeployWorkflowArtifacts.json
 */
async function deleteWorkflowArtifacts() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
  const client = new WebSiteManagementClient(credential, subscriptionId);
  await client.webApps.deployWorkflowArtifacts("testrg123", "testsite2", {
    workflowArtifacts: { filesToDelete: ["test/workflow.json", "test/"] },
  });
}
