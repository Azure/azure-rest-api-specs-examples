const { PowerBIDedicated } = require("@azure/arm-powerbidedicated");
const { DefaultAzureCredential } = require("@azure/identity");

async function listAutoScaleVCoresInSubscription() {
  const subscriptionId = "613192d7-503f-477a-9cfe-4efc3ee2bd60";
  const credential = new DefaultAzureCredential();
  const client = new PowerBIDedicated(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.autoScaleVCores.listBySubscription()) {
    resArray.push(item);
  }
  console.log(resArray);
}

listAutoScaleVCoresInSubscription().catch(console.error);
