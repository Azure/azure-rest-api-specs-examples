const { AppPlatformManagementClient } = require("@azure/arm-appplatform");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create or update custom domain of one lifecycle application.
 *
 * @summary Create or update custom domain of one lifecycle application.
 * x-ms-original-file: specification/appplatform/resource-manager/Microsoft.AppPlatform/preview/2022-09-01-preview/examples/CustomDomains_CreateOrUpdate.json
 */
async function customDomainsCreateOrUpdate() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "myResourceGroup";
  const serviceName = "myservice";
  const appName = "myapp";
  const domainName = "mydomain.com";
  const domainResource = {
    properties: {
      certName: "mycert",
      thumbprint: "934367bf1c97033f877db0f15cb1b586957d3133",
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new AppPlatformManagementClient(credential, subscriptionId);
  const result = await client.customDomains.beginCreateOrUpdateAndWait(
    resourceGroupName,
    serviceName,
    appName,
    domainName,
    domainResource
  );
  console.log(result);
}

customDomainsCreateOrUpdate().catch(console.error);
