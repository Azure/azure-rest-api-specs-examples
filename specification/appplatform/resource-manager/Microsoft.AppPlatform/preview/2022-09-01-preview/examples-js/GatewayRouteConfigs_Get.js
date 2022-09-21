const { AppPlatformManagementClient } = require("@azure/arm-appplatform");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get the Spring Cloud Gateway route configs.
 *
 * @summary Get the Spring Cloud Gateway route configs.
 * x-ms-original-file: specification/appplatform/resource-manager/Microsoft.AppPlatform/preview/2022-09-01-preview/examples/GatewayRouteConfigs_Get.json
 */
async function gatewayRouteConfigsGet() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "myResourceGroup";
  const serviceName = "myservice";
  const gatewayName = "default";
  const routeConfigName = "myRouteConfig";
  const credential = new DefaultAzureCredential();
  const client = new AppPlatformManagementClient(credential, subscriptionId);
  const result = await client.gatewayRouteConfigs.get(
    resourceGroupName,
    serviceName,
    gatewayName,
    routeConfigName
  );
  console.log(result);
}

gatewayRouteConfigsGet().catch(console.error);
