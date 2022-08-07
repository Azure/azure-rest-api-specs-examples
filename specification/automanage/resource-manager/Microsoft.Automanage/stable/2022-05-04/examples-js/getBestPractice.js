const { AutomanageClient } = require("@azure/arm-automanage");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get information about a Automanage best practice
 *
 * @summary Get information about a Automanage best practice
 * x-ms-original-file: specification/automanage/resource-manager/Microsoft.Automanage/stable/2022-05-04/examples/getBestPractice.json
 */
async function getAnAutomanageBestPractice() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const bestPracticeName = "azureBestPracticesProduction";
  const credential = new DefaultAzureCredential();
  const client = new AutomanageClient(credential, subscriptionId);
  const result = await client.bestPractices.get(bestPracticeName);
  console.log(result);
}

getAnAutomanageBestPractice().catch(console.error);
