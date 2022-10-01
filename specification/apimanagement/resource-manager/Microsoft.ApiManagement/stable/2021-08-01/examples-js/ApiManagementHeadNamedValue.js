const { ApiManagementClient } = require("@azure/arm-apimanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the entity state (Etag) version of the named value specified by its identifier.
 *
 * @summary Gets the entity state (Etag) version of the named value specified by its identifier.
 * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementHeadNamedValue.json
 */
async function apiManagementHeadNamedValue() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const serviceName = "apimService1";
  const namedValueId = "testarmTemplateproperties2";
  const credential = new DefaultAzureCredential();
  const client = new ApiManagementClient(credential, subscriptionId);
  const result = await client.namedValue.getEntityTag(resourceGroupName, serviceName, namedValueId);
  console.log(result);
}

apiManagementHeadNamedValue().catch(console.error);
