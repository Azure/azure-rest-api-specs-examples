const { ApiManagementClient } = require("@azure/arm-apimanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the details of the Private Endpoint Connection specified by its identifier.
 *
 * @summary Gets the details of the Private Endpoint Connection specified by its identifier.
 * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementGetPrivateEndpointConnection.json
 */
async function apiManagementGetPrivateEndpointConnection() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const serviceName = "apimService1";
  const privateEndpointConnectionName = "privateEndpointConnectionName";
  const credential = new DefaultAzureCredential();
  const client = new ApiManagementClient(credential, subscriptionId);
  const result = await client.privateEndpointConnectionOperations.getByName(
    resourceGroupName,
    serviceName,
    privateEndpointConnectionName
  );
  console.log(result);
}

apiManagementGetPrivateEndpointConnection().catch(console.error);
