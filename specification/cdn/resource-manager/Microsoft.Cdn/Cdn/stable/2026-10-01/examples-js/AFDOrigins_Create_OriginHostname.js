const { CdnManagementClient } = require("@azure/arm-cdn");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates a new origin within the specified origin group.
 *
 * @summary creates a new origin within the specified origin group.
 * x-ms-original-file: 2026-10-01/AFDOrigins_Create_OriginHostname.json
 */
async function afdOriginsCreateOriginHostname() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new CdnManagementClient(credential, subscriptionId);
  const result = await client.afdOrigins.create("RG", "profile1", "origingroup1", "origin1", {
    certificateNameCheckValidationMode: "OriginHostname",
    enabledState: "Enabled",
    enforceCertificateNameCheck: true,
    hostName: "host1.blob.core.windows.net",
    httpPort: 80,
    httpsPort: 443,
    originHostHeader: "host1.foo.com",
  });
  console.log(result);
}
