const { PrivateLinkScopesClient } = require("@azure/arm-kubernetesconfiguration-privatelinkscopes");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Deletes a private endpoint connection with a given name.
 *
 * @summary Deletes a private endpoint connection with a given name.
 * x-ms-original-file: specification/kubernetesconfiguration/resource-manager/Microsoft.KubernetesConfiguration/privateLinkScopes/preview/2024-11-01-preview/examples/PrivateEndpointConnectionDelete.json
 */
async function deletesAPrivateEndpointConnectionWithAGivenName() {
  const subscriptionId =
    process.env["KUBERNETESCONFIGURATION_SUBSCRIPTION_ID"] ||
    "00000000-1111-2222-3333-444444444444";
  const resourceGroupName =
    process.env["KUBERNETESCONFIGURATION_RESOURCE_GROUP"] || "myResourceGroup";
  const scopeName = "myPrivateLinkScope";
  const privateEndpointConnectionName = "private-endpoint-connection-name";
  const credential = new DefaultAzureCredential();
  const client = new PrivateLinkScopesClient(credential, subscriptionId);
  const result = await client.privateEndpointConnections.delete(
    resourceGroupName,
    scopeName,
    privateEndpointConnectionName,
  );
  console.log(result);
}
