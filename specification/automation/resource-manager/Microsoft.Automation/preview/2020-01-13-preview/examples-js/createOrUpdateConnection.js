const { AutomationClient } = require("@azure/arm-automation");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create or update a connection.
 *
 * @summary Create or update a connection.
 * x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/preview/2020-01-13-preview/examples/createOrUpdateConnection.json
 */
async function createOrUpdateConnection() {
  const subscriptionId = process.env["AUTOMATION_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["AUTOMATION_RESOURCE_GROUP"] || "rg";
  const automationAccountName = "myAutomationAccount28";
  const connectionName = "mysConnection";
  const parameters = {
    name: "mysConnection",
    description: "my description goes here",
    connectionType: { name: "Azure" },
    fieldDefinitionValues: {
      automationCertificateName: "mysCertificateName",
      subscriptionID: "subid",
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new AutomationClient(credential, subscriptionId);
  const result = await client.connectionOperations.createOrUpdate(
    resourceGroupName,
    automationAccountName,
    connectionName,
    parameters
  );
  console.log(result);
}
