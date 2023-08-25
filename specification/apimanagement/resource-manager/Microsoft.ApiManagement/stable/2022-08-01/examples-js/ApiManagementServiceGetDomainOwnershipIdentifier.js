const { ApiManagementClient } = require("@azure/arm-apimanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get the custom domain ownership identifier for an API Management service.
 *
 * @summary Get the custom domain ownership identifier for an API Management service.
 * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementServiceGetDomainOwnershipIdentifier.json
 */
async function apiManagementServiceGetDomainOwnershipIdentifier() {
  const subscriptionId = process.env["APIMANAGEMENT_SUBSCRIPTION_ID"] || "subid";
  const credential = new DefaultAzureCredential();
  const client = new ApiManagementClient(credential, subscriptionId);
  const result = await client.apiManagementService.getDomainOwnershipIdentifier();
  console.log(result);
}
