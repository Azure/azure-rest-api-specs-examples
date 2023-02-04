const { AutomationClient } = require("@azure/arm-automation");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates the Dsc compilation job of the configuration.
 *
 * @summary Creates the Dsc compilation job of the configuration.
 * x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/preview/2020-01-13-preview/examples/createCompilationJob.json
 */
async function createOrUpdateADscCompilationJob() {
  const subscriptionId = process.env["AUTOMATION_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["AUTOMATION_RESOURCE_GROUP"] || "rg";
  const automationAccountName = "myAutomationAccount33";
  const compilationJobName = "TestCompilationJob";
  const parameters = {
    configuration: { name: "SetupServer" },
  };
  const credential = new DefaultAzureCredential();
  const client = new AutomationClient(credential, subscriptionId);
  const result = await client.dscCompilationJobOperations.beginCreateAndWait(
    resourceGroupName,
    automationAccountName,
    compilationJobName,
    parameters
  );
  console.log(result);
}
