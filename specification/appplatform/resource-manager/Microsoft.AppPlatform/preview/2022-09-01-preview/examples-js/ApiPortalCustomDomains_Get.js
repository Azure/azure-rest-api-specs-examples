const { AppPlatformManagementClient } = require("@azure/arm-appplatform");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get the API portal custom domain.
 *
 * @summary Get the API portal custom domain.
 * x-ms-original-file: specification/appplatform/resource-manager/Microsoft.AppPlatform/preview/2022-09-01-preview/examples/ApiPortalCustomDomains_Get.json
 */
async function apiPortalCustomDomainsGet() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "myResourceGroup";
  const serviceName = "myservice";
  const apiPortalName = "default";
  const domainName = "myDomainName";
  const credential = new DefaultAzureCredential();
  const client = new AppPlatformManagementClient(credential, subscriptionId);
  const result = await client.apiPortalCustomDomains.get(
    resourceGroupName,
    serviceName,
    apiPortalName,
    domainName
  );
  console.log(result);
}

apiPortalCustomDomainsGet().catch(console.error);
