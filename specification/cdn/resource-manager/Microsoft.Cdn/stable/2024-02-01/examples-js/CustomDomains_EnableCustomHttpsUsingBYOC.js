const { CdnManagementClient } = require("@azure/arm-cdn");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Enable https delivery of the custom domain.
 *
 * @summary Enable https delivery of the custom domain.
 * x-ms-original-file: specification/cdn/resource-manager/Microsoft.Cdn/stable/2024-02-01/examples/CustomDomains_EnableCustomHttpsUsingBYOC.json
 */
async function customDomainsEnableCustomHttpsUsingYourOwnCertificate() {
  const subscriptionId = process.env["CDN_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["CDN_RESOURCE_GROUP"] || "RG";
  const profileName = "profile1";
  const endpointName = "endpoint1";
  const customDomainName = "www-someDomain-net";
  const credential = new DefaultAzureCredential();
  const client = new CdnManagementClient(credential, subscriptionId);
  const result = await client.customDomains.beginEnableCustomHttpsAndWait(
    resourceGroupName,
    profileName,
    endpointName,
    customDomainName,
  );
  console.log(result);
}
