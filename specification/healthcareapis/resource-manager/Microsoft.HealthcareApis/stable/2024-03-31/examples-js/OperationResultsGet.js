const { HealthcareApisManagementClient } = require("@azure/arm-healthcareapis");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get the operation result for a long running operation.
 *
 * @summary Get the operation result for a long running operation.
 * x-ms-original-file: specification/healthcareapis/resource-manager/Microsoft.HealthcareApis/stable/2024-03-31/examples/OperationResultsGet.json
 */
async function getOperationResult() {
  const subscriptionId = process.env["HEALTHCAREAPIS_SUBSCRIPTION_ID"] || "subid";
  const locationName = "westus";
  const operationResultId = "exampleid";
  const credential = new DefaultAzureCredential();
  const client = new HealthcareApisManagementClient(credential, subscriptionId);
  const result = await client.operationResults.get(locationName, operationResultId);
  console.log(result);
}
