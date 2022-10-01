const { ApiManagementClient } = require("@azure/arm-apimanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes the specified Private Endpoint Connection.
 *
 * @summary Deletes the specified Private Endpoint Connection.
 * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementDeletePrivateEndpointConnection.json
 */
async function apiManagementDeletePrivateEndpointConnection() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const serviceName = "apimService1";
  const privateEndpointConnectionName = "privateEndpointConnectionName";
  const credential = new DefaultAzureCredential();
  const client = new ApiManagementClient(credential, subscriptionId);
  const result = await client.privateEndpointConnectionOperations.beginDeleteAndWait(
    resourceGroupName,
    serviceName,
    privateEndpointConnectionName
  );
  console.log(result);
}

apiManagementDeletePrivateEndpointConnection().catch(console.error);
