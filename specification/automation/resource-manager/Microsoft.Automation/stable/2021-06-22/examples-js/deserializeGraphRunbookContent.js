const { AutomationClient } = require("@azure/arm-automation");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Post operation to serialize or deserialize GraphRunbookContent
 *
 * @summary Post operation to serialize or deserialize GraphRunbookContent
 * x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/stable/2021-06-22/examples/deserializeGraphRunbookContent.json
 */
async function getGraphicalRunbookContentFromRawContent() {
  const subscriptionId = process.env["AUTOMATION_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["AUTOMATION_RESOURCE_GROUP"] || "rg";
  const automationAccountName = "MyAutomationAccount";
  const parameters = {
    rawContent: {
      runbookDefinition:
        "AAEAAADAQAAAAAAAAAMAgAAAGJPcmNoZXN0cmF0b3IuR3JhcGhSdW5ib29rLk1vZGVsLCBWZXJzaW9uPTcuMy4wLjAsIEN1bHR....",
      runbookType: "GraphPowerShell",
      schemaVersion: "1.10",
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new AutomationClient(credential, subscriptionId);
  const result = await client.convertGraphRunbookContent(
    resourceGroupName,
    automationAccountName,
    parameters
  );
  console.log(result);
}
