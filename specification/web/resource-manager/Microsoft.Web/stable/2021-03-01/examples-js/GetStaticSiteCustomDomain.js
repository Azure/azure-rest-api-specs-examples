const { WebSiteManagementClient } = require("@azure/arm-appservice");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Description for Gets an existing custom domain for a particular static site.
 *
 * @summary Description for Gets an existing custom domain for a particular static site.
 * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2021-03-01/examples/GetStaticSiteCustomDomain.json
 */
async function getCustomDomainForAStaticSite() {
  const subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
  const resourceGroupName = "rg";
  const name = "testStaticSite0";
  const domainName = "custom.domain.net";
  const credential = new DefaultAzureCredential();
  const client = new WebSiteManagementClient(credential, subscriptionId);
  const result = await client.staticSites.getStaticSiteCustomDomain(
    resourceGroupName,
    name,
    domainName
  );
  console.log(result);
}

getCustomDomainForAStaticSite().catch(console.error);
