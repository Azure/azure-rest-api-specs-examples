const { HelpRP } = require("@azure/arm-selfhelp");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates a diagnostic for the specific resource using solutionId and requiredInputs* from discovery solutions. <br/>Diagnostics are powerful solutions that access product resources or other relevant data and provide the root cause of the issue and the steps to address the issue.<br/><br/> <b>Note: </b> ‘requiredInputs’ from Discovery solutions response must be passed via ‘additionalParameters’ as an input to Diagnostics API.
 *
 * @summary Creates a diagnostic for the specific resource using solutionId and requiredInputs* from discovery solutions. <br/>Diagnostics are powerful solutions that access product resources or other relevant data and provide the root cause of the issue and the steps to address the issue.<br/><br/> <b>Note: </b> ‘requiredInputs’ from Discovery solutions response must be passed via ‘additionalParameters’ as an input to Diagnostics API.
 * x-ms-original-file: specification/help/resource-manager/Microsoft.Help/preview/2023-09-01-preview/examples/CreateDiagnosticForKeyVaultResource.json
 */
async function createsADiagnosticForAKeyVaultResource() {
  const scope =
    "subscriptions/0d0fcd2e-c4fd-4349-8497-200edb3923c6/resourceGroups/myresourceGroup/providers/Microsoft.KeyVault/vaults/test-keyvault-non-read";
  const diagnosticsResourceName = "VMNotWorkingInsight";
  const diagnosticResourceRequest = {
    insights: [{ solutionId: "SampleSolutionId1" }, { solutionId: "SampleSolutionId2" }],
  };
  const options = {
    diagnosticResourceRequest,
  };
  const credential = new DefaultAzureCredential();
  const client = new HelpRP(credential);
  const result = await client.diagnostics.beginCreateAndWait(
    scope,
    diagnosticsResourceName,
    options,
  );
  console.log(result);
}
