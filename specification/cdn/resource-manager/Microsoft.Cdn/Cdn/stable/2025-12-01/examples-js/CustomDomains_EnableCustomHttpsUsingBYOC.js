const { CdnManagementClient } = require("@azure/arm-cdn");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to enable https delivery of the custom domain.
 *
 * @summary enable https delivery of the custom domain.
 * x-ms-original-file: 2025-12-01/CustomDomains_EnableCustomHttpsUsingBYOC.json
 */
async function customDomainsEnableCustomHttpsUsingYourOwnCertificate() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new CdnManagementClient(credential, subscriptionId);
  const result = await client.customDomains.enableCustomHttps(
    "RG",
    "profile1",
    "endpoint1",
    "www-someDomain-net",
  );
  console.log(result);
}
