const { HelpRP } = require("@azure/arm-selfhelp");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get the solution using the applicable solutionResourceName while creating the solution.
 *
 * @summary Get the solution using the applicable solutionResourceName while creating the solution.
 * x-ms-original-file: specification/help/resource-manager/Microsoft.Help/preview/2023-09-01-preview/examples/Solution_Get.json
 */
async function solutionGet() {
  const scope =
    "subscriptions/mySubscription/resourcegroups/myresourceGroup/providers/Microsoft.KeyVault/vaults/test-keyvault-rp";
  const solutionResourceName = "SolutionResource1";
  const credential = new DefaultAzureCredential();
  const client = new HelpRP(credential);
  const result = await client.solution.get(scope, solutionResourceName);
  console.log(result);
}
