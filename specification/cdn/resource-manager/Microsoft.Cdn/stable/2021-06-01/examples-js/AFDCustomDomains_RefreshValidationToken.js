const { CdnManagementClient } = require("@azure/arm-cdn");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Updates the domain validation token.
 *
 * @summary Updates the domain validation token.
 * x-ms-original-file: specification/cdn/resource-manager/Microsoft.Cdn/stable/2021-06-01/examples/AFDCustomDomains_RefreshValidationToken.json
 */
async function afdCustomDomainsDelete() {
  const subscriptionId = "subid";
  const resourceGroupName = "RG";
  const profileName = "profile1";
  const customDomainName = "domain1";
  const credential = new DefaultAzureCredential();
  const client = new CdnManagementClient(credential, subscriptionId);
  const result = await client.afdCustomDomains.beginRefreshValidationTokenAndWait(
    resourceGroupName,
    profileName,
    customDomainName
  );
  console.log(result);
}

afdCustomDomainsDelete().catch(console.error);
