const { AutomationClient } = require("@azure/arm-automation");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Update a connection.
 *
 * @summary Update a connection.
 * x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/preview/2020-01-13-preview/examples/updateConnection.json
 */
async function updateAConnection() {
  const subscriptionId = process.env["AUTOMATION_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["AUTOMATION_RESOURCE_GROUP"] || "rg";
  const automationAccountName = "myAutomationAccount28";
  const connectionName = "myConnection";
  const parameters = {
    name: "myConnection",
    description: "my description goes here",
    fieldDefinitionValues: {
      automationCertificateName: "myCertificateName",
      subscriptionID: "b5e4748c-f69a-467c-8749-e2f9c8cd3009",
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new AutomationClient(credential, subscriptionId);
  const result = await client.connectionOperations.update(
    resourceGroupName,
    automationAccountName,
    connectionName,
    parameters
  );
  console.log(result);
}
