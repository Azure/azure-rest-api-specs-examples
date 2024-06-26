const { AppPlatformManagementClient } = require("@azure/arm-appplatform");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Delete the custom domain of one lifecycle application.
 *
 * @summary Delete the custom domain of one lifecycle application.
 * x-ms-original-file: specification/appplatform/resource-manager/Microsoft.AppPlatform/preview/2023-03-01-preview/examples/CustomDomains_Delete.json
 */
async function customDomainsDelete() {
  const subscriptionId =
    process.env["APPPLATFORM_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["APPPLATFORM_RESOURCE_GROUP"] || "myResourceGroup";
  const serviceName = "myservice";
  const appName = "myapp";
  const domainName = "mydomain.com";
  const credential = new DefaultAzureCredential();
  const client = new AppPlatformManagementClient(credential, subscriptionId);
  const result = await client.customDomains.beginDeleteAndWait(
    resourceGroupName,
    serviceName,
    appName,
    domainName
  );
  console.log(result);
}
