const { SignalRManagementClient } = require("@azure/arm-signalr");
const { DefaultAzureCredential } = require("@azure/identity");

async function signalRRegenerateKey() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "myResourceGroup";
  const resourceName = "mySignalRService";
  const parameters = { keyType: "Primary" };
  const credential = new DefaultAzureCredential();
  const client = new SignalRManagementClient(credential, subscriptionId);
  const result = await client.signalR.beginRegenerateKeyAndWait(
    resourceGroupName,
    resourceName,
    parameters
  );
  console.log(result);
}

signalRRegenerateKey().catch(console.error);
