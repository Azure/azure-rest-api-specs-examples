const { WebSiteManagementClient } = require("@azure/arm-appservice");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Gets a workflow version.
 *
 * @summary Gets a workflow version.
 * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2024-11-01/examples/WorkflowVersions_Get.json
 */
async function getAWorkflowVersion() {
  const subscriptionId =
    process.env["APPSERVICE_SUBSCRIPTION_ID"] || "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
  const resourceGroupName = process.env["APPSERVICE_RESOURCE_GROUP"] || "test-resource-group";
  const name = "test-name";
  const workflowName = "test-workflow";
  const versionId = "08586676824806722526";
  const credential = new DefaultAzureCredential();
  const client = new WebSiteManagementClient(credential, subscriptionId);
  const result = await client.workflowVersions.get(
    resourceGroupName,
    name,
    workflowName,
    versionId,
  );
  console.log(result);
}
