const { ApiManagementClient } = require("@azure/arm-apimanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets an API Management service resource description.
 *
 * @summary Gets an API Management service resource description.
 * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementServiceGetServiceHavingMsi.json
 */
async function apiManagementServiceGetServiceHavingMsi() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const serviceName = "apimService1";
  const credential = new DefaultAzureCredential();
  const client = new ApiManagementClient(credential, subscriptionId);
  const result = await client.apiManagementService.get(resourceGroupName, serviceName);
  console.log(result);
}

apiManagementServiceGetServiceHavingMsi().catch(console.error);
