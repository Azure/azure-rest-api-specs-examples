const { OperationalInsightsManagementClient } = require("@azure/arm-operationalinsights");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists all the intelligence packs possible and whether they are enabled or disabled for a given workspace.
 *
 * @summary Lists all the intelligence packs possible and whether they are enabled or disabled for a given workspace.
 * x-ms-original-file: specification/operationalinsights/resource-manager/Microsoft.OperationalInsights/stable/2020-08-01/examples/WorkspacesListIntelligencePacks.json
 */
async function intelligencePacksList() {
  const subscriptionId = "00000000-0000-0000-0000-00000000000";
  const resourceGroupName = "rg1";
  const workspaceName = "TestLinkWS";
  const credential = new DefaultAzureCredential();
  const client = new OperationalInsightsManagementClient(credential, subscriptionId);
  const result = await client.intelligencePacks.list(resourceGroupName, workspaceName);
  console.log(result);
}

intelligencePacksList().catch(console.error);
