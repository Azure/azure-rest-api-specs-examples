const { ApiManagementClient } = require("@azure/arm-apimanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Checks that hostname configuration entity specified by identifier exists for specified Gateway entity.
 *
 * @summary Checks that hostname configuration entity specified by identifier exists for specified Gateway entity.
 * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementHeadGatewayHostnameConfiguration.json
 */
async function apiManagementHeadGatewayHostnameConfiguration() {
  const subscriptionId = process.env["APIMANAGEMENT_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["APIMANAGEMENT_RESOURCE_GROUP"] || "rg1";
  const serviceName = "apimService1";
  const gatewayId = "gw1";
  const hcId = "default";
  const credential = new DefaultAzureCredential();
  const client = new ApiManagementClient(credential, subscriptionId);
  const result = await client.gatewayHostnameConfiguration.getEntityTag(
    resourceGroupName,
    serviceName,
    gatewayId,
    hcId
  );
  console.log(result);
}
