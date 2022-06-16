const { OperationalInsightsManagementClient } = require("@azure/arm-operationalinsights");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the workspaces in a subscription.
 *
 * @summary Gets the workspaces in a subscription.
 * x-ms-original-file: specification/operationalinsights/resource-manager/Microsoft.OperationalInsights/stable/2021-06-01/examples/WorkspacesSubscriptionList.json
 */
async function workspacesSubscriptionList() {
  const subscriptionId = "00000000-0000-0000-0000-00000000000";
  const credential = new DefaultAzureCredential();
  const client = new OperationalInsightsManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.workspaces.list()) {
    resArray.push(item);
  }
  console.log(resArray);
}

workspacesSubscriptionList().catch(console.error);
