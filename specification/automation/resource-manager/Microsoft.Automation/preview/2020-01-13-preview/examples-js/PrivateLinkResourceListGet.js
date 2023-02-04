const { AutomationClient } = require("@azure/arm-automation");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the private link resources that need to be created for Automation account.
 *
 * @summary Gets the private link resources that need to be created for Automation account.
 * x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/preview/2020-01-13-preview/examples/PrivateLinkResourceListGet.json
 */
async function getsPrivateEndpointConnection() {
  const subscriptionId =
    process.env["AUTOMATION_SUBSCRIPTION_ID"] || "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = process.env["AUTOMATION_RESOURCE_GROUP"] || "rg1";
  const automationAccountName = "ddb1";
  const credential = new DefaultAzureCredential();
  const client = new AutomationClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.privateLinkResources.listAutomation(
    resourceGroupName,
    automationAccountName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}
