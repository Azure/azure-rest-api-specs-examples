const { AutomationClient } = require("@azure/arm-automation");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Retrieve the connection type identified by connection type name.
 *
 * @summary Retrieve the connection type identified by connection type name.
 * x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/preview/2020-01-13-preview/examples/getConnectionType.json
 */
async function getConnectionType() {
  const subscriptionId = process.env["AUTOMATION_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["AUTOMATION_RESOURCE_GROUP"] || "rg";
  const automationAccountName = "myAutomationAccount22";
  const connectionTypeName = "myCT";
  const credential = new DefaultAzureCredential();
  const client = new AutomationClient(credential, subscriptionId);
  const result = await client.connectionTypeOperations.get(
    resourceGroupName,
    automationAccountName,
    connectionTypeName
  );
  console.log(result);
}
