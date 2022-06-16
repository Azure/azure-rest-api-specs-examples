const { RecoveryServicesBackupClient } = require("@azure/arm-recoveryservicesbackup");
const { DefaultAzureCredential } = require("@azure/identity");

async function updatePrivateEndpointConnection() {
  const subscriptionId = "04cf684a-d41f-4550-9f70-7708a3a2283b";
  const vaultName = "gaallavaultbvtd2msi";
  const resourceGroupName = "gaallaRG";
  const privateEndpointConnectionName = "gaallatestpe2.5704c932-249a-490b-a142-1396838cd3b";
  const parameters = {
    properties: {
      privateEndpoint: {
        id: "/subscriptions/04cf684a-d41f-4550-9f70-7708a3a2283b/resourceGroups/gaallaRG/providers/Microsoft.Network/privateEndpoints/gaallatestpe3",
      },
      privateLinkServiceConnectionState: {
        description: "Approved by johndoe@company.com",
        status: "Approved",
      },
      provisioningState: "Succeeded",
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new RecoveryServicesBackupClient(credential, subscriptionId);
  const result = await client.privateEndpointConnectionOperations.beginPutAndWait(
    vaultName,
    resourceGroupName,
    privateEndpointConnectionName,
    parameters
  );
  console.log(result);
}

updatePrivateEndpointConnection().catch(console.error);
