const { SignalRManagementClient } = require("@azure/arm-signalr");
const { DefaultAzureCredential } = require("@azure/identity");

async function usagesList() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const location = "eastus";
  const credential = new DefaultAzureCredential();
  const client = new SignalRManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.usages.list(location)) {
    resArray.push(item);
  }
  console.log(resArray);
}

usagesList().catch(console.error);
