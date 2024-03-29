const { AutomationClient } = require("@azure/arm-automation");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Update a certificate.
 *
 * @summary Update a certificate.
 * x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/preview/2020-01-13-preview/examples/updateCertificate.json
 */
async function updateACertificate() {
  const subscriptionId = process.env["AUTOMATION_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["AUTOMATION_RESOURCE_GROUP"] || "rg";
  const automationAccountName = "myAutomationAccount33";
  const certificateName = "testCert";
  const parameters = {
    name: "testCert",
    description: "sample certificate. Description updated",
  };
  const credential = new DefaultAzureCredential();
  const client = new AutomationClient(credential, subscriptionId);
  const result = await client.certificateOperations.update(
    resourceGroupName,
    automationAccountName,
    certificateName,
    parameters
  );
  console.log(result);
}
