const { AutomationClient } = require("@azure/arm-automation");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Retrieve the Dsc node report list by node id.
 *
 * @summary Retrieve the Dsc node report list by node id.
 * x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/preview/2020-01-13-preview/examples/listPagedDscNodeReportsByNode.json
 */
async function listPagedDscReportsByNodeId() {
  const subscriptionId = process.env["AUTOMATION_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["AUTOMATION_RESOURCE_GROUP"] || "rg";
  const automationAccountName = "myAutomationAccount33";
  const nodeId = "nodeId";
  const credential = new DefaultAzureCredential();
  const client = new AutomationClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.nodeReports.listByNode(
    resourceGroupName,
    automationAccountName,
    nodeId
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}
