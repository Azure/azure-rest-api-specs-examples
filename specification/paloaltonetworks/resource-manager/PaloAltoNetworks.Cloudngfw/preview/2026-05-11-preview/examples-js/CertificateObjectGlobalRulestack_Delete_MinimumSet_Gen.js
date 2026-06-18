const { PaloAltoNetworksCloudngfw } = require("@azure/arm-paloaltonetworksngfw");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to delete a CertificateObjectGlobalRulestackResource
 *
 * @summary delete a CertificateObjectGlobalRulestackResource
 * x-ms-original-file: 2026-05-11-preview/CertificateObjectGlobalRulestack_Delete_MinimumSet_Gen.json
 */
async function certificateObjectGlobalRulestackDeleteMinimumSetGen() {
  const credential = new DefaultAzureCredential();
  const client = new PaloAltoNetworksCloudngfw(credential);
  await client.certificateObjectGlobalRulestack.delete("praval", "armid1");
}
