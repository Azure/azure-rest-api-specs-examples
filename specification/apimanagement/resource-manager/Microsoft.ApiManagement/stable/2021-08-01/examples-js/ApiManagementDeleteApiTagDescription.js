const { ApiManagementClient } = require("@azure/arm-apimanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Delete tag description for the Api.
 *
 * @summary Delete tag description for the Api.
 * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementDeleteApiTagDescription.json
 */
async function apiManagementDeleteApiTagDescription() {
  const subscriptionId = process.env["APIMANAGEMENT_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["APIMANAGEMENT_RESOURCE_GROUP"] || "rg1";
  const serviceName = "apimService1";
  const apiId = "59d5b28d1f7fab116c282650";
  const tagDescriptionId = "59d5b28e1f7fab116402044e";
  const ifMatch = "*";
  const credential = new DefaultAzureCredential();
  const client = new ApiManagementClient(credential, subscriptionId);
  const result = await client.apiTagDescription.delete(
    resourceGroupName,
    serviceName,
    apiId,
    tagDescriptionId,
    ifMatch
  );
  console.log(result);
}
