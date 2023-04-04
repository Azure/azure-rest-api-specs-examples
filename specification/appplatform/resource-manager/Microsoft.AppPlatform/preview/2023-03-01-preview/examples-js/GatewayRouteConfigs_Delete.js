const { AppPlatformManagementClient } = require("@azure/arm-appplatform");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Delete the Spring Cloud Gateway route config.
 *
 * @summary Delete the Spring Cloud Gateway route config.
 * x-ms-original-file: specification/appplatform/resource-manager/Microsoft.AppPlatform/preview/2023-03-01-preview/examples/GatewayRouteConfigs_Delete.json
 */
async function gatewayRouteConfigsDelete() {
  const subscriptionId =
    process.env["APPPLATFORM_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["APPPLATFORM_RESOURCE_GROUP"] || "myResourceGroup";
  const serviceName = "myservice";
  const gatewayName = "default";
  const routeConfigName = "myRouteConfig";
  const credential = new DefaultAzureCredential();
  const client = new AppPlatformManagementClient(credential, subscriptionId);
  const result = await client.gatewayRouteConfigs.beginDeleteAndWait(
    resourceGroupName,
    serviceName,
    gatewayName,
    routeConfigName
  );
  console.log(result);
}
