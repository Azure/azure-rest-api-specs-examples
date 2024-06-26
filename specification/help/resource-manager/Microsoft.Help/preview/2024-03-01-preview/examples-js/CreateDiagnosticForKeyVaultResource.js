const { HelpRP } = require("@azure/arm-selfhelp");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates a diagnostic for the specific resource using solutionId from discovery solutions. <br/>Diagnostics are powerful solutions that access product resources or other relevant data and provide the root cause of the issue and the steps to address the issue.<br/><br/>
 *
 * @summary Creates a diagnostic for the specific resource using solutionId from discovery solutions. <br/>Diagnostics are powerful solutions that access product resources or other relevant data and provide the root cause of the issue and the steps to address the issue.<br/><br/>
 * x-ms-original-file: specification/help/resource-manager/Microsoft.Help/preview/2024-03-01-preview/examples/CreateDiagnosticForKeyVaultResource.json
 */
async function createsADiagnosticForAKeyVaultResource() {
  const scope =
    "subscriptions/0d0fcd2e-c4fd-4349-8497-200edb3923c6/resourceGroups/myresourceGroup/providers/Microsoft.KeyVault/vaults/test-keyvault-non-read";
  const diagnosticsResourceName = "VMNotWorkingInsight";
  const credential = new DefaultAzureCredential();
  const client = new HelpRP(credential);
  const result = await client.diagnostics.beginCreateAndWait(scope, diagnosticsResourceName);
  console.log(result);
}
