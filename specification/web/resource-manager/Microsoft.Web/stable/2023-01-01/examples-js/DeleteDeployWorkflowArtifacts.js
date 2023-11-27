const { WebSiteManagementClient } = require("@azure/arm-appservice");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Description for Creates the artifacts for web site, or a deployment slot.
 *
 * @summary Description for Creates the artifacts for web site, or a deployment slot.
 * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2023-01-01/examples/DeleteDeployWorkflowArtifacts.json
 */
async function deleteWorkflowArtifacts() {
  const subscriptionId =
    process.env["APPSERVICE_SUBSCRIPTION_ID"] || "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
  const resourceGroupName = process.env["APPSERVICE_RESOURCE_GROUP"] || "testrg123";
  const name = "testsite2";
  const workflowArtifacts = {
    filesToDelete: ["test/workflow.json", "test/"],
  };
  const options = {
    workflowArtifacts,
  };
  const credential = new DefaultAzureCredential();
  const client = new WebSiteManagementClient(credential, subscriptionId);
  const result = await client.webApps.deployWorkflowArtifacts(resourceGroupName, name, options);
  console.log(result);
}
