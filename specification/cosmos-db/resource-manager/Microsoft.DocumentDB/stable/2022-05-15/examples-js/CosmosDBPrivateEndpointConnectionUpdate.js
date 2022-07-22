const { CosmosDBManagementClient } = require("@azure/arm-cosmosdb");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Approve or reject a private endpoint connection with a given name.
 *
 * @summary Approve or reject a private endpoint connection with a given name.
 * x-ms-original-file: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2022-05-15/examples/CosmosDBPrivateEndpointConnectionUpdate.json
 */
async function approveOrRejectAPrivateEndpointConnectionWithAGivenName() {
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = "rg1";
  const accountName = "ddb1";
  const privateEndpointConnectionName = "privateEndpointConnectionName";
  const parameters = {
    privateLinkServiceConnectionState: {
      description: "Approved by johndoe@contoso.com",
      status: "Approved",
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new CosmosDBManagementClient(credential, subscriptionId);
  const result = await client.privateEndpointConnections.beginCreateOrUpdateAndWait(
    resourceGroupName,
    accountName,
    privateEndpointConnectionName,
    parameters
  );
  console.log(result);
}

approveOrRejectAPrivateEndpointConnectionWithAGivenName().catch(console.error);
