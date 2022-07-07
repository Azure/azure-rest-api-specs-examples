const { WebSiteManagementClient } = require("@azure/arm-appservice");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Description for Gets a private endpoint connection
 *
 * @summary Description for Gets a private endpoint connection
 * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2022-03-01/examples/AppServiceEnvironments_GetPrivateEndpointConnection.json
 */
async function getsAPrivateEndpointConnection() {
  const subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
  const resourceGroupName = "test-rg";
  const name = "test-ase";
  const privateEndpointConnectionName = "fa38656c-034e-43d8-adce-fe06ce039c98";
  const credential = new DefaultAzureCredential();
  const client = new WebSiteManagementClient(credential, subscriptionId);
  const result = await client.appServiceEnvironments.getPrivateEndpointConnection(
    resourceGroupName,
    name,
    privateEndpointConnectionName
  );
  console.log(result);
}

getsAPrivateEndpointConnection().catch(console.error);
