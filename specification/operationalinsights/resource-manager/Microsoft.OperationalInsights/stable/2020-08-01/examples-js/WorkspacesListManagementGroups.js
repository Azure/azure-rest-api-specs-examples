const { OperationalInsightsManagementClient } = require("@azure/arm-operationalinsights");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a list of management groups connected to a workspace.
 *
 * @summary Gets a list of management groups connected to a workspace.
 * x-ms-original-file: specification/operationalinsights/resource-manager/Microsoft.OperationalInsights/stable/2020-08-01/examples/WorkspacesListManagementGroups.json
 */
async function workspacesListManagementGroups() {
  const subscriptionId = "00000000-0000-0000-0000-00000000000";
  const resourceGroupName = "rg1";
  const workspaceName = "TestLinkWS";
  const credential = new DefaultAzureCredential();
  const client = new OperationalInsightsManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.managementGroups.list(resourceGroupName, workspaceName)) {
    resArray.push(item);
  }
  console.log(resArray);
}

workspacesListManagementGroups().catch(console.error);
