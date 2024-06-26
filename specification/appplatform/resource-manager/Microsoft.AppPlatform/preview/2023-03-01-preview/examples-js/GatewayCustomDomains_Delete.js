const { AppPlatformManagementClient } = require("@azure/arm-appplatform");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Delete the Spring Cloud Gateway custom domain.
 *
 * @summary Delete the Spring Cloud Gateway custom domain.
 * x-ms-original-file: specification/appplatform/resource-manager/Microsoft.AppPlatform/preview/2023-03-01-preview/examples/GatewayCustomDomains_Delete.json
 */
async function gatewayCustomDomainsDelete() {
  const subscriptionId =
    process.env["APPPLATFORM_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["APPPLATFORM_RESOURCE_GROUP"] || "myResourceGroup";
  const serviceName = "myservice";
  const gatewayName = "default";
  const domainName = "myDomainName";
  const credential = new DefaultAzureCredential();
  const client = new AppPlatformManagementClient(credential, subscriptionId);
  const result = await client.gatewayCustomDomains.beginDeleteAndWait(
    resourceGroupName,
    serviceName,
    gatewayName,
    domainName
  );
  console.log(result);
}
