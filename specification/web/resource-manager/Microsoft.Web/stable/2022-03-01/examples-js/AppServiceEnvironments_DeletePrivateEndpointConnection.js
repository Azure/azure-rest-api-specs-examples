const { WebSiteManagementClient } = require("@azure/arm-appservice");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Description for Deletes a private endpoint connection
 *
 * @summary Description for Deletes a private endpoint connection
 * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2022-03-01/examples/AppServiceEnvironments_DeletePrivateEndpointConnection.json
 */
async function deletesAPrivateEndpointConnection() {
  const subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
  const resourceGroupName = "test-rg";
  const name = "test-ase";
  const privateEndpointConnectionName = "fa38656c-034e-43d8-adce-fe06ce039c98";
  const credential = new DefaultAzureCredential();
  const client = new WebSiteManagementClient(credential, subscriptionId);
  const result = await client.appServiceEnvironments.beginDeletePrivateEndpointConnectionAndWait(
    resourceGroupName,
    name,
    privateEndpointConnectionName
  );
  console.log(result);
}

deletesAPrivateEndpointConnection().catch(console.error);
