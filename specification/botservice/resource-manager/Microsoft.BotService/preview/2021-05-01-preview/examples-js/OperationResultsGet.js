const { AzureBotService } = require("@azure/arm-botservice");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get the operation result for a long running operation.
 *
 * @summary Get the operation result for a long running operation.
 * x-ms-original-file: specification/botservice/resource-manager/Microsoft.BotService/preview/2021-05-01-preview/examples/OperationResultsGet.json
 */
async function getOperationResult() {
  const subscriptionId = "subid";
  const operationResultId = "exampleid";
  const credential = new DefaultAzureCredential();
  const client = new AzureBotService(credential, subscriptionId);
  const result = await client.operationResults.beginGetAndWait(operationResultId);
  console.log(result);
}

getOperationResult().catch(console.error);
