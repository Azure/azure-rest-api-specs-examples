const { HybridComputeManagementClient } = require("@azure/arm-hybridcompute");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Deletes a private endpoint connection with a given name.
 *
 * @summary Deletes a private endpoint connection with a given name.
 * x-ms-original-file: specification/hybridcompute/resource-manager/Microsoft.HybridCompute/preview/2025-02-19-preview/examples/privateEndpoint/PrivateEndpointConnection_Delete.json
 */
async function deletesAPrivateEndpointConnectionWithAGivenName() {
  const subscriptionId =
    process.env["HYBRIDCOMPUTE_SUBSCRIPTION_ID"] || "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = process.env["HYBRIDCOMPUTE_RESOURCE_GROUP"] || "myResourceGroup";
  const scopeName = "myPrivateLinkScope";
  const privateEndpointConnectionName = "private-endpoint-connection-name";
  const credential = new DefaultAzureCredential();
  const client = new HybridComputeManagementClient(credential, subscriptionId);
  const result = await client.privateEndpointConnections.beginDeleteAndWait(
    resourceGroupName,
    scopeName,
    privateEndpointConnectionName,
  );
  console.log(result);
}
