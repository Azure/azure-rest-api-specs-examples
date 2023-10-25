const { HelpRP } = require("@azure/arm-selfhelp");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get the diagnostics using the 'diagnosticsResourceName' you chose while creating the diagnostic.
 *
 * @summary Get the diagnostics using the 'diagnosticsResourceName' you chose while creating the diagnostic.
 * x-ms-original-file: specification/help/resource-manager/Microsoft.Help/preview/2023-09-01-preview/examples/GetDiagnosticForKeyVaultResource.json
 */
async function getsADiagnosticForAKeyVaultResource() {
  const scope =
    "subscriptions/0d0fcd2e-c4fd-4349-8497-200edb3923c6/resourceGroups/myresourceGroup/providers/Microsoft.KeyVault/vaults/test-keyvault-non-read";
  const diagnosticsResourceName = "VMNotWorkingInsight";
  const credential = new DefaultAzureCredential();
  const client = new HelpRP(credential);
  const result = await client.diagnostics.get(scope, diagnosticsResourceName);
  console.log(result);
}
