const { AutomationClient } = require("@azure/arm-automation");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Retrieve the source control identified by source control name.
 *
 * @summary Retrieve the source control identified by source control name.
 * x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/preview/2020-01-13-preview/examples/sourceControl/getSourceControl.json
 */
async function getASourceControl() {
  const subscriptionId = process.env["AUTOMATION_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["AUTOMATION_RESOURCE_GROUP"] || "rg";
  const automationAccountName = "sampleAccount9";
  const sourceControlName = "sampleSourceControl";
  const credential = new DefaultAzureCredential();
  const client = new AutomationClient(credential, subscriptionId);
  const result = await client.sourceControlOperations.get(
    resourceGroupName,
    automationAccountName,
    sourceControlName
  );
  console.log(result);
}
