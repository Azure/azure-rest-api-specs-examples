const { WeightsAndBiasesClient } = require("@azure/arm-weightsandbiases");
const { DefaultAzureCredential } = require("@azure/identity");

async function operationsListGeneratedByMaximumSetRule() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-00000000000";
  const client = new WeightsAndBiasesClient(credential, subscriptionId);
  const resArray = new Array();
  for await (const item of client.operations.list()) {
    resArray.push(item);
  }

  console.log(resArray);
}
