const { AlertRuleRecommendationsManagementClient } = require("@azure/arm-alertrulerecommendations");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to retrieve alert rule recommendations for a resource.
 *
 * @summary retrieve alert rule recommendations for a resource.
 * x-ms-original-file: 2023-08-01-preview/AlertRuleRecommendations_GetByResource_VM.json
 */
async function listAlertRuleRecommendationsForVirtualMachinesAtResourceLevel() {
  const credential = new DefaultAzureCredential();
  const client = new AlertRuleRecommendationsManagementClient(credential);
  const resArray = new Array();
  for await (const item of client.alertRuleRecommendations.listByResource(
    "subscriptions/2f00cc51-6809-498f-9ffc-48c42aff570d/resourcegroups/test/providers/Microsoft.Compute/virtualMachines/testMachineCanBeSafelyDeleted",
  )) {
    resArray.push(item);
  }

  console.log(resArray);
}
