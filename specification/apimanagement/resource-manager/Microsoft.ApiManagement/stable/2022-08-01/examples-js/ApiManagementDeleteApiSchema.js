const { ApiManagementClient } = require("@azure/arm-apimanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes the schema configuration at the Api.
 *
 * @summary Deletes the schema configuration at the Api.
 * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementDeleteApiSchema.json
 */
async function apiManagementDeleteApiSchema() {
  const subscriptionId = process.env["APIMANAGEMENT_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["APIMANAGEMENT_RESOURCE_GROUP"] || "rg1";
  const serviceName = "apimService1";
  const apiId = "59d5b28d1f7fab116c282650";
  const schemaId = "59d5b28e1f7fab116402044e";
  const ifMatch = "*";
  const credential = new DefaultAzureCredential();
  const client = new ApiManagementClient(credential, subscriptionId);
  const result = await client.apiSchema.delete(
    resourceGroupName,
    serviceName,
    apiId,
    schemaId,
    ifMatch
  );
  console.log(result);
}
