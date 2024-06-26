const { AzureAnalysisServices } = require("@azure/arm-analysisservices");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List the result of the specified operation.
 *
 * @summary List the result of the specified operation.
 * x-ms-original-file: specification/analysisservices/resource-manager/Microsoft.AnalysisServices/stable/2017-08-01/examples/operationResults.json
 */
async function getDetailsOfAServer() {
  const subscriptionId = "613192d7-503f-477a-9cfe-4efc3ee2bd60";
  const location = "West US";
  const operationId = "00000000000000000000000000000000";
  const credential = new DefaultAzureCredential();
  const client = new AzureAnalysisServices(credential, subscriptionId);
  const result = await client.servers.listOperationResults(location, operationId);
  console.log(result);
}

getDetailsOfAServer().catch(console.error);
