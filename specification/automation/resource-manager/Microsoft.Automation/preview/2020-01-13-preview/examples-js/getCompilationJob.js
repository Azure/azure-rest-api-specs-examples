const { AutomationClient } = require("@azure/arm-automation");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Retrieve the Dsc configuration compilation job identified by job id.
 *
 * @summary Retrieve the Dsc configuration compilation job identified by job id.
 * x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/preview/2020-01-13-preview/examples/getCompilationJob.json
 */
async function getADscCompilationJob() {
  const subscriptionId = process.env["AUTOMATION_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["AUTOMATION_RESOURCE_GROUP"] || "rg";
  const automationAccountName = "myAutomationAccount33";
  const compilationJobName = "TestCompilationJob";
  const credential = new DefaultAzureCredential();
  const client = new AutomationClient(credential, subscriptionId);
  const result = await client.dscCompilationJobOperations.get(
    resourceGroupName,
    automationAccountName,
    compilationJobName
  );
  console.log(result);
}
