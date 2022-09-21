const { AppPlatformManagementClient } = require("@azure/arm-appplatform");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Handle requests to list all Spring Cloud Gateway route configs.
 *
 * @summary Handle requests to list all Spring Cloud Gateway route configs.
 * x-ms-original-file: specification/appplatform/resource-manager/Microsoft.AppPlatform/preview/2022-09-01-preview/examples/GatewayRouteConfigs_List.json
 */
async function gatewayRouteConfigsList() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "myResourceGroup";
  const serviceName = "myservice";
  const gatewayName = "default";
  const credential = new DefaultAzureCredential();
  const client = new AppPlatformManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.gatewayRouteConfigs.list(
    resourceGroupName,
    serviceName,
    gatewayName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

gatewayRouteConfigsList().catch(console.error);
