const { AutomationClient } = require("@azure/arm-automation");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Retrieve the Dsc node reports by node id and report id.
 *
 * @summary Retrieve the Dsc node reports by node id and report id.
 * x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/preview/2020-01-13-preview/examples/getDscNodeReportContent.json
 */
async function getContentOfNode() {
  const subscriptionId = process.env["AUTOMATION_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["AUTOMATION_RESOURCE_GROUP"] || "rg";
  const automationAccountName = "myAutomationAccount33";
  const nodeId = "nodeId";
  const reportId = "reportId";
  const credential = new DefaultAzureCredential();
  const client = new AutomationClient(credential, subscriptionId);
  const result = await client.nodeReports.getContent(
    resourceGroupName,
    automationAccountName,
    nodeId,
    reportId
  );
  console.log(result);
}
