const { ContainerServiceClient } = require("@azure/arm-containerservice");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Updates a private endpoint connection.
 *
 * @summary Updates a private endpoint connection.
 * x-ms-original-file: specification/containerservice/resource-manager/Microsoft.ContainerService/aks/preview/2023-08-02-preview/examples/PrivateEndpointConnectionsUpdate.json
 */
async function updatePrivateEndpointConnection() {
  const subscriptionId = process.env["CONTAINERSERVICE_SUBSCRIPTION_ID"] || "subid1";
  const resourceGroupName = process.env["CONTAINERSERVICE_RESOURCE_GROUP"] || "rg1";
  const resourceName = "clustername1";
  const privateEndpointConnectionName = "privateendpointconnection1";
  const parameters = {
    privateLinkServiceConnectionState: { status: "Approved" },
  };
  const credential = new DefaultAzureCredential();
  const client = new ContainerServiceClient(credential, subscriptionId);
  const result = await client.privateEndpointConnections.update(
    resourceGroupName,
    resourceName,
    privateEndpointConnectionName,
    parameters
  );
  console.log(result);
}
