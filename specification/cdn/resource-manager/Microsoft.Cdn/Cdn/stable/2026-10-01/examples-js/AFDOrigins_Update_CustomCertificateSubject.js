const { CdnManagementClient } = require("@azure/arm-cdn");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to updates an existing origin within an origin group.
 *
 * @summary updates an existing origin within an origin group.
 * x-ms-original-file: 2026-10-01/AFDOrigins_Update_CustomCertificateSubject.json
 */
async function afdOriginsUpdateCustomCertificateSubject() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new CdnManagementClient(credential, subscriptionId);
  const result = await client.afdOrigins.update("RG", "profile1", "origingroup1", "origin1", {
    certificateNameCheckValidationMode: "CustomCertificateSubject",
    customCertificateSubjects: ["host1.foo.com", "*.foo.com"],
    enabledState: "Enabled",
    enforceCertificateNameCheck: true,
    hostName: "host1.blob.core.windows.net",
    httpPort: 80,
    httpsPort: 443,
    originHostHeader: "host1.foo.com",
  });
  console.log(result);
}
