const { PaloAltoNetworksCloudngfw } = require("@azure/arm-paloaltonetworksngfw");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Delete a CertificateObjectGlobalRulestackResource
 *
 * @summary Delete a CertificateObjectGlobalRulestackResource
 * x-ms-original-file: specification/paloaltonetworks/resource-manager/PaloAltoNetworks.Cloudngfw/stable/2022-08-29/examples/CertificateObjectGlobalRulestack_Delete_MaximumSet_Gen.json
 */
async function certificateObjectGlobalRulestackDeleteMaximumSetGen() {
  const globalRulestackName = "praval";
  const name = "armid1";
  const credential = new DefaultAzureCredential();
  const client = new PaloAltoNetworksCloudngfw(credential);
  const result = await client.certificateObjectGlobalRulestack.beginDeleteAndWait(
    globalRulestackName,
    name
  );
  console.log(result);
}
