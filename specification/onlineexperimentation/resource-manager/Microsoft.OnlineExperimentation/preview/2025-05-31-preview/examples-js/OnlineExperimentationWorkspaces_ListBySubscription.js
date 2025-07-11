const { OnlineExperimentationClient } = require("@azure/arm-onlineexperimentation");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to gets all online experimentation workspaces in the specified subscription.
 *
 * @summary gets all online experimentation workspaces in the specified subscription.
 * x-ms-original-file: 2025-05-31-preview/OnlineExperimentationWorkspaces_ListBySubscription.json
 */
async function listOnlineExperimentationWorkspacesInASubscription() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "fa5fc227-a624-475e-b696-cdd604c735bc";
  const client = new OnlineExperimentationClient(credential, subscriptionId);
  const resArray = new Array();
  for await (const item of client.onlineExperimentationWorkspaces.listBySubscription()) {
    resArray.push(item);
  }

  console.log(resArray);
}
