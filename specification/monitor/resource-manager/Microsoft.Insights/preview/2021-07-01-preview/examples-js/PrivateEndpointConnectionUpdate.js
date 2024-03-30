const { MonitorClient } = require("@azure/arm-monitor");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Approve or reject a private endpoint connection with a given name.
 *
 * @summary Approve or reject a private endpoint connection with a given name.
 * x-ms-original-file: specification/monitor/resource-manager/Microsoft.Insights/preview/2021-07-01-preview/examples/PrivateEndpointConnectionUpdate.json
 */
async function approveOrRejectAPrivateEndpointConnectionWithAGivenName() {
  const subscriptionId =
    process.env["MONITOR_SUBSCRIPTION_ID"] || "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = process.env["MONITOR_RESOURCE_GROUP"] || "MyResourceGroup";
  const scopeName = "MyPrivateLinkScope";
  const privateEndpointConnectionName = "private-endpoint-connection-name";
  const parameters = {
    privateLinkServiceConnectionState: {
      description: "Approved by johndoe@contoso.com",
      status: "Approved",
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new MonitorClient(credential, subscriptionId);
  const result = await client.privateEndpointConnections.beginCreateOrUpdateAndWait(
    resourceGroupName,
    scopeName,
    privateEndpointConnectionName,
    parameters,
  );
  console.log(result);
}
