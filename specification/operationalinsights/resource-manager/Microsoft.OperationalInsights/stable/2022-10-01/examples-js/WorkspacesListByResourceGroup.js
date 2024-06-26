const { OperationalInsightsManagementClient } = require("@azure/arm-operationalinsights");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets recently deleted workspaces in a resource group, available for recovery.
 *
 * @summary Gets recently deleted workspaces in a resource group, available for recovery.
 * x-ms-original-file: specification/operationalinsights/resource-manager/Microsoft.OperationalInsights/stable/2022-10-01/examples/WorkspacesListByResourceGroup.json
 */
async function workspacesGet() {
  const subscriptionId =
    process.env["OPERATIONALINSIGHTS_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-00000000000";
  const resourceGroupName = process.env["OPERATIONALINSIGHTS_RESOURCE_GROUP"] || "oiautorest6685";
  const credential = new DefaultAzureCredential();
  const client = new OperationalInsightsManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.deletedWorkspaces.listByResourceGroup(resourceGroupName)) {
    resArray.push(item);
  }
  console.log(resArray);
}
