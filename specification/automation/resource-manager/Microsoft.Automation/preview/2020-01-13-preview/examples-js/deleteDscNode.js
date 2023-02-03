const { AutomationClient } = require("@azure/arm-automation");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Delete the dsc node identified by node id.
 *
 * @summary Delete the dsc node identified by node id.
 * x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/preview/2020-01-13-preview/examples/deleteDscNode.json
 */
async function deleteADscNode() {
  const subscriptionId = process.env["AUTOMATION_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["AUTOMATION_RESOURCE_GROUP"] || "rg";
  const automationAccountName = "myAutomationAccount9";
  const nodeId = "e1243a76-a9bd-432f-bde3-ad8f317ee786";
  const credential = new DefaultAzureCredential();
  const client = new AutomationClient(credential, subscriptionId);
  const result = await client.dscNodeOperations.delete(
    resourceGroupName,
    automationAccountName,
    nodeId
  );
  console.log(result);
}
