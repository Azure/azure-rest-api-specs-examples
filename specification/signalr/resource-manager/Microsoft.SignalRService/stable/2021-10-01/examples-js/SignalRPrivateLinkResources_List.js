const { SignalRManagementClient } = require("@azure/arm-signalr");
const { DefaultAzureCredential } = require("@azure/identity");

async function signalRPrivateLinkResourcesList() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "myResourceGroup";
  const resourceName = "mySignalRService";
  const credential = new DefaultAzureCredential();
  const client = new SignalRManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.signalRPrivateLinkResources.list(resourceGroupName, resourceName)) {
    resArray.push(item);
  }
  console.log(resArray);
}

signalRPrivateLinkResourcesList().catch(console.error);
