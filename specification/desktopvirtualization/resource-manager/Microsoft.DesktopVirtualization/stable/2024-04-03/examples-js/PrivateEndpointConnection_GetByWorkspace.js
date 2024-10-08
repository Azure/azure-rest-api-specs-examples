const { DesktopVirtualizationAPIClient } = require("@azure/arm-desktopvirtualization");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get a private endpoint connection.
 *
 * @summary Get a private endpoint connection.
 * x-ms-original-file: specification/desktopvirtualization/resource-manager/Microsoft.DesktopVirtualization/stable/2024-04-03/examples/PrivateEndpointConnection_GetByWorkspace.json
 */
async function privateEndpointConnectionGetByWorkspace() {
  const subscriptionId =
    process.env["DESKTOPVIRTUALIZATION_SUBSCRIPTION_ID"] || "daefabc0-95b4-48b3-b645-8a753a63c4fa";
  const resourceGroupName = process.env["DESKTOPVIRTUALIZATION_RESOURCE_GROUP"] || "resourceGroup1";
  const workspaceName = "workspace1";
  const privateEndpointConnectionName = "workspace1.377103f1-5179-4bdf-8556-4cdd3207cc5b";
  const credential = new DefaultAzureCredential();
  const client = new DesktopVirtualizationAPIClient(credential, subscriptionId);
  const result = await client.privateEndpointConnections.getByWorkspace(
    resourceGroupName,
    workspaceName,
    privateEndpointConnectionName,
  );
  console.log(result);
}
