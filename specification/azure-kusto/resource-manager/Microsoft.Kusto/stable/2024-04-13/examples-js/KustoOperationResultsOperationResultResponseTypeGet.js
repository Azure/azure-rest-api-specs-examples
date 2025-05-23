const { KustoManagementClient } = require("@azure/arm-kusto");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Returns operation results.
 *
 * @summary Returns operation results.
 * x-ms-original-file: specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2024-04-13/examples/KustoOperationResultsOperationResultResponseTypeGet.json
 */
async function kustoOperationsResultsLocationGet() {
  const subscriptionId =
    process.env["KUSTO_SUBSCRIPTION_ID"] || "12345678-1234-1234-1234-123456789098";
  const location = "westus";
  const operationId = "30972f1b-b61d-4fd8-bd34-3dcfa24670f3";
  const credential = new DefaultAzureCredential();
  const client = new KustoManagementClient(credential, subscriptionId);
  const result = await client.operationsResultsLocation.get(location, operationId);
  console.log(result);
}
