const { ApiManagementClient } = require("@azure/arm-apimanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Description for Gets the private link resources
 *
 * @summary Description for Gets the private link resources
 * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementGetPrivateLinkGroupResource.json
 */
async function apiManagementGetPrivateLinkGroupResource() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const serviceName = "apimService1";
  const privateLinkSubResourceName = "privateLinkSubResourceName";
  const credential = new DefaultAzureCredential();
  const client = new ApiManagementClient(credential, subscriptionId);
  const result = await client.privateEndpointConnectionOperations.getPrivateLinkResource(
    resourceGroupName,
    serviceName,
    privateLinkSubResourceName
  );
  console.log(result);
}

apiManagementGetPrivateLinkGroupResource().catch(console.error);
