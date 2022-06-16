const { MicrosoftStorageSync } = require("@azure/arm-storagesync");
const { DefaultAzureCredential } = require("@azure/identity");

async function privateEndpointConnectionsCreate() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "res7687";
  const storageSyncServiceName = "sss2527";
  const privateEndpointConnectionName = "{privateEndpointConnectionName}";
  const properties = {
    privateLinkServiceConnectionState: {
      description: "Auto-Approved",
      status: "Approved",
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new MicrosoftStorageSync(credential, subscriptionId);
  const result = await client.privateEndpointConnections.beginCreateAndWait(
    resourceGroupName,
    storageSyncServiceName,
    privateEndpointConnectionName,
    properties
  );
  console.log(result);
}

privateEndpointConnectionsCreate().catch(console.error);
