const { OperationalInsightsManagementClient } = require("@azure/arm-operationalinsights");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Disables an intelligence pack for a given workspace.
 *
 * @summary Disables an intelligence pack for a given workspace.
 * x-ms-original-file: specification/operationalinsights/resource-manager/Microsoft.OperationalInsights/stable/2020-08-01/examples/WorkspacesDisableIntelligencePack.json
 */
async function intelligencePacksDisable() {
  const subscriptionId =
    process.env["OPERATIONALINSIGHTS_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-00000000000";
  const resourceGroupName = process.env["OPERATIONALINSIGHTS_RESOURCE_GROUP"] || "rg1";
  const workspaceName = "TestLinkWS";
  const intelligencePackName = "ChangeTracking";
  const credential = new DefaultAzureCredential();
  const client = new OperationalInsightsManagementClient(credential, subscriptionId);
  const result = await client.intelligencePacks.disable(
    resourceGroupName,
    workspaceName,
    intelligencePackName
  );
  console.log(result);
}
