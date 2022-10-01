const { ApiManagementClient } = require("@azure/arm-apimanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the details of the Api Version Set specified by its identifier.
 *
 * @summary Gets the details of the Api Version Set specified by its identifier.
 * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementGetApiVersionSet.json
 */
async function apiManagementGetApiVersionSet() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const serviceName = "apimService1";
  const versionSetId = "vs1";
  const credential = new DefaultAzureCredential();
  const client = new ApiManagementClient(credential, subscriptionId);
  const result = await client.apiVersionSet.get(resourceGroupName, serviceName, versionSetId);
  console.log(result);
}

apiManagementGetApiVersionSet().catch(console.error);
