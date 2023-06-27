const { HelpRP } = require("@azure/arm-selfhelp");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to This API is used to check the uniqueness of a resource name used for a diagnostic check.
 *
 * @summary This API is used to check the uniqueness of a resource name used for a diagnostic check.
 * x-ms-original-file: specification/help/resource-manager/Microsoft.Help/stable/2023-06-01/examples/CheckNameAvailabilityForDiagnosticWhenNameIsNotAvailable.json
 */
async function exampleWhenNameIsNotAvailableForADiagnosticResource() {
  const scope = "subscriptions/0d0fcd2e-c4fd-4349-8497-200edb3923c6";
  const checkNameAvailabilityRequest = {
    name: "sampleName",
    type: "Microsoft.Help/diagnostics",
  };
  const options = {
    checkNameAvailabilityRequest,
  };
  const credential = new DefaultAzureCredential();
  const client = new HelpRP(credential);
  const result = await client.diagnostics.checkNameAvailability(scope, options);
  console.log(result);
}
