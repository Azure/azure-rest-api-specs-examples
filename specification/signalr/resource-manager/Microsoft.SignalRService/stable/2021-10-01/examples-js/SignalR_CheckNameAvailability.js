const { SignalRManagementClient } = require("@azure/arm-signalr");
const { DefaultAzureCredential } = require("@azure/identity");

async function signalRCheckNameAvailability() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const location = "eastus";
  const parameters = {
    name: "mySignalRService",
    type: "Microsoft.SignalRService/SignalR",
  };
  const credential = new DefaultAzureCredential();
  const client = new SignalRManagementClient(credential, subscriptionId);
  const result = await client.signalR.checkNameAvailability(location, parameters);
  console.log(result);
}

signalRCheckNameAvailability().catch(console.error);
