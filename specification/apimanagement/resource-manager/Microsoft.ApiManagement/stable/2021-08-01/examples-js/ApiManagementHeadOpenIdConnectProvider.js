const { ApiManagementClient } = require("@azure/arm-apimanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the entity state (Etag) version of the openIdConnectProvider specified by its identifier.
 *
 * @summary Gets the entity state (Etag) version of the openIdConnectProvider specified by its identifier.
 * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementHeadOpenIdConnectProvider.json
 */
async function apiManagementHeadOpenIdConnectProvider() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const serviceName = "apimService1";
  const opid = "templateOpenIdConnect2";
  const credential = new DefaultAzureCredential();
  const client = new ApiManagementClient(credential, subscriptionId);
  const result = await client.openIdConnectProvider.getEntityTag(
    resourceGroupName,
    serviceName,
    opid
  );
  console.log(result);
}

apiManagementHeadOpenIdConnectProvider().catch(console.error);
