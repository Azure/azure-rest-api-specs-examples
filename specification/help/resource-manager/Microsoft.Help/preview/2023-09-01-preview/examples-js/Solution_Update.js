const { HelpRP } = require("@azure/arm-selfhelp");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Update the requiredInputs or additional information needed to execute the solution
 *
 * @summary Update the requiredInputs or additional information needed to execute the solution
 * x-ms-original-file: specification/help/resource-manager/Microsoft.Help/preview/2023-09-01-preview/examples/Solution_Update.json
 */
async function solutionUpdate() {
  const scope =
    "subscriptions/mySubscription/resourcegroups/myresourceGroup/providers/Microsoft.KeyVault/vaults/test-keyvault-rp";
  const solutionResourceName = "SolutionResourceName1";
  const solutionPatchRequestBody = {
    triggerCriteria: [
      {
        name: "ReplacementKey",
        value: "<!--12345678-BBBb-cCCCC-0000-123456789012-->",
      },
    ],
  };
  const options = { solutionPatchRequestBody };
  const credential = new DefaultAzureCredential();
  const client = new HelpRP(credential);
  const result = await client.solution.beginUpdateAndWait(scope, solutionResourceName, options);
  console.log(result);
}
