const { MachineLearningWorkspacesManagementClient } = require("@azure/arm-workspaces");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists all the available machine learning workspaces under the specified subscription.
 *
 * @summary Lists all the available machine learning workspaces under the specified subscription.
 * x-ms-original-file: specification/machinelearning/resource-manager/Microsoft.MachineLearning/stable/2019-10-01/examples/ListWorkspaces.json
 */
async function workspaceGetBySubscription() {
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const credential = new DefaultAzureCredential();
  const client = new MachineLearningWorkspacesManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.workspaces.list()) {
    resArray.push(item);
  }
  console.log(resArray);
}

workspaceGetBySubscription().catch(console.error);
