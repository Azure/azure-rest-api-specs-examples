const { SecurityCenter } = require("@azure/arm-security");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Settings about where we should store your security data and logs. If the result is empty, it means that no custom-workspace configuration was set
 *
 * @summary Settings about where we should store your security data and logs. If the result is empty, it means that no custom-workspace configuration was set
 * x-ms-original-file: specification/security/resource-manager/Microsoft.Security/preview/2017-08-01-preview/examples/WorkspaceSettings/GetWorkspaceSettings_example.json
 */
async function getWorkspaceSettingsOnSubscription() {
  const subscriptionId =
    process.env["SECURITY_SUBSCRIPTION_ID"] || "20ff7fc3-e762-44dd-bd96-b71116dcdc23";
  const credential = new DefaultAzureCredential();
  const client = new SecurityCenter(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.workspaceSettings.list()) {
    resArray.push(item);
  }
  console.log(resArray);
}
