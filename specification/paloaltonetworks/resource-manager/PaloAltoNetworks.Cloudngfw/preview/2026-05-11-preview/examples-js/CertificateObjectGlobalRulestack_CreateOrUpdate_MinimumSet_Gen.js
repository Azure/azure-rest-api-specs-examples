const { PaloAltoNetworksCloudngfw } = require("@azure/arm-paloaltonetworksngfw");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to create a CertificateObjectGlobalRulestackResource
 *
 * @summary create a CertificateObjectGlobalRulestackResource
 * x-ms-original-file: 2026-05-11-preview/CertificateObjectGlobalRulestack_CreateOrUpdate_MinimumSet_Gen.json
 */
async function certificateObjectGlobalRulestackCreateOrUpdateMinimumSetGen() {
  const credential = new DefaultAzureCredential();
  const client = new PaloAltoNetworksCloudngfw(credential);
  const result = await client.certificateObjectGlobalRulestack.createOrUpdate("praval", "armid1", {
    certificateSelfSigned: "TRUE",
  });
  console.log(result);
}
