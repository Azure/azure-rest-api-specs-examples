const { CognitiveServicesManagementClient } = require("@azure/arm-cognitiveservices");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to updates a Cognitive Services account
 *
 * @summary updates a Cognitive Services account
 * x-ms-original-file: 2026-11-01/UpdateAccountPreserveTrustedCertificateSources.json
 */
async function updateAccountTagsWhilePreservingPrivateCATrustWithNull() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const client = new CognitiveServicesManagementClient(credential, subscriptionId);
  const result = await client.accounts.update("myResourceGroup", "myAccount", {
    tags: { environment: "production" },
    properties: { trustedCertificateSources: [] },
  });
  console.log(result);
}
