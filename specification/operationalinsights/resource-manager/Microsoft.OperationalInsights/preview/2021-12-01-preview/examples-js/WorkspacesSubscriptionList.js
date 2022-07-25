const { OperationalInsightsManagementClient } = require("@azure/arm-operationalinsights");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets recently deleted workspaces in a subscription, available for recovery.
 *
 * @summary Gets recently deleted workspaces in a subscription, available for recovery.
 * x-ms-original-file: specification/operationalinsights/resource-manager/Microsoft.OperationalInsights/preview/2021-12-01-preview/examples/WorkspacesSubscriptionList.json
 */
async function workspacesSubscriptionList() {
  const subscriptionId = "00000000-0000-0000-0000-00000000000";
  const credential = new DefaultAzureCredential();
  const client = new OperationalInsightsManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.deletedWorkspaces.list()) {
    resArray.push(item);
  }
  console.log(resArray);
}

workspacesSubscriptionList().catch(console.error);
