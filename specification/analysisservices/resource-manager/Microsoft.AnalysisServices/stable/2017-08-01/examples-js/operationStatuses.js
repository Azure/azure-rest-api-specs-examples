const { AzureAnalysisServices } = require("@azure/arm-analysisservices");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List the status of operation.
 *
 * @summary List the status of operation.
 * x-ms-original-file: specification/analysisservices/resource-manager/Microsoft.AnalysisServices/stable/2017-08-01/examples/operationStatuses.json
 */
async function getDetailsOfAServer() {
  const subscriptionId = "613192d7-503f-477a-9cfe-4efc3ee2bd60";
  const location = "West US";
  const operationId = "testoperationid";
  const credential = new DefaultAzureCredential();
  const client = new AzureAnalysisServices(credential, subscriptionId);
  const result = await client.servers.listOperationStatuses(location, operationId);
  console.log(result);
}

getDetailsOfAServer().catch(console.error);
