const { AppPlatformManagementClient } = require("@azure/arm-appplatform");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create or update the API portal custom domain.
 *
 * @summary Create or update the API portal custom domain.
 * x-ms-original-file: specification/appplatform/resource-manager/Microsoft.AppPlatform/preview/2022-09-01-preview/examples/ApiPortalCustomDomains_CreateOrUpdate.json
 */
async function apiPortalCustomDomainsCreateOrUpdate() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "myResourceGroup";
  const serviceName = "myservice";
  const apiPortalName = "default";
  const domainName = "myDomainName";
  const apiPortalCustomDomainResource = {
    properties: { thumbprint: "*" },
  };
  const credential = new DefaultAzureCredential();
  const client = new AppPlatformManagementClient(credential, subscriptionId);
  const result = await client.apiPortalCustomDomains.beginCreateOrUpdateAndWait(
    resourceGroupName,
    serviceName,
    apiPortalName,
    domainName,
    apiPortalCustomDomainResource
  );
  console.log(result);
}

apiPortalCustomDomainsCreateOrUpdate().catch(console.error);
