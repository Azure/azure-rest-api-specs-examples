const { AppPlatformManagementClient } = require("@azure/arm-appplatform");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Handle requests to list all API portal custom domains.
 *
 * @summary Handle requests to list all API portal custom domains.
 * x-ms-original-file: specification/appplatform/resource-manager/Microsoft.AppPlatform/preview/2022-11-01-preview/examples/ApiPortalCustomDomains_List.json
 */
async function apiPortalCustomDomainsList() {
  const subscriptionId =
    process.env["APPPLATFORM_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["APPPLATFORM_RESOURCE_GROUP"] || "myResourceGroup";
  const serviceName = "myservice";
  const apiPortalName = "default";
  const credential = new DefaultAzureCredential();
  const client = new AppPlatformManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.apiPortalCustomDomains.list(
    resourceGroupName,
    serviceName,
    apiPortalName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}
