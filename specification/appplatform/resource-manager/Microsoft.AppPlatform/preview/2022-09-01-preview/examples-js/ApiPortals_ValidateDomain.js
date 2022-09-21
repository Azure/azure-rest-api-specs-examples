const { AppPlatformManagementClient } = require("@azure/arm-appplatform");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Check the domains are valid as well as not in use.
 *
 * @summary Check the domains are valid as well as not in use.
 * x-ms-original-file: specification/appplatform/resource-manager/Microsoft.AppPlatform/preview/2022-09-01-preview/examples/ApiPortals_ValidateDomain.json
 */
async function apiPortalsValidateDomain() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "myResourceGroup";
  const serviceName = "myservice";
  const apiPortalName = "default";
  const validatePayload = { name: "mydomain.io" };
  const credential = new DefaultAzureCredential();
  const client = new AppPlatformManagementClient(credential, subscriptionId);
  const result = await client.apiPortals.validateDomain(
    resourceGroupName,
    serviceName,
    apiPortalName,
    validatePayload
  );
  console.log(result);
}

apiPortalsValidateDomain().catch(console.error);
