const { AppPlatformManagementClient } = require("@azure/arm-appplatform");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create the default Spring Cloud Gateway or update the existing Spring Cloud Gateway.
 *
 * @summary Create the default Spring Cloud Gateway or update the existing Spring Cloud Gateway.
 * x-ms-original-file: specification/appplatform/resource-manager/Microsoft.AppPlatform/preview/2023-03-01-preview/examples/Gateways_CreateOrUpdate.json
 */
async function gatewaysCreateOrUpdate() {
  const subscriptionId =
    process.env["APPPLATFORM_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["APPPLATFORM_RESOURCE_GROUP"] || "myResourceGroup";
  const serviceName = "myservice";
  const gatewayName = "default";
  const gatewayResource = {
    properties: { public: true, resourceRequests: { cpu: "1", memory: "1G" } },
    sku: { name: "E0", capacity: 2, tier: "Enterprise" },
  };
  const credential = new DefaultAzureCredential();
  const client = new AppPlatformManagementClient(credential, subscriptionId);
  const result = await client.gateways.beginCreateOrUpdateAndWait(
    resourceGroupName,
    serviceName,
    gatewayName,
    gatewayResource
  );
  console.log(result);
}
