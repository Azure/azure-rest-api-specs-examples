const { MicrosoftStorageSync } = require("@azure/arm-storagesync");
const { DefaultAzureCredential } = require("@azure/identity");

async function privateEndpointConnectionsGet() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "res6977";
  const storageSyncServiceName = "sss2527";
  const privateEndpointConnectionName = "{privateEndpointConnectionName}";
  const credential = new DefaultAzureCredential();
  const client = new MicrosoftStorageSync(credential, subscriptionId);
  const result = await client.privateEndpointConnections.get(
    resourceGroupName,
    storageSyncServiceName,
    privateEndpointConnectionName
  );
  console.log(result);
}

privateEndpointConnectionsGet().catch(console.error);
