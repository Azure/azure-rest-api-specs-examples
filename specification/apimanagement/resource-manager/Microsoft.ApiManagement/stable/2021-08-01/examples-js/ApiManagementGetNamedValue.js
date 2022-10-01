const { ApiManagementClient } = require("@azure/arm-apimanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the details of the named value specified by its identifier.
 *
 * @summary Gets the details of the named value specified by its identifier.
 * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementGetNamedValue.json
 */
async function apiManagementGetNamedValue() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const serviceName = "apimService1";
  const namedValueId = "testarmTemplateproperties2";
  const credential = new DefaultAzureCredential();
  const client = new ApiManagementClient(credential, subscriptionId);
  const result = await client.namedValue.get(resourceGroupName, serviceName, namedValueId);
  console.log(result);
}

apiManagementGetNamedValue().catch(console.error);
