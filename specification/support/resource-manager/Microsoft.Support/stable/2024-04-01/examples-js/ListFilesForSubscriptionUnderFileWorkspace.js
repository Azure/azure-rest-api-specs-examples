const { MicrosoftSupport } = require("@azure/arm-support");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists all the Files information under a workspace for an Azure subscription.
 *
 * @summary Lists all the Files information under a workspace for an Azure subscription.
 * x-ms-original-file: specification/support/resource-manager/Microsoft.Support/stable/2024-04-01/examples/ListFilesForSubscriptionUnderFileWorkspace.json
 */
async function listFilesUnderAWorkspaceForASubscription() {
  const subscriptionId =
    process.env["SUPPORT_SUBSCRIPTION_ID"] || "132d901f-189d-4381-9214-fe68e27e05a1";
  const fileWorkspaceName = "testworkspace";
  const credential = new DefaultAzureCredential();
  const client = new MicrosoftSupport(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.files.list(fileWorkspaceName)) {
    resArray.push(item);
  }
  console.log(resArray);
}
