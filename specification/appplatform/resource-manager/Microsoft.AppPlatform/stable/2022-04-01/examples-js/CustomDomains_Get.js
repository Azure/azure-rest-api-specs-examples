const { AppPlatformManagementClient } = require("@azure/arm-appplatform");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get the custom domain of one lifecycle application.
 *
 * @summary Get the custom domain of one lifecycle application.
 * x-ms-original-file: specification/appplatform/resource-manager/Microsoft.AppPlatform/stable/2022-04-01/examples/CustomDomains_Get.json
 */
async function customDomainsGet() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "myResourceGroup";
  const serviceName = "myservice";
  const appName = "myapp";
  const domainName = "mydomain.com";
  const credential = new DefaultAzureCredential();
  const client = new AppPlatformManagementClient(credential, subscriptionId);
  const result = await client.customDomains.get(
    resourceGroupName,
    serviceName,
    appName,
    domainName
  );
  console.log(result);
}

customDomainsGet().catch(console.error);
