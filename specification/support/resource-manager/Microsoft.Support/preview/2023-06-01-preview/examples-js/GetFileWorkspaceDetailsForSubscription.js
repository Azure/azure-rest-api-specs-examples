const { MicrosoftSupport } = require("@azure/arm-support");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets details for a specific file workspace in an Azure subscription.
 *
 * @summary Gets details for a specific file workspace in an Azure subscription.
 * x-ms-original-file: specification/support/resource-manager/Microsoft.Support/preview/2023-06-01-preview/examples/GetFileWorkspaceDetailsForSubscription.json
 */
async function getDetailsOfASubscriptionFileWorkspace() {
  const subscriptionId = process.env["SUPPORT_SUBSCRIPTION_ID"] || "subid";
  const fileWorkspaceName = "testworkspace";
  const credential = new DefaultAzureCredential();
  const client = new MicrosoftSupport(credential, subscriptionId);
  const result = await client.fileWorkspaces.get(fileWorkspaceName);
  console.log(result);
}
