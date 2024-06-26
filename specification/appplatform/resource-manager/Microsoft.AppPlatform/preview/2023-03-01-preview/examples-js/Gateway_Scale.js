const { AppPlatformManagementClient } = require("@azure/arm-appplatform");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Operation to update an exiting Spring Cloud Gateway capacity.
 *
 * @summary Operation to update an exiting Spring Cloud Gateway capacity.
 * x-ms-original-file: specification/appplatform/resource-manager/Microsoft.AppPlatform/preview/2023-03-01-preview/examples/Gateway_Scale.json
 */
async function gatewayScale() {
  const subscriptionId =
    process.env["APPPLATFORM_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["APPPLATFORM_RESOURCE_GROUP"] || "myResourceGroup";
  const serviceName = "myservice";
  const gatewayName = "default";
  const gatewayCapacityResource = {
    sku: { name: "E0", capacity: 2, tier: "Enterprise" },
  };
  const credential = new DefaultAzureCredential();
  const client = new AppPlatformManagementClient(credential, subscriptionId);
  const result = await client.gateways.beginUpdateCapacityAndWait(
    resourceGroupName,
    serviceName,
    gatewayName,
    gatewayCapacityResource
  );
  console.log(result);
}
