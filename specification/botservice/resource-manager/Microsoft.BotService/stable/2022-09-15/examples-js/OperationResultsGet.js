const { AzureBotService } = require("@azure/arm-botservice");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get the operation result for a long running operation.
 *
 * @summary Get the operation result for a long running operation.
 * x-ms-original-file: specification/botservice/resource-manager/Microsoft.BotService/stable/2022-09-15/examples/OperationResultsGet.json
 */
async function getOperationResult() {
  const subscriptionId = process.env["BOTSERVICE_SUBSCRIPTION_ID"] || "subid";
  const operationResultId = "exampleid";
  const credential = new DefaultAzureCredential();
  const client = new AzureBotService(credential, subscriptionId);
  const result = await client.operationResults.beginGetAndWait(operationResultId);
  console.log(result);
}
