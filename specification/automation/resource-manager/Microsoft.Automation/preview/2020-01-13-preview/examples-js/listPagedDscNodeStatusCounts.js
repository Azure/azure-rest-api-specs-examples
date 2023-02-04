const { AutomationClient } = require("@azure/arm-automation");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Retrieve counts for Dsc Nodes.
 *
 * @summary Retrieve counts for Dsc Nodes.
 * x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/preview/2020-01-13-preview/examples/listPagedDscNodeStatusCounts.json
 */
async function getNodeStatusCounts() {
  const subscriptionId = process.env["AUTOMATION_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["AUTOMATION_RESOURCE_GROUP"] || "rg";
  const automationAccountName = "myAutomationAccount33";
  const countType = "status";
  const credential = new DefaultAzureCredential();
  const client = new AutomationClient(credential, subscriptionId);
  const result = await client.nodeCountInformation.get(
    resourceGroupName,
    automationAccountName,
    countType
  );
  console.log(result);
}
