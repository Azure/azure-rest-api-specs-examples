const { OperationalInsightsManagementClient } = require("@azure/arm-operationalinsights");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the available service tiers for the workspace.
 *
 * @summary Gets the available service tiers for the workspace.
 * x-ms-original-file: specification/operationalinsights/resource-manager/Microsoft.OperationalInsights/stable/2020-08-01/examples/WorkspacesAvailableServiceTiers.json
 */
async function availableServiceTiers() {
  const subscriptionId = "00000000-0000-0000-0000-00000000000";
  const resourceGroupName = "rg1";
  const workspaceName = "workspace1";
  const credential = new DefaultAzureCredential();
  const client = new OperationalInsightsManagementClient(credential, subscriptionId);
  const result = await client.availableServiceTiers.listByWorkspace(
    resourceGroupName,
    workspaceName
  );
  console.log(result);
}

availableServiceTiers().catch(console.error);
