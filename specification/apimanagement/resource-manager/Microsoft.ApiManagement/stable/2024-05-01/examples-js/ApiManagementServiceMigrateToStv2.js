const { ApiManagementClient } = require("@azure/arm-apimanagement");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Upgrades an API Management service to the Stv2 platform. For details refer to https://aka.ms/apim-migrate-stv2. This change is not reversible. This is long running operation and could take several minutes to complete.
 *
 * @summary Upgrades an API Management service to the Stv2 platform. For details refer to https://aka.ms/apim-migrate-stv2. This change is not reversible. This is long running operation and could take several minutes to complete.
 * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/ApiManagementServiceMigrateToStv2.json
 */
async function apiManagementMigrateService() {
  const subscriptionId =
    process.env["APIMANAGEMENT_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["APIMANAGEMENT_RESOURCE_GROUP"] || "rg1";
  const serviceName = "apimService1";
  const parameters = { mode: "PreserveIp" };
  const options = {
    parameters,
  };
  const credential = new DefaultAzureCredential();
  const client = new ApiManagementClient(credential, subscriptionId);
  const result = await client.apiManagementService.beginMigrateToStv2AndWait(
    resourceGroupName,
    serviceName,
    options,
  );
  console.log(result);
}
