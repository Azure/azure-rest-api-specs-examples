const { HelpRP } = require("@azure/arm-selfhelp");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Diagnostics tells you precisely the root cause of the issue and how to address it. You can get diagnostics once you discover and identify the relevant solution for your Azure issue.<br/><br/> You can create diagnostics using the ‘solutionId’  from Solution Discovery API response and ‘additionalParameters’ <br/><br/> <b>Note: </b>‘requiredParameterSets’ from Solutions Discovery API response must be passed via ‘additionalParameters’ as an input to Diagnostics API
 *
 * @summary Diagnostics tells you precisely the root cause of the issue and how to address it. You can get diagnostics once you discover and identify the relevant solution for your Azure issue.<br/><br/> You can create diagnostics using the ‘solutionId’  from Solution Discovery API response and ‘additionalParameters’ <br/><br/> <b>Note: </b>‘requiredParameterSets’ from Solutions Discovery API response must be passed via ‘additionalParameters’ as an input to Diagnostics API
 * x-ms-original-file: specification/help/resource-manager/Microsoft.Help/stable/2023-06-01/examples/CreateDiagnosticForKeyVaultResource.json
 */
async function createsADiagnosticForAKeyVaultResource() {
  const scope =
    "subscriptions/0d0fcd2e-c4fd-4349-8497-200edb3923c6/resourcegroups/myresourceGroup/providers/Microsoft.KeyVault/vaults/test-keyvault-non-read";
  const diagnosticsResourceName = "VMNotWorkingInsight";
  const credential = new DefaultAzureCredential();
  const client = new HelpRP(credential);
  const result = await client.diagnostics.beginCreateAndWait(scope, diagnosticsResourceName);
  console.log(result);
}
