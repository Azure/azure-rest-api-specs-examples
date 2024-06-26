const { AzureOrbital } = require("@azure/arm-orbital");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Returns operation results.
 *
 * @summary Returns operation results.
 * x-ms-original-file: specification/orbital/resource-manager/Microsoft.Orbital/stable/2022-11-01/examples/OperationResultsGet.json
 */
async function kustoOperationResultsGet() {
  const subscriptionId =
    process.env["ORBITAL_SUBSCRIPTION_ID"] || "12345678-1234-1234-1234-123456789098";
  const location = "eastus2";
  const operationId = "30972f1b-b61d-4fd8-bd34-3dcfa24670f3";
  const credential = new DefaultAzureCredential();
  const client = new AzureOrbital(credential, subscriptionId);
  const result = await client.operationsResults.beginGetAndWait(location, operationId);
  console.log(result);
}
