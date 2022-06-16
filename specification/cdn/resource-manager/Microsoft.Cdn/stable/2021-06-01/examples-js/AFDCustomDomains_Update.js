const { CdnManagementClient } = require("@azure/arm-cdn");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Updates an existing domain within a profile.
 *
 * @summary Updates an existing domain within a profile.
 * x-ms-original-file: specification/cdn/resource-manager/Microsoft.Cdn/stable/2021-06-01/examples/AFDCustomDomains_Update.json
 */
async function afdCustomDomainsUpdate() {
  const subscriptionId = "subid";
  const resourceGroupName = "RG";
  const profileName = "profile1";
  const customDomainName = "domain1";
  const customDomainUpdateProperties = {
    azureDnsZone: { id: "" },
    tlsSettings: {
      certificateType: "CustomerCertificate",
      minimumTlsVersion: "TLS12",
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new CdnManagementClient(credential, subscriptionId);
  const result = await client.afdCustomDomains.beginUpdateAndWait(
    resourceGroupName,
    profileName,
    customDomainName,
    customDomainUpdateProperties
  );
  console.log(result);
}

afdCustomDomainsUpdate().catch(console.error);
