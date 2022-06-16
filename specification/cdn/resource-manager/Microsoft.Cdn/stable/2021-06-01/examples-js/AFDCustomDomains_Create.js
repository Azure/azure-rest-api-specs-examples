const { CdnManagementClient } = require("@azure/arm-cdn");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates a new domain within the specified profile.
 *
 * @summary Creates a new domain within the specified profile.
 * x-ms-original-file: specification/cdn/resource-manager/Microsoft.Cdn/stable/2021-06-01/examples/AFDCustomDomains_Create.json
 */
async function afdCustomDomainsCreate() {
  const subscriptionId = "subid";
  const resourceGroupName = "RG";
  const profileName = "profile1";
  const customDomainName = "domain1";
  const customDomain = {
    azureDnsZone: { id: "" },
    hostName: "www.someDomain.net",
    tlsSettings: {
      certificateType: "ManagedCertificate",
      minimumTlsVersion: "TLS12",
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new CdnManagementClient(credential, subscriptionId);
  const result = await client.afdCustomDomains.beginCreateAndWait(
    resourceGroupName,
    profileName,
    customDomainName,
    customDomain
  );
  console.log(result);
}

afdCustomDomainsCreate().catch(console.error);
