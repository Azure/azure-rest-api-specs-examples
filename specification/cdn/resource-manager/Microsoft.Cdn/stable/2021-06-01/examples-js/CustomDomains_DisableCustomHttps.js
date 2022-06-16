const { CdnManagementClient } = require("@azure/arm-cdn");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Disable https delivery of the custom domain.
 *
 * @summary Disable https delivery of the custom domain.
 * x-ms-original-file: specification/cdn/resource-manager/Microsoft.Cdn/stable/2021-06-01/examples/CustomDomains_DisableCustomHttps.json
 */
async function customDomainsDisableCustomHttps() {
  const subscriptionId = "subid";
  const resourceGroupName = "RG";
  const profileName = "profile1";
  const endpointName = "endpoint1";
  const customDomainName = "www-someDomain-net";
  const credential = new DefaultAzureCredential();
  const client = new CdnManagementClient(credential, subscriptionId);
  const result = await client.customDomains.disableCustomHttps(
    resourceGroupName,
    profileName,
    endpointName,
    customDomainName
  );
  console.log(result);
}

customDomainsDisableCustomHttps().catch(console.error);
