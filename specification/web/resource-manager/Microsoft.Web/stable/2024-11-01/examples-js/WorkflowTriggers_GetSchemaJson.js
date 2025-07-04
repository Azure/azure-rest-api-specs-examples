const { WebSiteManagementClient } = require("@azure/arm-appservice");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Get the trigger schema as JSON.
 *
 * @summary Get the trigger schema as JSON.
 * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2024-11-01/examples/WorkflowTriggers_GetSchemaJson.json
 */
async function getTriggerSchema() {
  const subscriptionId =
    process.env["APPSERVICE_SUBSCRIPTION_ID"] || "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
  const resourceGroupName = process.env["APPSERVICE_RESOURCE_GROUP"] || "testResourceGroup";
  const name = "test-name";
  const workflowName = "testWorkflow";
  const triggerName = "testTrigger";
  const credential = new DefaultAzureCredential();
  const client = new WebSiteManagementClient(credential, subscriptionId);
  const result = await client.workflowTriggers.getSchemaJson(
    resourceGroupName,
    name,
    workflowName,
    triggerName,
  );
  console.log(result);
}
