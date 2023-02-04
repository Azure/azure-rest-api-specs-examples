const { AutomationClient } = require("@azure/arm-automation");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Delete the source control.
 *
 * @summary Delete the source control.
 * x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/preview/2020-01-13-preview/examples/sourceControl/deleteSourceControl.json
 */
async function deleteASourceControl() {
  const subscriptionId = process.env["AUTOMATION_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["AUTOMATION_RESOURCE_GROUP"] || "rg";
  const automationAccountName = "sampleAccount9";
  const sourceControlName = "sampleSourceControl";
  const credential = new DefaultAzureCredential();
  const client = new AutomationClient(credential, subscriptionId);
  const result = await client.sourceControlOperations.delete(
    resourceGroupName,
    automationAccountName,
    sourceControlName
  );
  console.log(result);
}
