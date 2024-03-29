const { AutomationClient } = require("@azure/arm-automation");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create or Update the python 2 package identified by package name.
 *
 * @summary Create or Update the python 2 package identified by package name.
 * x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/preview/2020-01-13-preview/examples/createOrUpdatePython2Package.json
 */
async function createOrUpdateAPython2Package() {
  const subscriptionId = process.env["AUTOMATION_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["AUTOMATION_RESOURCE_GROUP"] || "rg";
  const automationAccountName = "myAutomationAccount33";
  const packageName = "OmsCompositeResources";
  const parameters = {
    contentLink: {
      contentHash: {
        algorithm: "sha265",
        value: "07E108A962B81DD9C9BAA89BB47C0F6EE52B29E83758B07795E408D258B2B87A",
      },
      uri: "https://teststorage.blob.core.windows.net/dsccomposite/OmsCompositeResources.zip",
      version: "1.0.0.0",
    },
    tags: {},
  };
  const credential = new DefaultAzureCredential();
  const client = new AutomationClient(credential, subscriptionId);
  const result = await client.python2Package.createOrUpdate(
    resourceGroupName,
    automationAccountName,
    packageName,
    parameters
  );
  console.log(result);
}
